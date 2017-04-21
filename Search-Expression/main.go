/*

When the program is executed,
the file in which the full path is given is read and the expression is searched.
The offset value of the point is printed on the screen.

*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Printf("%s file-name string", os.Args[0])
}

func main() {
	if len(os.Args) != 3 { // check given two arguman
		usage()
		return
	}

	b, err := ioutil.ReadFile(os.Args[1]) // Get file content
	if err != nil {
		panic(err)
	}

	n := bytes.Index(b, []byte(os.Args[2])) // n = found offset of needle ( Find file content)

	if n != -1 {
		fmt.Printf("\"%s\" expression %d found on this address.\n", os.Args[2], n)
	} else {
		fmt.Printf("\"%s\" expression could not found anywhere.\n", os.Args[2])
	}
}
