package customhandler

import (
	"net/http"
	"time"
)

type TimeHandler struct {
	Format string
}

func(timehandler *TimeHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	tm := time.Now().Format(timehandler.Format)

	w.Write([]byte("The time is : " + tm + "and the url is " + r.URL.Host))
}
