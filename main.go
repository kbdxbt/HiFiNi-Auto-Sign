package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}
	success := SignIn(client)
	if success {
		fmt.Println("签到成功")
	} else {
		fmt.Println("签到失败")
		os.Exit(3)
	}

	success1 := SignIn1(client)
	if success1 {
		fmt.Println("获取成功")
	} else {
		fmt.Println("获取失败")
		os.Exit(3)
	}
}

// SignIn 签到
func SignIn(client *http.Client) bool {
	//生成要访问的url
	url := "https://www.hifini.com/sg_sign.htm"
	cookie := os.Getenv("COOKIE")
	if cookie == "" {
		fmt.Println("COOKIE不存在，请检查是否添加")
		return false
	}
	//提交请求
	reqest, err := http.NewRequest("POST", url, nil)
	reqest.Header.Add("Cookie", cookie)
	reqest.Header.Add("x-requested-with", "XMLHttpRequest")
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(buf))
	return strings.Contains(string(buf), "成功") || strings.Contains(string(buf), "今天已经签过")
}

func SignIn1(client *http.Client) bool {
	//生成要访问的url
	url := "https://w1.v2ai.top/user/checkin"
	token := os.Getenv("TOKEN")
	if token == "" {
		fmt.Println("TOKEN不存在，请检查是否添加")
		return false
	}
	//提交请求
	reqest, err := http.NewRequest("POST", url, nil)
	reqest.Header.Add("Cookie", token)
	reqest.Header.Add("x-requested-with", "XMLHttpRequest")
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(buf))
	return strings.Contains(string(buf), "获得")
}