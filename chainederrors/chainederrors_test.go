package chainederrors_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/isaacaflores2/chainederrors/chainederrors"
)

func TestErrorChain(t *testing.T) {
	var (
		errGetItem  = errors.New("get item")
		errMakeRequest = errors.New("make network request")
		errTimeout = errors.New("request timeout")
		errChain = chainederrors.Wrap(errMakeRequest, errTimeout)
	)
	testCases := []struct {
		name        string
		err         error
		targetErr error
	}{
		{
			name:        "Check for outer error",
			err:         chainederrors.Wrap(errGetItem, errMakeRequest),
			targetErr: errGetItem,
		},
		{
			name:        "Check for inner chained error",
			err:         chainederrors.Wrap(errGetItem, errMakeRequest),
			targetErr: errMakeRequest,
		},
		{
			name:        "Check for error in existing chain",
			err:         chainederrors.Wrap(errGetItem, errChain),
			targetErr: errTimeout,
		},
		{
			name:        "Outer error is nil",
			err:         chainederrors.Wrap(nil, errTimeout),
			targetErr: errTimeout,
		},
		{
			name:        "Inner error is nil",
			err:         chainederrors.Wrap(nil, errTimeout),
			targetErr: errTimeout,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			x, ok := tc.err.(interface{ Is(error) bool })
			
			if ok && x.Is(tc.targetErr) {
				fmt.Println("is target!")
			}	

			if isErr := errors.Is(tc.err, tc.targetErr); isErr == false {
				t.Errorf("errors.Is() got = %t, want true", isErr)
			}

			if len(tc.err.Error()) == 0 {
				t.Errorf("error() got empty string, want formatted error")
			}
		})
	}
}

func ExampleWrap() {
	var (
		errGetUserInfo = errors.New("get user info")
		errReadDocument = errors.New("read document")
		errNotFound = errors.New("document not found")
	)

	// wrap errors using fmt
	nestedErr := fmt.Errorf("%w : %s", errReadDocument, errNotFound)
	err := fmt.Errorf("%w : %s", errGetUserInfo, nestedErr)
	fmt.Printf("errors.Is(err, errNotFound) = %t \n", errors.Is(err, errNotFound)) // false

	// wrap errors using chainederrors
	nestedErr = chainederrors.Wrap(errReadDocument, errNotFound)
	err = chainederrors.Wrap(errGetUserInfo, nestedErr)
	fmt.Printf("errors.Is(err, errNotFound) = %t \n", errors.Is(err, errNotFound)) // true
	
	fmt.Printf("err.Error() = %s \n", err.Error()) // get user info : read document : document not found
}
