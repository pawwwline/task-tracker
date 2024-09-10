package e

import (
	"fmt"
)

func WrapError(msg string, err error) error {
	if err == nil {
		return nil
	}
	//fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	return fmt.Errorf("%s: %w", msg, err)
}
