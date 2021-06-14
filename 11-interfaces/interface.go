package main

import "fmt"

const DEBUG = true

func main(){
	func(){
		fmt.Println("prepairing...")
	}()
	defer test_pointers()
	
	fmt.Println("hello, nice to meet you!")

	majidStar := stars{id: 1, value: 3.5}
	majid := User{name: "majid", family: "majidi", age: 20, star: majidStar}
	defer fmt.Println(majid.name)
	defer fmt.Println(majid.get_age())
	defer fmt.Println(majid.star.id)
	defer fmt.Println(majid.star.value)
	majid.star.set_value_to_max()
	defer fmt.Println(majid.star.value)
	majid.change_name("abbas boazar")
	fmt.Println(majid.name)
	
	



}



type User struct{
	name string
	family string
	age int
	star stars
}

type stars struct{
	id int64
	value float64
}

type Profile struct{
	User
	bio string
}

func (obj *User) get_age() int {
	return obj.age
}

func(obj *stars) set_value_to_max() {
	obj.value = 5
}


type UserI interface{
	change_name(n_name string) 
}

func (obj *User) change_name(n_name string) {
	obj.name = n_name
}


// some pointers here
func test_pointers(){
	var name string = "Nima"
	var name_value *string = &name
	fmt.Println(*name_value) // get name value
	fmt.Println(name_value) // get address in memory
	*name_value = "Majid" // change name value
	fmt.Println(*name_value) // get name value
}


