package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	names := [...]string{"nima", "majid"}
	var grades [5]int
	grades[0] = 10
	grades[1] = 15
	grades[2] = 20
	fmt.Printf("%v\n %v\n", numbers, names[0])
	fmt.Printf("%v", grades)

	k := []int{1, 2, 3, 4, 5}

	fmt.Printf("\nLen : %v\n", len(k))
	fmt.Printf("Cap : %v", cap(k))

	{

		my_numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println("\n")
		fmt.Printf("%v \n", my_numbers[:])
		fmt.Printf("%v \n", my_numbers[3:])
		fmt.Printf("%v \n", my_numbers[:5])
		fmt.Printf("%v \n", my_numbers[2:4])

		h := make([]int, 5, 10)
		// diffrent cap

		b := []string{"ching", "chong", "chang"}
		fmt.Printf("\n%v", b)
		b = append(b, "chochang")
		var p = append(b[:1], b[2:]...)
		fmt.Printf("\n%v", b)
		fmt.Println(h, b, p)

	}

}
