package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("\a")

	for j, l := 0, 0; j <= 5; l, j = l+1, j+2 {
		fmt.Printf("%v %v", l, j)
	}

	for k := 0; k <= 100; k++ {
		if k%7 == 0 {
			fmt.Println(k)
		}
	}

	bixCoeitIndex := 0
	for {
		fmt.Println(bixCoeitIndex)
		bixCoeitIndex++
		if bixCoeitIndex == 3 {
			break
		}
	}

	numbersList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for g, h := range numbersList {
		fmt.Println(g, h)
	}

	mystring := "Summon more codes"

LoopSummon:
	for b, m := range mystring {
		fmt.Println(b, string(m))
		if string(m) == "m" {
			break LoopSummon
		}
	}

}
