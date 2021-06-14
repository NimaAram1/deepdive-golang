package main

import (
	"fmt"
	"reflect"
)


func main(){
	names := make(map[string]string)
	names = map[string]string{
		"karafs": "hayejan-zade",
		"hollo": "bepar to galo",
	}
	fmt.Println(names)
	fmt.Println(names["hollo"])
	fmt.Println(names["karafs"])
	names["hollo"] = "napar to galo"
	names["new"] = "new"
	fmt.Println(names["hollo"])
	fmt.Println(names)
	delete(names, "new")
	fmt.Println(names)

//	get length of map
	fmt.Println(len(names))

// get object with status
	pop, status := names["holl"]
	pop2, status2 := names["hollo"]
	fmt.Println(status,pop)
	fmt.Println(status2,pop2)

// create struct (OOP)
	type Fruit struct {
		init string `required max: "100"`
	}
	type Kahoo struct {
		//Fruit
		id int
		model string
		year_created []int64
	}
	//kahoo1 := Kahoo{id: 2,model: "tracked",year_created: []int64{2005,2010}}
	kahoo1 := Kahoo{2,"tracked",[]int64{2005,2010}}
	fmt.Println(kahoo1)
	fmt.Println(kahoo1.id)
	fmt.Println(kahoo1.year_created)
	kahoo2 := struct{model string; id int64}{"majid",20}
	fmt.Println(kahoo2.model)
	fmt.Println(kahoo2.id)
	kahoo2.id = 30
	fmt.Println(kahoo2.id)
	FruitStruct := reflect.TypeOf(Fruit{})
	field, _ := FruitStruct.FieldByName("init")
	fmt.Println(field.Tag)

	// end
}
