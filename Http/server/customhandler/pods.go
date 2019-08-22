package customhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	core_v1 "k8s.io/api/core/v1"
)

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
	_, _ = fmt.Fprint(w, string(bytes))
}