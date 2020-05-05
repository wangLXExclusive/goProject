package main

import (
	_ "goInAction/sample/matchers"
	"goInAction/sample/search"
	"log"
	"os"
)

func init(){
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main(){
	//使用特定的项做搜索
	search.Run("汽车")
}
