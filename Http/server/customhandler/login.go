package customhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"net/http"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	core_v1 "k8s.io/api/core/v1"
)

type BaseJsonBean struct {
	Code int `json:"code"`
	Message string 	`json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean{
	return &BaseJsonBean{}
}


func InitClient()(clientset *kubernetes.Clientset, err error){
	var (
		restConf *rest.Config
	)
	if restConf, err = GetRestConfig(); err != nil{
		return
	}
	if clientset, err = kubernetes.NewForConfig(restConf);err != nil{
		goto END
	}
END:
	return
}


func GetRestConfig()(restConf *rest.Config, err error){
	var (
		kubeconfig []byte
	)
	if kubeconfig, err = ioutil.ReadFile("./config");err!=nil{
		goto END
	}
	if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeconfig);err!=nil{
		goto  END
	}
END:
	return
}




func Getpods(w http.ResponseWriter, req *http.Request){

	var(
		clientset *kubernetes.Clientset
		podsList *core_v1.PodList
		err error
	)
	result := NewBaseJsonBean()


	if clientset, err = InitClient();err !=nil{
		return
	}
	if podsList, err = clientset.CoreV1().Pods("default").List(meta_v1.ListOptions{});err != nil{
		return
	}
	for _, pod := range podsList.Items{
		result.Message += fmt.Sprintf("%v    ",pod.Name)
	}
	bytes,_ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))


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

