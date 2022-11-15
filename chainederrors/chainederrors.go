package chainederrors

import "fmt"

type errorChain struct {
	msg  string
	err  error
	next error
}

// Wrap creates or adds err to an error chain returning
// a new error where err wraps the wrappedErr.
//
// The error chain allows errors.Is() to work on any
// nested errors that were wrapped using this method.
func Wrap(err, wrappedErr error) error {
	if err == nil {
		return wrappedErr
	}

	return &errorChain{
		msg:  fmt.Sprintf("%s : %s", err, wrappedErr),
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
