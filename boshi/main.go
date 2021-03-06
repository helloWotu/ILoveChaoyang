package main

import (
	"boshi/model"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const usrID string = "5cf220836282c907f331de21"

func main() {

	//findAnswer();
	//httprou();

postNet()

}

func postNet()  {
	url := "https://bs.yunshicloud.com/api/evaluation/answer/exams"
	//填写表单
	args := &fasthttp.Args{}
	args.Add("_id",usrID)
	args.Add("exam_name","")

	statusCode, body, err := fasthttp.Post(nil, url, args)
	if err != nil {
		fmt.Printf("请求失败 error:%s\n", err.Error())
		return
	}

	if statusCode != fasthttp.StatusOK{
		fmt.Println("请求没有成功:", statusCode)
		return
	}

	fmt.Println(string(body))

}

func httprou()  {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprintf(writer, "Blog:%s \nWechat:%s","www.flysnow.org","flysnow_org")
	})

	log.Fatal(http.ListenAndServe(":8080",router))
}

func findAnswer()  {
	file, _ := os.Open("boshi.json")
	bytes, _ := ioutil.ReadAll(file)
	file.Close()

	jsonData := model.AutoGenerated{}
	err := jsoniter.Unmarshal(bytes, &jsonData)
	if err != nil {
		fmt.Printf("error:%s\n", err.Error())
	}
	//fmt.Println(jsonData.Data.User.Nickname)

	for index,answer_info := range jsonData.Data.Exam.AnswerInfo {
		fmt.Printf("题目%d：\n",index+1)
		for idx,topic_option := range answer_info.TopicOption {
			if topic_option.CorrectOption == 1 {
				//说明这是一个正确答案
				fmt.Print(idx+1)
			}else {
				continue
			}
		}
		fmt.Print("\n")
	}
}

