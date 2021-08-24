package gokoans

import (
	"fmt"
	"strings"
)

// A string in Go is a slice of bytes.
// Resources:
// https://golangbot.com/strings/
// https://medium.com/rungo/string-data-type-in-go-8af2b639478
// https://blog.golang.org/strings

func aboutStrings() {
	assert("a"+__string__ == "abc") // string concatenation need not be difficult
	assert(len("abc") == __int__)   // and bounds are thoroughly checked

	assert("abc"[0] == __byte__) // their contents are reminiscent of C

	assert("smith"[2:] == __string__)  // slicing may omit the end point
	assert("smith"[:4] == __string__)  // or the beginning
	assert("smith"[2:4] == __string__) // or neither
	assert("smith"[:] == __string__)   // or both

	assert("smith" == __string__) // they can be compared directly
	assert("smith" < __string__)  // i suppose maybe this could be useful.. someday

	bytes := []byte{'a', 'b', 'c'}
	assert(string(bytes) == __string__) // strings can be created from byte-slices

	bytes[0] = 'z'
	assert(string(bytes) == __string__) // byte-slices can be mutated, although strings cannot

	assert(fmt.Sprintf("hello %s", __string__) == "hello world")               // our old friend sprintf returns
	assert(fmt.Sprintf("hello \"%s\"", "world") == __string__)                 // quoting is familiar
	assert(fmt.Sprintf("hello %q", "world") == __string__)                     // although it can be done more easily
	assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == __string__) // "the root of all evil" is actually a misquotation, by the way

	// For the next koans, see the
	assert(strings.Contains("Yale", "ale") == __bool__)                                     // You can't spell Yale without ale.
	assert(strings.Count("Yale School of Management", "o") == __int__)                      // And you need some o's for SOM.
	assert(strings.HasPrefix("MGT656", "MGT") == __bool__)                                  // Management is the basis of this class.
	assert(strings.Index("Beyond Beef", "Beef") == __int__)                                 // Where's the beef?
	assert(strings.Join([]string{"business", "society", "chocolate"}, "and") == __string__) // Our purpose
	assert(strings.Split("business❤️society", "❤️")[0] == __string__)
	assert(strings.Split("business❤️society", "❤️")[0] == __string__)
}
