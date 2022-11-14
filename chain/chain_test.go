package chain_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/isaacaflores2/chainederrors/chain"
)

func TestErrorChain(t *testing.T) {
	var (
		errGetItem  = errors.New("get item")
		errMakeRequest = errors.New("make network request")
		errTimeout = errors.New("request timeout")
		errChain = chain.Wrap(errMakeRequest, errTimeout)
	)
	testCases := []struct {
		name        string
		err         error
		targetErr error
	}{
		{
			name:        "Check for outer error",
			err:         chain.Wrap(errGetItem, errMakeRequest),
			targetErr: errGetItem,
		},
		{
			name:        "Check for inner chained error",
			err:         chain.Wrap(errGetItem, errMakeRequest),
			targetErr: errMakeRequest,
		},
		{
			name:        "Check for error in existing chain",
			err:         chain.Wrap(errGetItem, errChain),
			targetErr: errTimeout,
		},
		{
			name:        "Outer error is nil",
			err:         chain.Wrap(nil, errTimeout),
			targetErr: errTimeout,
		},
		{
			name:        "Inner error is nil",
			err:         chain.Wrap(nil, errTimeout),
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
		})
	}
}
