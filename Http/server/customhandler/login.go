package customhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"os"
)

// New logger you define you change the logger preix and flag
//
// it is same as you use log  and set the log.SetPrefix and log.SetFlags
var Error *log.Logger


func init(){
	log.SetPrefix("[Login]")
	log.SetFlags(log.Ldate|log.Lshortfile|log.Llongfile)

	Error = log.New(os.Stderr, "[Login ]", log.Ldate|log.Llongfile|log.Ltime)
}

type BaseJsonBean struct {
	Code int `json:"code"`
	Message string 	`json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean{
	return &BaseJsonBean{}
}



func Login(w http.ResponseWriter, req *http.Request){
	req.ParseForm()
	username, nameexsit := req.Form["name"]
	password, passwordexsit := req.Form["password"]
	log.Printf("usename is %v",username)
	log.Printf("pas is %v", password)
	if !(nameexsit && passwordexsit){
		log.Printf("you cant visit this site")
		return
	}
	result := NewBaseJsonBean()
	name := username[0]
	pass := password[0]

	s := fmt.Sprintf("UserName is %s, Password is %s", name, pass)
	log.Printf(s)

	if name == "yuzhipeng" && pass == "123456"{
		result.Message = "登录成功"
		result.Code = 100
		log.Printf("success")
	}else{
		result.Message= "登录失败"
		result.Code = 401
		Error.Println("Failed")
	}


	bytes,_ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))

}





