# Go Koans

## Overview

In this assignment, you will complete a series of [koans](https://en.wikipedia.org/wiki/Kōan) in Go. In this repository, you'll notice a series of `about_something` files. Each of these files has a series of koans (or tests) that currently fail. Your mission is to complete the koans, learning about the topic assocaited with each file along the way.

While completing these koans, you may only change the placeholders, which are written in the files as data types surrounded by underscores. The placeholders you will see include: `__int__`, `__float32__`, `__string__`, and `__bool__`. The data types written in each placeholder are hints for what you should replace them with. _Do not_ change any other code, otherwise you may lose points on the assignment.

For example, in the `about_basics.go` file, you'll notice the following koan at the top:

```
assert(__bool__ == true)  // what is truth?
assert(__bool__ != false) // in it there is nothing false
```

To complete that koan, you'll need to replace the `__bool__` placeholders with a value that makes the tests passed. In this case, the koan is solved as follows:

```
assert(true == true)  // what is truth?
assert(true != false) // in it there is nothing false
```

## Getting started

To begin, you need to accept the GitHub Classroom invite for the
assignment, which will create a GitHub repo for you in the class
organization on GitHub (either 656 or 660). Actually, if you're
reading this, its likely you already did that. High five . Then,
you should clone the repo to your computer (or cloud9). Then, you're
ready to go.

In your terminal or PowerShell (windows) change directories to be
in the same directory as this file and then run `go test` to test
your completion of the koans. You'll start off with zero completed,
naturally. You'll edit the code in order to make the tests (koans)
pass. When they all pass you've completed the assignment.

I suggest working on the files in the following order:

- `about_basics.go`
- `about_strings.go`
- `about_arrays.go`
- `about_slices.go`
- `about_control_flow.go`
- `about_enumeration.go`
- `about_anonymous_functions.go`
- `about_variadic_functions.go`
- `about_defer.go`
- `about_files.go`
- `about_interfaces.go`
- `about_common_interfaces.go`
- `about_maps.go`
- `about_pointers.go`
- `about_structs.go`

Each file contains resources, such as links, to help you learn about the
concepts covered in that file so that you can complete the koans. If you
want a general preparation before working on these files, I suggest spending
a few minutes with the [Tour of Go](https://tour.golang.org/welcome/1).
Many of the concepts covered in thi assignment are covered in the tour.

You can test your progress by running `go test` in your terminal
at any time. You should do this often. The test script will show
you the file name and line numbers for incomplete koans, as well
as the number of koans you've completed.

Remember to save your progress along the way by adding and committing your files!
If you just worked on `about_slice.go`, you might do something like
`git status` to see which files have edits, `git diff` to see
the content of the edits, `git add about_slice.go` to add the file
to the "staging area" in preparation for commit and
`git commit -m "work on slices, woot!"` to commit your work.
(Recall, the "work on slices, woot!" part is the comment for your commit.)

## Submitting and Updating

In order to receive credit for this assignment, you must submit
your GitHub repo URL through the assignment page on the class
website. That way, we'll know which repo on GitHub is yours. You
should [git push](https://www.atlassian.com/git/tutorials/syncing/git-push)
regularly to GitHub as you complete the assignment---you can do so
right up until the deadline. If you don't push to GitHub, we won't
have your work and therefore our grading will not reflect your true
awesomeness!

## Credits

This code comes largely from [Jason Thigpen's Go Koans](https://github.com/cdarwin/go-koans).
