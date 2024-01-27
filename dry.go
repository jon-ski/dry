package dry

import "fmt"

func AsError(r interface{}) error {
	// No error, return nil
	if r == nil {
		return nil
	}

	// Already an error, return it
	if e, ok := r.(error); ok {
		return e
	}

	// Not an error, return it as a string
	return fmt.Errorf("%v", r)
}
