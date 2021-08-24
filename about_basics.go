package gokoans

// In Go, as well as many other programming languages, we declare a data type 
// as we initialize variables. For example a word or sentence (ex: hello world) 
// is a "string", a whole number (ex: 6) is an "int", and a decimal (ex: 1.2)
// is a "float."

// More resources: 
// https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go
// https://go101.org/article/type-system-overview.html

func aboutBasics() {
	assert(__bool__ == true)  // what is truth?
	assert(__bool__ != false) // what is not false?

	var i int = __int__
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := __int__ //short assignment can be used, as well
	assert(k == 1.0000000000000000000000000000000000000)

	assert(5%2 == __int__)
	assert(5*2 == __int__)
	assert(5^2 == __int__)

	// In Go, we can initialize variables without giving it value.
	// Each data type has a "zero-value" that the variable defaults to in these cases.
	// Learn about them here: https://golang.org/ref/spec#The_zero_value
	// Run this script to see the zero values: https://tour.golang.org/basics/12

	var x int
	assert(x == __int__) // zero values are valued in Go

	var f float32
	assert(f == __float32__) // for types of all types

	var s string
	assert(s == __string__) // both typical or atypical types

	// This is an anonymous structure. A structure represents a collection of fields.
	// See: https://golangbot.com/structs/  and https://www.callicoder.com/golang-structs/

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == __int__)     // and types within composite types
	assert(c.f == __float32__) // which match the other types
	assert(c.s == __string__)  // in a typical way
}
