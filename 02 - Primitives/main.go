package main

import "fmt"

func main() {
	n := 2 == 2
	m := 1 == 2
	yy := 9
	kk := 7

	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T\n", m, m)
	fmt.Println(yy > kk)
	fmt.Println(yy < kk)
	fmt.Println(yy != kk)
	fmt.Println(yy <= kk)
	fmt.Println(yy >= kk)
	fmt.Println((2 == 2 && 3 == 3))
	fmt.Println((2 == 3 || 3 == 3))

	a := 5
	b := 10

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	var d int = 20
	var j int8 = 4

	fmt.Println(d + int(j))

	// binary primitives

	num1 := 30 // 11110
	num2 := 20 // 10100

	x := 16 // 2^4

	fmt.Println(num1 & num2)  // 10100
	fmt.Println(num1 | num2)  // 11110
	fmt.Println(num1 ^ num2)  // 01010
	fmt.Println(num1 &^ num2) // 10101
	fmt.Println(x << 5)       // 2^4 * 2^5 = 2^9
	fmt.Println(x >> 5)       // 2^4 / 2^5 = 2^-1

	var tt complex64 = 5 + 2i
	var bb complex128 = 20 + 220i
	fmt.Printf("%v %T\n", imag(tt), imag(tt))
	fmt.Printf("%v %T\n", real(tt), real(tt))
	fmt.Printf("%v %T\n", tt, tt)
	fmt.Printf("%v %T\n", imag(bb), imag(bb))
	fmt.Printf("%v %T\n", real(bb), real(bb))
	fmt.Printf("%v %T\n", bb, bb)

	// string

	my_name := "nima"

	fmt.Printf("%v %T\n", my_name[1], my_name[1])
	fmt.Printf("%v %T\n", string(my_name[1]), string(my_name[1]))

	var into_the_byte = []byte(my_name)

	var my_alter_name rune = 'i'
	fmt.Printf("%v %T\n", my_alter_name, my_alter_name)
	fmt.Println(into_the_byte)

}
