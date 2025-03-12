package utilities

// Must is a helper function that takes a value of any type and an error.
// If the error is nil, it returns the value; if the error is non-nil, it panics.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func IsStringEmpty(value string) bool {
	return value == "" || len(value) == 0
}

//func MustError[T any](value T, err responses.ResponseError) T {
//	if err != nil {
//		panic(err)
//	}
//	return value
//}
// change something
