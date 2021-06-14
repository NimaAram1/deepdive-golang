package main

import (
	"fmt"
	"sync"
	"time"
)

const attention = `Don't Repeat Yourself!
				   DRY`

var wg = sync.WaitGroup{}
var mark int8 = 0

func main(){

	var msg = "Hello"
	wg.Add(4)
	go func(msg string){
		fmt.Println(msg)
		wg.Done()
	}(msg)
	go increase()
	go hardProcess()
	go hardestProcess()
	msg = "Goodbye"
	defer fmt.Println("Hello!")
	wg.Wait()
	defer fmt.Println("Bye!")
	
	// we can use this without wait group
	
}

func increase(){
	mark++
	fmt.Println(mark)
	wg.Done()
	
}

func hardProcess(){
	for i := 0; i<=10; i++{
		time.Sleep(time.Second / 3)
		fmt.Println("Processed ", i)
	}
	wg.Done()
	
}

func hardestProcess(){
	for i := 0; i<=20; i++{
		time.Sleep(time.Second / 2)
		fmt.Println("Bad ", i)
	}
	wg.Done()
}