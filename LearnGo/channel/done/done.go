package main

import (
	"fmt"
	"sync"
)

func main() {
	channelDomo()
}

//type  worker struct {//分别并发大小写的worker
//	in chan int
//	done chan bool
//}
type worker struct {//waitGroup需要的worker
	in chan int
	wg *sync.WaitGroup
}
func doWorker(id int,in chan int,wg *sync.WaitGroup){
	for{
		n,ok:=<-in
		if ok {
			fmt.Printf("id:%d,n:%c\n", id, n)
			//go func() {
			//	done<-true
			//}()
			wg.Done()
		}
	}
}

func creatWorker(id int,Wg *sync.WaitGroup)worker{
	work:=worker{
		in:   make(chan int),
		wg: Wg,
	}
	go doWorker(id,work.in,work.wg)
	return work
}
func channelDomo(){
	var wg sync.WaitGroup
	var  workers [10]worker
	for i:=0;i<10;i++{
		workers[i]=creatWorker(i,&wg)
	}
	//**********分别并发的打大小写（先小写后大写）*****************
	//for i, worker:=range  workers{
	//	worker.in<-'a'+i
	//	//<-worker.done
	//}
	////for i,worker:=range workers{//放在这会卡住
	////	worker.in<-'A'+i
	////}
	//for _,worker:=range workers{
	//	<-worker.done
	//}
	//for i,worker:=range workers{
	//	worker.in<-'A'+i
	//}
	//for _,worker:=range workers{
	//	<-worker.done
	//}
	//**********************waitGroup的应用(大小写并发打印)****************
	wg.Add(20)
	for i, worker:=range  workers{
		worker.in<-'a'+i
		//<-worker.done
	}

	for i,worker:=range workers{
		worker.in<-'A'+i
	}

	wg.Wait()
}
