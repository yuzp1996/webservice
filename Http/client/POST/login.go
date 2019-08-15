package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"k8s.io/klog"
	"webservice/Http/server/customhandler"
)

func Login(){

	client := http.Client{}
	req, err := http.NewRequest("POST","http://127.0.0.1:3000/login?name=yuzhipeng&password=123456", nil)
	if err != nil{
		fmt.Printf("error is %v", err)
		klog.Error(err)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Errorf("body is %v", string(body))

	result := &customhandler.BaseJsonBean{}
	err = json.Unmarshal(body,result)
	if err!=nil{
		klog.Errorf("err happend when unmarshal %v", err)
		return
	}

	if result.Code == 100{
		klog.Errorf("result code is %v",result.Code)
	}else{
		klog.Errorf("resutl else is %v message is %v ", result.Code, result.Message)
	}
	return
}