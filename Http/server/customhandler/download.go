package customhandler
//https://www.jianshu.com/p/0c3e005dbf75

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func fileHandler(rw http.ResponseWriter, r *http.Request) {
	zipName := "zpyu.zip"
	// 设置rw的header信息中的ctontent-type，对于zip可选以下两种
	rw.Header().Set("Content-Type", "application/octet-stream")
	//rw.Header().Set("Content-Type", "text/plain")
	//rw.Header().Set("Content-Type", "application/zip")
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
		Path:"/job/zpyu/job/zpyutest/22/artifact/*zip*/archive.zip",

		Scheme: "http",
		//http://127.0.0.1:8080/job/zpyu/job/zpyutest/2/artifact/*zip*/archive.zip
	}
	request := http.Request{
		Method: http.MethodGet,
		URL: &url,
		Header:  http.Header{"Authorization":[]string{"Basic YWRtaW46MTFlOTg4NGI5MjI5MTRhNjc0Njk1MjY2N2Y3NjI2YWIyZg=="}},
	}


	resp, err :=httpclient.Do(&request)

	resbody, _:=ioutil.ReadAll(resp.Body)
	log.Printf("header is %v", resp.Header)
	log.Printf("resp is %v", string(resbody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 这里将data写入了rw中，应该也就是一个reader之类的东西吧

	rw.Header()["Content-Length"] = []string{resp.Header.Get("Content-Length")}

	_, err = io.Copy(rw,resp.Body)

	if err != nil {
		panic(err)
	}
	return


}

//func downloadjenkisnfile(rw http.ResponseWriter){
//
//	// Get the data
//	httpclient:=http.Client{}
//	url := url.URL{
//		Host:"192.168.16.71",
//		//Path:"/job/zpyu/job/zpyutest/3/artifact/zpyu2",
//		Path:"/devops/api/v1/pipeline/zpyu/dsfds-xb62h/download",
//
//		Scheme: "http",
//		//http://127.0.0.1:8080/job/zpyu/job/zpyutest/2/artifact/*zip*/archive.zip
//	}
//	request := http.Request{
//		Method: http.MethodGet,
//		URL: &url,
//		Header:  http.Header{"Authorization":[]string{"Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImUyYTVlYTU2ZmYyNDQ4NjM3ZDk4Yjg0Yzc5MmIzZjUxYTlkZjM1YmIifQ.eyJpc3MiOiJodHRwczovLzE5Mi4xNjguMTYuNzEvZGV4Iiwic3ViIjoiQ2lRd09HRTROamcwWWkxa1lqZzRMVFJpTnpNdE9UQmhPUzB6WTJReE5qWXhaalUwTmpZU0JXeHZZMkZzIiwiYXVkIjoiYWxhdWRhLWF1dGgiLCJleHAiOjE1Nzk1OTQyMzEsImlhdCI6MTU3OTUwNzgzMSwibm9uY2UiOiJhbGF1ZGEtY29uc29sZSIsImF0X2hhc2giOiJVTFlmMEZNcW1palRmekVCMXNvQWdRIiwiZW1haWwiOiJhZG1pbkBjcGFhcy5pbyIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJuYW1lIjoiYWRtaW4iLCJleHQiOnsiaXNfYWRtaW4iOnRydWUsImNvbm5faWQiOiJsb2NhbCJ9fQ.wQG6c92dCYUp6Di4ZAl--5kwCm9pbIopHXfWz_-K6Vey6ucrlSQuXsFsVqM8PxQN0BgIGCbDz27ZoIfVykZlK67bydcPGj065bm6kE2imR1uLblI_Q-Pqn8HlEmmlhzs-s-hHNSspwa-uX0MIs8rD34s6gzyDpDN3LZ0BwoFEX0Tehk9Hw3h5vvuq3Vwx-itsgB7Gi6NAIXCzP5NVgE7ItKOTAMzEvVkEmM1jhRBEIUZ8MoZ-5KBdBEdpvjVD7zDxoxZ22stFO7zYB3GEVza_B2HaB8dC6SdorqUheTxcXTXN9mVtyQBIxe1Z87q0FE4I9Yhh3aUz87toMiyhmzYvA"}},
//	}
//
//
//	resp, err :=httpclient.Do(&request)
//	//resbody, _:=ioutil.ReadAll(resp.Body)
//	//log.Printf("resp is %v", string(resbody))
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	// 这里将data写入了rw中，应该也就是一个reader之类的东西吧
//	_, err = io.Copy(rw,resp.Body)
//
//	if err != nil {
//		panic(err)
//	}
//	return
//
//
//}