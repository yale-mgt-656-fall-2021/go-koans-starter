package gokoans

import (
	"fmt"
	"io/ioutil"

	// "os"
	"path"
	"runtime"
	"strings"
	"testing"
)

const (
	__string__       string  = "impossibly lame value"
	__int__          int     = -1
	__positive_int__ int     = 42
	__byte__         byte    = 255
	__bool__         bool    = false // ugh
	__boolean__      bool    = true  // oh well
	__float32__      float32 = -1.0
	__delete_me__    bool    = false
)

var passedTests int
var totalTests int

func TestKoans(t *testing.T) {
	aboutBasics()
	aboutStrings()
	aboutArrays()
	aboutSlices()
	aboutControlFlow()
	aboutEnumeration()
	aboutAnonymousFunctions()
	aboutVariadicFunctions()
	aboutDefer()
	aboutFiles()
	aboutInterfaces()
	aboutCommonInterfaces()
	aboutMaps()
	aboutPointers()
	aboutStructs()

	if totalTests == passedTests {
		fmt.Printf("\n%c[32;1m%d/%d tests passed. Woohoo! %c[0m\n\n", 27, passedTests, totalTests, 27)
	} else {
		fmt.Printf("\n%c[32;1m%d/%d tests passed.%c[0m\n\n", 27, passedTests, totalTests, 27)
	}
}

func assert(o bool) {
	totalTests = totalTests + 1
	if !o {
		fmt.Printf("\n%c[31;3mTEST FAILED\n%s%c[0m\n\n", 27, getRecentLine(), 27)
		// os.Exit(1)
	} else {
		//fmt.Printf("\n%c[32;3mTEST Passed%c\n%c[31;3m%s%c[0m\n\n",27,27, 27, getRecentLine(), 27)
		passedTests = passedTests + 1
	}
}

func getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
