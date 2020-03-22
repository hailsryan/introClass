package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

// host on port 8080
// build three http handlers

// one with the route /wipe
// reset the counter to 0
// send back a 200 status code with no body

// one with the route /add
// read a number from the request and add to cumulative sum
// send back a 200 status code, but no body
// note: the number will come in as a string, so once reading
// the body, you will need to call strconv.AtoI() to transform
// the string into an integer

// one with the route /total
// send back a 200 status code and the body should
// just be a string that is the number for the total count
// you will need to transform your integer to a string using
// strconv.AtoI().
