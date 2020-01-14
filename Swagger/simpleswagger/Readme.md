# Simple Server的使用方法

通过swagger生成代码  cd ~/go/src/webservice/Swagger/simpleswagger & swagger generate server -A TodoList -f ./swagger.yml


cd ~/go/src/webservice/Swagger/simpleswagger/cmd/todo-list-server  & go install & todo-list-server

会有提示提醒在哪个端口上运行  直接访问对应端口即可

例如如果想下载文件 直接在浏览器访问 http://127.0.0.1:53442/1 即可
