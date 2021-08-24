package gokoans

import "io/ioutil"
import "strings"

// Learn about reading files: https://gobyexample.com/reading-files
// Learn about strings.Split (and run the example!): https://gobyexample.com/reading-files

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	assert(lines[5] == __string__) // handling files is too trivial
}
