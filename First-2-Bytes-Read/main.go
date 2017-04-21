/*

When the program is run,
the first 2 bytes of a file to be given the full path in the code
are read and printed on the screen.

*/

package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("%s file-name\n", os.Args[0]) // If not given parameter, this'll show.
}

func main() {
	if len(os.Args) != 2 { // Check if arguman exists pass
		usage()
		return
	}

	f, err := os.Open(os.Args[1]) //os.Args[1] filename
	if err != nil {
		panic(err)
	}

	defer f.Close() // close file handler

	data := make([]byte, 2) // we create array(size 2)

	_, err = f.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("First 2 bytes : %s\n", data)
}
