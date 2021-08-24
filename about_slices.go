package gokoans

// A slice is like an array, except that it can grow or shrink in size
// It also has a slice operator that retrieves a subset of the slice data type
// Resources:
// https://www.geeksforgeeks.org/slices-in-golang/
// https://gobyexample.com/slices
// https://tour.golang.org/moretypes/7

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	assert(fruits[0] == __string__) // slices seem like arrays
	assert(len(fruits) == __int__)  // in nearly all respects

	tastyFruits := fruits[1:3]           // we can even slice slices
	assert(tastyFruits[0] == __string__) // slices of slices also share the underlying data

	carSlots := []string{"good one", "good one", "lemon"}
	assert(cap(carSlots) == __int__) // the capacity is initially the length

	carSlots = append(carSlots, "good one!")
	assert(len(carSlots) == __int__) // slices can be extended with append(), much like realloc in C
	assert(cap(carSlots) == __int__) // but with better optimizations

	carSlots = append(carSlots, "another good one!?", "yet another!", "good luck!")

	assert(len(carSlots) == __int__) // append() can take N arguments to append to the slice
	assert(cap(carSlots) == __int__) // the capacity optimizations have a guessable algorithm
}
