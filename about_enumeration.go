package gokoans

// When we "enumerate" a collection, we count off or list each element
// in the set. In Go, we use "range" to iterate over each element in a
// data structure, such as an array.
// Examples of range:
// https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/
// https://gobyexample.com/range
// https://tour.golang.org/flowcontrol/1

func aboutEnumeration() {
	{
		var concatenated string
		var total int

		strings := []string{"hello", " world", "!"}
		for i, v := range strings {
			total += i
			concatenated += v
		}

		assert(concatenated == __string__) // for loops have a modern variation
		assert(total == __int__)           // which offers both a value and an index
	}

	{
		var totalLength int

		strings := []string{"hello", " world", "!"}
		for _, v := range strings {
			totalLength += len(v)
		}

		assert(totalLength == __int__) // although we may omit either value
	}
}
