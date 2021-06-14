package main

import "fmt"

func main() {
	const chance = 1
	birds_age := map[string]int64{
		"parrot":  11,
		"bird":    8,
		"chicken": 12,
	}
	if chance == 1 {
		fmt.Println(birds_age)
	} else if chance == 2 {
		fmt.Println("bye")

	} else {
		fmt.Printf("wow!")
	}
	b := 5
	switch {
	case b == 5:
		fmt.Println("5")
		// break
		fallthrough
	case b > 5:
		fmt.Println("bigger than 5")
		fallthrough
	case b <= 5:
		fmt.Println("5 , maybe smaller than")
		fallthrough
	default:
		fmt.Println("num not found")
	}
}
