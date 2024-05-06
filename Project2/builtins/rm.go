// rm.go
package builtins

import (
	"fmt"
	"os"
)

func Remove(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("rm: missing operand")
	}
	for _, file := range args {
		err := os.Remove(file)
		if err != nil {
			return err
		}
	}
	return nil
}
