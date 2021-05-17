package main

import (
	"fmt"
)

func main(){
	var a int = 53;
	var b *int = &a;

	fmt.Print(a,b,"\n")
	fmt.Println("new: ",*b)

	var j int = 5;
	k := &j;

	fmt.Printf("%v %p \n",k,k)

	type logo struct{
		foo int
	}

	var sm *logo
	sm = new(logo)
	sm.foo = 5;
	fmt.Println(sm.foo)



}
