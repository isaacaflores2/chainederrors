package chain

import "fmt"

type errorChain struct {
	msg  string
	err  error
	next error
}

func Wrap(err, wrappedErr error) error {
	if err == nil {
		return wrappedErr
	}

	return &errorChain{
		msg:  fmt.Sprintf("%s , %s", err, wrappedErr),
		err:  err,
		next: wrappedErr,
	}
}

func (e *errorChain) Error() string {
	if e == nil {
		return ""
	}
	return e.msg
}

func (e *errorChain) Is(target error) bool {
	if e == nil {
		return false
	}
	return e.err == target
}

func (e *errorChain) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.next
}
