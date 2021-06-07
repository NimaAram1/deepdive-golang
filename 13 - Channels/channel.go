package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){
	// when we say func m(v *int){} we want a integer varibale address (&v)
	var b int = 1
	numberPlusOne(&b)
	ch := make(chan int, 3)
	wg.Add(2)
	// go func(){
	// 	resault := <- ch
	// 	fmt.Println(resault)
	// 	resault = <- ch
	// 	fmt.Println(resault)
	// 	resault = <- ch
	// 	fmt.Println(resault)
	// 	wg.Done()

	// }()
	go func(c chan<- int){
		c <- 50
		c <- 25
		c <- 100
		close(c)
		wg.Done()
	}(ch)
	go func(ch <-chan int ){
		for {
			if i, ok := <- ch; ok{
				fmt.Println(i, ok)
			}else{
				break
			}
		}
		wg.Done()
	}(ch)
	// go func(){
	// 	for o := range ch{
	// 		fmt.Println(o)
	// 	}
	// 	wg.Done()
	// }()
	wg.Wait()
	
}

func numberPlusOne(num *int){
	 *num = *num + 1
	 fmt.Println(*num)
}