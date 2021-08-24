package gokoans

// "defer" is used to do something like clean up after yourself
// or perform when a task when a function exits.
// See https://gobyexample.com/defer

func aboutDefer() {
	var acc int

	increment := func(value int) {
		acc += value
	}

	decrement := func(value int) {
		acc -= value
	}

	func() {
		acc = 0
		defer increment(1)
	}()

	assert(acc == __int__) // defer function will be execute after main function body

	func() {
		acc = 0
		defer increment(5)
		defer decrement(3)
	}()

	assert(acc == __int__) // list of functions also allowed

}
