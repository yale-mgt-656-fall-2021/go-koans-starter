package gokoans

// An "interface" in Go is a named collection of
// functions (methods). A struct is said to satisfy
// an interface if it has the function that the
// interface needs.
// See
// * https://www.alexedwards.net/blog/interfaces-explained
// * https://gobyexample.com/interfaces
// * https://golang.org/doc/effective_go.html#interfaces
// * https://www.golangprograms.com/go-language/interface.html


func aboutInterfaces() {
	alice := new(teacher) // bob is a kind of *human, see below
	bob := new(student)   // rspec is a kind of *program

	runTwice(alice) // alice fits the profile for an 'attendee'
	runTwice(bob)   // bob also fits the profile for an 'attendee'

	assert(alice.paychecks == __int__) // alice is affected by attending in her own unique way
	assert(bob.credits == __int__)     // bob gets credits instead of paychecks, thanks to interfaces
}

// abstract interface and function that requires it

type attendee interface {
	attendClass()
}

func runTwice(r attendee) {
	r.attendClass()
	r.attendClass()
}

// concrete type implementing the interface

// when a variable is declared without an explicit value,
// a "zero-value" is assigned. see: https://golang.org/ref/spec#The_zero_value
type teacher struct {
	paychecks int
}

func (h *teacher) attendClass() {
	h.paychecks++
}

// another concrete type implementing the interface

type student struct {
	credits int
}

func (p *student) attendClass() {
	p.credits++
}
