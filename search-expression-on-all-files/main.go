/*

When the program is run,
the search is executed to include all the files (except subfolders) in a folder
that will be given as a program argument. The end user is informed.


*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Printf("%s directory string", os.Args[0])
}

func scan(path string, f os.FileInfo, err error) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	n := bytes.Index(b, []byte(os.Args[2]))

	if n != -1 {
		fmt.Printf("\"%s\": found (%d on this address).\n", path, n)
	} else {
		fmt.Printf("\"%s\": could not found.\n", path)
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		usage()
		return
	}

	err := filepath.Walk(os.Args[1], scan) // filepath.walks  take a func as an arguman. The func execute for each file.
	if err != nil {
		panic(err)
	}
}
