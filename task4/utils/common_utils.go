package utils

import "fmt"

func AppendError(exitError, newError error) error {
	if newError == nil {
		return exitError
	}

	if exitError == nil {
		return newError
	} else {
		return fmt.Errorf("%v,%w", exitError, newError)
	}
}
