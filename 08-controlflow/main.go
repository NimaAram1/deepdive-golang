package main

import (
	//"net/http"
	"fmt"
	"log"
)

func main() {

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	//	w.Write([]byte("Magic"))
	//})
	//
	//err := http.ListenAndServe(":8080",nil)
	//
	//if err != nil{
	//	panic(err.Error())
	//}


	fmt.Println("Hello!")
	defer fmt.Println("I dont' think so")
	//panic("errors") // stops orinery flow
	fmt.Println("End!")


	fmt.Println("START")
	run_app()
	fmt.Println("END")



}

func run_app(){

	defer func(){
		if err := recover(); err != nil{
			log.Println(err)
		}
	}()
	panic("stopping flow")
	fmt.Println("panicked")

}