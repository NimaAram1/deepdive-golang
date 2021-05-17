package main

import (
	"fmt"
)

func main(){
	const CHECK_OBJECTS = true;
	PersianLanguage := map[string]string{
		"f1": "karafs",
		"f2": "limoo",
		"f3": "hollo",
		"f4": "sib",
	}
	for first, value := range PersianLanguage{
		fmt.Print(value + " \n ")
	}
}