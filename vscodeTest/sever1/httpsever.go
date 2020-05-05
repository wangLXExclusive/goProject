//*********************************************
//					爬取页面
//*********************************************
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main(){
	//发送http请求
	resp,err:=http.Get("http://www.ustc.edu.cn/")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	//读取响应数据
	body,_:=ioutil.ReadAll(resp.Body)
	//解析 a标签里面的href数据
	reg,_:=regexp.Compile(`<a.+href="(http.+?)">`)
	res:=reg.FindAllStringSubmatch(string(body),-1)
	for idx:=range res{
		fmt.Println(res[idx][1])
	}
}
