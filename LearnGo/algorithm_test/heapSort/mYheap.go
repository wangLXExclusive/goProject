package myHeap

import "fmt"

type HEAP struct {
	MyHeap []int
}

func (h *HEAP)Push(x int){
	if len(h.MyHeap)==0{
		h.MyHeap=append(h.MyHeap,x)
		return
	}
	current:=len(h.MyHeap)
	h.MyHeap=append(h.MyHeap,x)
	for current>0&&h.MyHeap[(current-1)/2]>x{
		h.MyHeap[current]=h.MyHeap[(current-1)/2]
		current=(current-1)/2
	}
	h.MyHeap[current]=x
}
func (h *HEAP)Pop(){//弹出顶部元素
	if len(h.MyHeap)==0{
		fmt.Println("heap is null")
		return
	}
	tem:=h.MyHeap[len(h.MyHeap)-1]//将最后一个值存起来
	//删除最后一个元素
	if len(h.MyHeap)==1{
		h.MyHeap=h.MyHeap[:len(h.MyHeap)-1]
		//h.MyHeap=make([]int,0,1)
		return
	}else {
		h.MyHeap = h.MyHeap[:len(h.MyHeap)-1]
	}
	cur,child:=0,1
	for child<len(h.MyHeap){
		if child+1<len(h.MyHeap) {//判断有没有右孩子
			if h.MyHeap[child] > h.MyHeap[child+1] {//有的话比较大小
				child = child + 1
			}
		}
		//判断tem和他孩子的大小，小的话
		if tem<h.MyHeap[child]{
			break
		}
		//大的话
		h.MyHeap[cur]=h.MyHeap[child]
		//改变cur,child
		cur,child=child,child*2+1
	}
	h.MyHeap[cur]=tem
}
func (h *HEAP)IsEmpty()bool{
	return len(h.MyHeap)==0
}
func (h *HEAP)InitHeap(arr []int){
	for _,v:=range arr{
		h.Push(v)
	}
}
func (h *HEAP)Head()int{
	if len(h.MyHeap)==0{
		fmt.Println("Heap is null!!!")
	}
	return h.MyHeap[0]
}