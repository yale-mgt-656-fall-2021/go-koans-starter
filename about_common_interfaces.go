package gokoans

import "bytes"

// There are many interfaces defined in the standard library.
// E.g.
// Reader - https://golang.org/pkg/io/#Reader
// ReadWriter - https://golang.org/pkg/io/#ReadWriter
// Writer - https://golang.org/pkg/io/#Writer
// These are commonly encountered and useful.

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
