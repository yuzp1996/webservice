package GETWAY

import (
	"bytes"
	"fmt"
	"net/http"
)

func Getmethod(){
	res, err := http.Get("http://127.0.0.1:3000/time")
	if err != nil{
		fmt.Printf("err happend %v", err)
	}
	defer res.Body.Close()

	headers := res.Header
	for k,v := range headers{
		fmt.Printf("k is %v and v is %v \n", k,v)
	}
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	buf.ReadFrom(res.Body)
	fmt.Printf("res body is %#v", string(buf.Bytes()))

}
