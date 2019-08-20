package businesslogic

import "fmt"

type Hellointer interface {
	Hello()
}

func NewCommendOptions()Hellointer{
	return CommendOption{
		Options:[]Hellointer{
			optoner1{"logoption"},
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

type optoner1 struct {
	name string
}
func (O1 optoner1)Hello(){
	fmt.Println(O1.name)
}

type optoner2 struct {
	name string
}

func (O2 optoner2)Hello(){
	fmt.Println(O2.name)
}