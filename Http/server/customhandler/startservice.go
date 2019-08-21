package customhandler

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"net/http"
	"time"
	"webservice/Http/server/businesslogic"
)


func StartServer(){

	// make some init function

	name := viper.Get("name")
	port := viper.GetString("port")

	color.Cyan("name is %v",name)
	color.Blue("listening port %v", port)

	businesslogic.NewCommendOptions().Hello()



	mux := http.NewServeMux()
	// first way
	redirehandler := http.RedirectHandler("http://www.baidu.com",307)
	mux.Handle("/foo", redirehandler)

	//second way  struct with function ServeHTTP
	timehandler := &TimeHandler{Format: time.RFC822}
	mux.Handle("/time", timehandler)

	timehandler3339 := &TimeHandler{Format: time.RFC3339}
	mux.Handle("/time3339", timehandler3339)


	//third way handlerfunc only use function
	timehandlerfunc := http.HandlerFunc(timeHandler)
	mux.Handle("/timefunc", timehandlerfunc)

	//fourth way the better way of third  the usual way
	mux.HandleFunc("/shorttime",timeHandler)


	//var func you can input parama just like third way
	mux.Handle("/paratime", timeHandlerwithPara(time.RFC3339))


	// post way
	mux.HandleFunc("/login",Login)


	http.ListenAndServe(fmt.Sprintf(":%s",port), mux)

}


func timeHandler(w http.ResponseWriter, r *http.Request){
	tm := time.Now().Format(time.RFC3339)
	w.Write([]byte("The tiem is: "+tm))
}

func timeHandlerwithPara(format string)http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("time is: "+tm))
	}
	return http.HandlerFunc(fn)
}


