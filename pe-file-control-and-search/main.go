/*

When the program is run,
the search is executed to include all the files (except subfolders) in a folder
that will be given as a program argument. However, the search should only take place in Portable Executable (PE) format files.
Other files should be skipped and user information should be given.
It is sufficient to check whether the first 2 bytes are "MZ" to determine whether a file is in Pe format.


*/

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Printf("%s directory string", os.Args[0])
}

func isPE(b []byte) bool {
	// if len(b) < 0x120 { //file size control
	// 	return false
	// }

	if string(b[:2]) != "MZ" { // first 2 bytes control.
		return false
	}
	//3c de pe header nerde başladiği var, onu ordan okuyup peOffsete attım diğer kontrolleri ona göre yaptım.
	peOffset := binary.LittleEndian.Uint32(b[0x3c : 0x3c+4]) //PE signature at offset stored in MZ header at 0x3C.

	// if peOffset > uint32(len(b)+5) {
	// 	return false
	// }

	// if string(b[peOffset:peOffset+2]) != "PE" { // PE control
	// 	return false
	// }

	if (b[peOffset+4] == 0x4c && b[peOffset+5] == 0x01) || (b[peOffset+4] == 0x64 && b[peOffset+5] == 0x86) { // 32bit or 64bit PE control
		return true
	}

	return false
}

func scan(path string, f os.FileInfo, err error) error { //This func  scan arguments with the given file path.
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	if !isPE(b) {
		fmt.Printf("\"%s\": is not PE file.\n", path)
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

	err := filepath.Walk(os.Args[1], scan)
	if err != nil {
		panic(err)
	}
}
