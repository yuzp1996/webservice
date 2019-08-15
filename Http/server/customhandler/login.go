package customhandler

import (
	"encoding/json"
	"fmt"
	"k8s.io/klog"
	"net/http"
)

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
	klog.Errorf("usename is %v",username)
	klog.Errorf("pas is %v", password)
	if !(nameexsit && passwordexsit){
		klog.Errorf("you cant visit this site")
		return
	}
	result := NewBaseJsonBean()
	name := username[0]
	pass := password[0]

	s := fmt.Sprintf("UserName is %s, Password is %s", name, pass)
	klog.Error(s)

	if name == "yuzhipeng" && pass == "123456"{
		result.Message = "登录成功"
		result.Code = 100
		klog.Error("success")
	}else{
		result.Message= "登录失败"
		result.Code = 401
		klog.Error("failed")
	}
	fmt.Printf("result is %#v", result)
	bytes,_ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))

}

