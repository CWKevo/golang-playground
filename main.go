package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"runtime"
	"strings"
)



func main() {
	var newline string = "\n"; if runtime.GOOS == "windows" { newline = "\r\n" }
	var to_strip []string = []string{newline, "\t", " ", ";"}


	files, err := os.ReadDir(".")
	if err != nil { panic(err) }

	for i := 0; i < len(files); i++ {
		fileinfo, err := files[i].Info()
		if err != nil { panic(err) }
		if fileinfo.IsDir() { continue }


		filename := fileinfo.Name()

		contents, err := os.ReadFile(filename)
		if err != nil { panic(err) }


		clean := string(contents)
		for i := 0; i < len(to_strip); i++ {
			clean = strings.ReplaceAll(clean, to_strip[i], "")
		}


		chs := sha256.Sum256([]byte(clean))
		fmt.Printf("%s - %x\n", filename, chs)
	}
}
