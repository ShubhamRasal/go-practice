package main


import (
	"fmt"

)


type Animal interface{
	Speak() string
}



func MakeAnimalSpeak( animal Animal){

	msg := animal.Speak()
	fmt.Println(msg)
}



type Dog struct {
	Name string
}

func (d Dog) Speak() string{
	return "Woofff!!!"
}


type Cat struct {
        Name string
}

func (c Cat) Speak() string{
        return "Meowww!!!"
}



func main(){

	fmt.Println("welcome to interface-3")
	animals := []Animal{
		
		Dog{Name: "tommy"}, 
		Cat{Name :"billy" },
	}	

	for _, animal := range animals{
		MakeAnimalSpeak(animal)
	}

}





