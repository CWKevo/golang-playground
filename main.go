package main

import (
	"fmt"
	"os"
	"strings"
	"runtime"
	"github.com/codingsince1985/checksum"
)


func main() {
	var newline string = "\n"; if runtime.GOOS == "windows" { newline = "\r\n" }

	var to_strip []string = []string{newline, "\t", "\n", " ", ";"}

	files, err := os.ReadDir(".")
	if err != nil { panic(err) }

	for i := 0; i < len(files); i++ {
		fileinfo, err := files[i].Info()
		if err != nil { panic(err) }
		if fileinfo.IsDir() { continue }

		filename := fileinfo.Name()

		contents, err := os.ReadFile(filename)
		if err != nil { panic(err) }

		var clean string = string(contents)

		for i := 0; i < len(to_strip); i++ {
			clean = strings.ReplaceAll(clean, to_strip[i], "")
		}

		chs, _ := checksum.SHA256sum(filename)
		fmt.Println(filename, chs)
	}
}
