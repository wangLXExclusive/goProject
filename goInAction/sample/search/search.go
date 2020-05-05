package search

import (
	"log"
	"sync"
)

var matchers=make(map[string]Matcher)

func Run(searchTerm string){
	 feeds, err := RetrieveFeeds()

	if err!=nil{
		log.Fatal(err)
	}
	//创建一个无缓存的通道，接收匹配后的结果
	restults:=make(chan *Result)
	//构造一个waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	//设置需要等待处理
	//每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	//为每一个数据源启动一个goroutine来查找结果
	for _,feed:=range  feeds{
		//获取一个匹配器用于查找
		matcher,exists:=matchers[feed.Type]
		if !exists{
			matcher=matchers["default"]
		}

		//启动一个goroutine执行查找
		go func(matcher Matcher,feed *Feed) {
			Match(matcher,feed,searchTerm,restults)
			waitGroup.Done()
		}(matcher,feed)
	}

	//启动一个goroutine来监控是否所有的工作都做完了
	go func() {
		//等候所有任务完成
		waitGroup.Wait()

		//用关闭通道的方式通知Display函数
		//可以退出程序了
		close(restults)
	}()

	//启动函数显示返回结果，并且
	//在最后一个结果显示完成后返回
	Display(restults)
}

//Register调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string,matcher Matcher){
	if _,exists:=matchers[feedType];exists{
		log.Fatalln(feedType,"Matcher already registered")
	}


	log.Println("Register",feedType,"matcher")
	matchers[feedType]=matcher
}