package businesslogic

import "fmt"

type Hellointerface interface {
	Hello()
}

func NewCommendOptions()Hellointerface{
	return CommendOption{
		Options:[]Hellointerface{
			option{"first option"},
			optione1{"second option"},
		}}
}

type CommendOption struct {
	Options []Hellointerface
}

func(co CommendOption)Hello(){
	for _,value := range co.Options{
		value.Hello()
	}
}

type option struct {
	name string
}
func (O option)Hello(){
	fmt.Println(O.name)
}

type optione1 struct {
	name string
}

func (O optione1)Hello(){
	fmt.Println(O.name)
}