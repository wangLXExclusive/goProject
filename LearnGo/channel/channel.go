package main

import (
	"fmt"
	"time"
)

func main() {
	//channelDomo()
	bufferChannel()
	time.Sleep(time.Millisecond)
}
func worker(id int,channel chan int){
	for{
		//fmt.Println("id:",id,"n:",<-channel)
		n,ok:=<-channel
		//if !ok{
		//	break
		//}
		if ok {
			fmt.Printf("id:%d,n:%c\n", id, n)
		}
	}
}

func creatWorker(id int)chan<- int{
	channel:=make(chan int)
	go func() {
		for{
			//fmt.Println("id:",id,"n:",<-channel)
			fmt.Printf("id:%d,n:%c\n",id,<-channel)
		}
	}()
	return channel
}
func channelDomo(){
	//var channels [10]chan int
	//for i:=0;i<10;i++ {
	//	channels[i]=make(chan int)
	//	go worker(i, channels[i])
	//}
	var channels  [10]chan<- int
	for i:=0;i<10;i++{
		channels[i]=creatWorker(i)
	}
	for i:=0;i<10;i++{
		channels[i]<-'a'+i
	}
}
func bufferChannel(){
	c:=make(chan int)
	go worker(0,c)
	c<-'a'
	c<-'b'
	c<-'c'
	c<-'d'
	close(c)
}