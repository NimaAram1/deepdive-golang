package main

import (
	"fmt"
)
func main(){
	var root_path string = "/usr/locals";
	showErrorMessage()
	getAndCheckData("abbas boazar",1)
	changeRoot(&root_path)
	fmt.Println(root_path)
	multiple(56,89,12)
}

func showErrorMessage(){
	fmt.Println("an error found in your potato pc")
}

func getAndCheckData(username string, index int){
	fmt.Printf("User: %v with Id: %v \n",username,index)
}
func changeRoot(path *string){
	*path = "/usr/hashed"
}
func multiple(values ...int) *int {
	resault := 1
	for _, val := range values {
		resault = val * resault
	}
	fmt.Println(resault)
	return &resault
}
