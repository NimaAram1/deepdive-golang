package main

import (
	"fmt"
	"strconv"
)

// unused outside of main()
// another type of variables
var (
	old_age   = 70
	young_age = 18
	// we cand declare type
	mid_age   int64 = 40
	mid_age_n int   = 40
)

func main() {
	// we can declare variables in diffrent types

	// type 1
	var name = "go"

	// type 2
	var prefix string = "main"

	// type 3
	big_num := 9999

	// all of variables in main() must be used
	fmt.Printf(" The Values %v %v %v ", name, prefix, big_num)
	fmt.Printf(" The Values %T %T %T \n ", name, prefix, big_num)
	// string formatting with %v and %T and Printf

	var h float32 = 25.666
	var x int
	var y string
	// converting type
	x = int(h)
	// y = string(x) this code give error
	y = strconv.Itoa(x)
	fmt.Println(h, x, y)

}
