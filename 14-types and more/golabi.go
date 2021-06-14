package main

import (
	"fmt"
	"strings"
	"runtime"
	"time"
)

type Role string

type ErrorMigration struct{
	createdAt time.Time
	body string
}

type User struct{
	username string
	age int8
}

type UserBalancer interface{
	NameWithAge() string
}

func main(){
	fmt.Printf("version %v \n",runtime.Version())
	role1 := Role("khobe?")
	fmt.Println(role1.ToUppercase())

	majid := &User{
		username: "nima",
		age: 17,
	}
	fmt.Println(HelloToUser(majid))

	log_app := runApp()
	fmt.Println(log_app)

	go printText("hello")
	printTextPlus("world")
	
}

func (r Role) ToUppercase() string {
	return strings.ToUpper(string(r))
}

func (u *User) NameWithAge() string{
	return fmt.Sprintf("%v age: %v",u.username, u.age)
}

func HelloToUser(UserB UserBalancer) string{
	return fmt.Sprintf("say: %v", UserB.NameWithAge())
} 

func (err *ErrorMigration) Error() string{
	return fmt.Sprintf("At %v Error : %v", err.createdAt, err.body)
}

func runApp() error{
	newError := &ErrorMigration{
		createdAt: time.Now(),
		body: "bad request",
	}
	return newError
}

func printText(text string){
	for i := 0; i < 5; i++{
		fmt.Println(text)
		time.Sleep(time.Second / 2)				
	}
}

func printTextPlus(text string){
	for i := 0; i < 5; i++{
		fmt.Println(text)
		time.Sleep(time.Second / 4)				
	}
}