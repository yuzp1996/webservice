package customhandler

import (
	"log"
	"net/http"
	. "webservice/Http/server/interface"
)

func DogandCat(w http.ResponseWriter, req *http.Request){
	defer func() {
		if r := recover(); r != nil{
			log.Printf("Panic err happend bug Recovered err:{%v}", r)
		}
	}()

	//在方法中使用func作为参数
	NewDogwithfriend()



	testmap := map[string]bool{
		"cat":true,
	}

	if testmap["dog"]{
		log.Printf("dog is small I cant woo woo ...")
	}
	if testmap["cat"]{
		log.Printf("cat mi mi ...")
	}



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

	newCat.Say(newCat)
	newDog.Say(newDog)

	newCat.Say(newDog)
	newDog.Say(newCat)

}

// for test interface
type Human struct {
	Pet Animal
}

func (h *Human)Mypetname()string{
	return h.Pet.Myname()
}





type Dog struct {
	Name string
	Friend []Dog
	//只需要定义一个方法  你需要实现的方法只要满足我的这个定义就行
	AddFriendFunc []func(name string)Dog
}

func NewDogwithfriend(){
	mydog := Dog{
		Name:"xiaobai",
	}
	//实现在外部定义方法  然后内部调用   这里主要是用实验性的用方法作为参数  这里只需要提供参数
	mydog.AddFriendFuncs(Newdogfriend)
	for _, friendfunc := range mydog.AddFriendFunc{
		mydog.Friend = append(mydog.Friend,friendfunc("xiaohei"))
	}
	log.Printf("I have now friend it's name is %v", mydog.Friend[0].Name)

}

// 这里只需要提供方法  name这个参数  我们默认是调用方会提供的
func Newdogfriend(name string)Dog{
	return Dog{
		Name:name,
	}
}

func (dog *Dog)AddFriendFuncs(GenrateFriend func(name string)Dog)*Dog{
	dog.AddFriendFunc = append(dog.AddFriendFunc,GenrateFriend)
	return dog
}

func(dog Dog)Say(object Animal){
	// I want to konw if the object is dog
	shouldadog := object.(Dog)

	log.Printf("wangwang... your name is %v", shouldadog.Myname())
}
func (dog Dog)Myname()string  {
	return dog.Name
}




type Cat struct {
	Name string
}
func(cat Cat)Say(object Animal){
	shouldacat := object.(Cat)
	log.Printf("miaomiao...  your name is %v",shouldacat.Myname())
}
func (cat Cat)Myname()string  {
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






