// ls.go
package builtins

import (
	"fmt"
	"io"
	"io/ioutil"
)

func ListDirectory(w io.Writer, args ...string) error {
	var dir string
	if len(args) > 0 {
		dir = args[0]
	} else {
		dir = "."
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
	return nil
}
