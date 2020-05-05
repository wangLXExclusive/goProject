package search

import (
	"fmt"
	"log"
	"sync"
)

type Result struct {
	Filed string
	Content string
}

func Run(searchTerm string){
	feeds,err:=RetrieverFeeds()
	if err!=nil{
		log.Fatal(err)
	}
	results:=make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))
	for _,feed:=range feeds{
		go func(feed *Feed) {
			searchResult,err:=Search(feed,searchTerm)
			if err!=nil{
				log.Println(err)
				return
			}
			for _,result:=range searchResult{
				results<-result
			}
			waitGroup.Done()
		}(feed)
	}
	go func() {
		//等候所有任务完成
		waitGroup.Wait()

		//用关闭通道的方式通知Display函数
		//可以退出程序了
		close(results)
	}()
	Display(results)
}
func Display(results chan *Result)  {
	//通道会一直堵塞，直到有结果写入
	//一旦通道被关闭，for循环就会终止
	for result:=range results{
		fmt.Printf("%s:\n%s\n\n",result.Filed,result.Content)
	}
}