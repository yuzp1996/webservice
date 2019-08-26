package customhandler

import (
	"log"
	"net/http"
)



func DogandCat(w http.ResponseWriter, req *http.Request){
	defer func() {
		if r := recover(); r != nil{
			log.Printf("Panic err happend bug Recovered err:{%v}", r)
		}
	}()
	newDog := Dog{
		Name:"DOG",
	}
	newCat := Cat{
		Name:"CAT",
	}

	newCat.GetChild()

	childCat1 := Cat{
		Name:"childCat1mimi",
	}
	childCat2 := Cat{
		Name:"childCat2miomio",
	}
	childernCat := []Cat{
		childCat1,
		childCat2,
	}

	newCat.GetChild(childernCat...)

	newCat.say(newCat)
	newDog.say(newDog)

	newCat.say(newDog)
	newDog.say(newCat)

}


type animal interface{
	say(animal)
	name() string
}


type Dog struct {
	Name string
}

func(dog Dog)say(object animal){
	// I want to konw if the object is dog
	shouldadog := object.(Dog)

	log.Printf("wangwang... your name is %v", shouldadog.name())
}
func (dog Dog)name()string  {
	return dog.Name
}




type Cat struct {
	Name string
}
func(cat Cat)say(object animal){
	shouldacat := object.(Cat)
	log.Printf("miaomiao...  your name is %v",shouldacat.name())
}
func (cat Cat)name()string  {
	return cat.Name
}

// I can recive null param
func(cat Cat)GetChild(cats ...Cat){
	var childname string
	for _, catname := range cats{
		childname += catname.Name+" "
	}
	log.Printf("cat has child cat and name is %v\n",childname)

}






