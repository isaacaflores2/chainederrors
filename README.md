# chainederrors
Go library which supports error chains so `errors.Is()` works with nested errors. 
The `Wrap` method is designed to replace `fmt.Errorf` and improve the limitation where the original error wrapped by `fmt.Errorf` is hidden. 

```go
var (
    errGetUserInfo = errors.New("get user info")
    errReadDocument = errors.New("read document")
    errNotFound = errors.New("document not found")
)

// wrap errors using fmt
nestedErr := fmt.Errorf("%w : %s", errReadDocument, errNotFound)
err := fmt.Errorf("%w : %s", errGetUserInfo, nestedErr)

// original error is hidden
fmt.Printf("errors.Is(err, errNotFound) = %t \n", errors.Is(err, errNotFound)) // false

// wrap errors using chainederrors
nestedErr = chainederrors.Wrap(errReadDocument, errNotFound)
err = chainederrors.Wrap(errGetUserInfo, nestedErr)

// original error is now nested and "checkable"
fmt.Printf("errors.Is(err, errNotFound) = %t \n", errors.Is(err, errNotFound)) // true

// nested error String() format
fmt.Printf("err.Error() = %s \n", err.Error()) // get user info : read document : document not found
```