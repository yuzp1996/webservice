package businesslogic

import "fmt"

type Hellointer interface {
	Hello()
}

func NewCommendOptions()Hellointer{
	return CommendOption{
		Options:[]Hellointer{
			port{"logoption"},
			optoner2{"pamaration"},
		}}
}

type CommendOption struct {
	Options []Hellointer
}

func(co CommendOption)Hello(){
	for _,value := range co.Options{
		value.Hello()
	}
}

type port struct {
	name string
}
func (O1 port)Hello(){
	fmt.Println(O1.name)
}

type optoner2 struct {
	name string
}

func (O2 optoner2)Hello(){
	fmt.Println(O2.name)
}