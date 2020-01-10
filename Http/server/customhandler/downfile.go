package customhandler


import (
"fmt"
"io"
"net/http"
	"net/url"
)

func fileHandler(rw http.ResponseWriter, r *http.Request) {
	zipName := "zpyu.zip"
	// 设置rw的header信息中的ctontent-type，对于zip可选以下两种
	// rw.Header().Set("Content-Type", "application/octet-stream")
	//rw.Header().Set("Content-Type", "text/plain")
	rw.Header().Set("Content-Type", "application/zip")
	// 设置rw的header信息中的Content-Disposition为attachment类型
	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))
	downloadjenkisnfile(rw)
}


func downloadjenkisnfile(rw http.ResponseWriter){

	// Get the data
	httpclient:=http.Client{}
	url := url.URL{
		Host:"127.0.0.1:8080",
		//Path:"/job/zpyu/job/zpyutest/3/artifact/zpyu2",
		Path:"/job/zpyu/job/zpyutest/3/artifact/*zip*/archive.zip",

		Scheme: "http",
		//http://127.0.0.1:8080/job/zpyu/job/zpyutest/2/artifact/*zip*/archive.zip
	}
	request := http.Request{
		Method: http.MethodGet,
		URL: &url,
		Header:  http.Header{"Authorization":[]string{"Basic YWRtaW46MTFlOTg4NGI5MjI5MTRhNjc0Njk1MjY2N2Y3NjI2YWIyZg=="}},
	}
	resp, err :=httpclient.Do(&request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(rw,resp.Body)

	if err != nil {
		panic(err)
	}
	return

}
