// mkdir.go
package builtins

import (
	"fmt"
	"os"
)

func MakeDirectory(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("mkdir: missing operand")
	}
	for _, dir := range args {
		err := os.Mkdir(dir, 0755) // creates directory with rwxr-xr-x permissions
		if err != nil {
			return err
		}
	}
	return nil
}
