package main

import (
	GETWAY "webservice/Http/client/GET"
	"webservice/Http/client/POST"
)

func main(){
	//even import GET  but I need GETWAY the name of package
	GETWAY.Getmethod()
	POST.Login()

}


