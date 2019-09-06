package POST

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
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

type Postdata struct{
	Jql string `json:"jql"`
	StartAt int `json:"startAt"`
	MaxResults int `json:"maxResults"`
	Fields []string `json:"fields"`
	Expand []string `json:"expand"`
}

type ListIssuesOptions struct {
	Page     int
	PageSize int
	Project  string
	Type     string
	Priority string
	Status   string
	Summary  string
	IssueKey string
	OrderBy  string
	Sort     string
}


func GetIssueList(){

	listissueoption := ListIssuesOptions{
		Page:1,
		PageSize:4,
		Project:"TEST",
		Type:"task",
		Priority:"Medium",
		Status:"3",
		Summary:"错吧还",
		IssueKey:"TEST-1",
		OrderBy:"updated",
		Sort    :"DESC",
	}


	jql,startat,maxresults:=GetInfofromListOption(listissueoption)



	postdata := Postdata{


		//Jql:"(project = TEST or project = TEST) and priority=Medium and type=task and status=3 and summary ~ '错吧还' and issuekey=TEST-1 ORDER BY updated DESC",
		Jql:jql,
		StartAt:startat,
		MaxResults:maxresults,
		Fields:[]string{
		"created",
		"priority",
		"status",
		"project",
		"updated",
		"creator",
		"assignee",
		"summary",
	},
		Expand:[]string{},
	}


	bodyJson, _ := json.Marshal(postdata)

	klog.Errorf("bodyJson is %v", string(bodyJson))

	httpclient := http.Client{}
	addurl := fmt.Sprintf("http://127.0.0.1:30000/rest/api/2/search")
	req, _ := http.NewRequest(http.MethodPost, addurl, bytes.NewBuffer(bodyJson))
	req.SetBasicAuth("admin", "123456")
	req.Header.Set("Content-Type", "application/json")
	resp, _ := httpclient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	klog.Errorf("resp code is %v and response is %v", resp.StatusCode, string(body))

}


func GetInfofromListOption(option ListIssuesOptions)(jql string,startat,maxresult int){
	//jql = fmt.Sprintf("(project = TEST or project = TEST) and assignee=admin and creator=admin and priority=Medium and type=task and status=3 and summary ~ '错吧还' and issuekey=TEST-1 ORDER BY updated DESC",)
	jql = fmt.Sprintf("ORDER BY %s %s",option.OrderBy,option.Sort)
	Projectlist := []string{"TEST","TEST"}
	var ProjectKeyInfo string
	if option.Project == ""{
		ProjectKeyInfo = fmt.Sprintf("(project = %s",Projectlist[0])
		for _, projectkey := range Projectlist[1:len(Projectlist)]{
			projectforjql := fmt.Sprintf("or Project = %s",projectkey)
			ProjectKeyInfo += projectforjql
		}
		ProjectKeyInfo = fmt.Sprintf("%s)",ProjectKeyInfo)

	}else{
		ProjectKeyInfo = fmt.Sprintf("project = %v", option.Project)
	}

	if option.Type != ""{
		typeforjql := fmt.Sprintf("type = %s", option.Type)
		ProjectKeyInfo = fmt.Sprintf("%s and %s", typeforjql,ProjectKeyInfo)
	}
	if option.Status != ""{
		typeforstatus := fmt.Sprintf("status = %s", option.Status)
		ProjectKeyInfo = fmt.Sprintf("%s and %s", typeforstatus,ProjectKeyInfo)
	}
	if option.IssueKey != ""{
		typeforissuekey := fmt.Sprintf("issuekey = %s", option.IssueKey)
		ProjectKeyInfo = fmt.Sprintf("%s and %s", typeforissuekey,ProjectKeyInfo)
	}
	if option.Summary != ""{
		typeforsummary := fmt.Sprintf("summary ~ '%s'", option.Summary)
		ProjectKeyInfo = fmt.Sprintf("%s and %s", typeforsummary,ProjectKeyInfo)
	}


	return ProjectKeyInfo, (option.Page-1)*option.PageSize,option.PageSize


}