package main

import (
	"bufio"
	"fmt"
	"io"
	"main/retriever/mock"
	"strings"
	"time"
)

type Retriever interface {
	Get(source string) string
}

func Download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	case mock.Retriever1:
		fmt.Println("Content:", v.Content)
	}
}
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func fibonacci() intGen { //斐波那契数列生成器
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//主函数
func main() {
	//接口测试部分
	var r Retriever
	r = &mock.Retriever{
		UserAgent: "wang",
		TimeOut:   time.Minute,
	}
	r1 := mock.Retriever1{"ksdvjsdvksd"}
	//fmt.Println(Download(r))
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	inspect(r1)
	//闭包测试部分
	//a := adder()
	f := fibonacci()
	printFileContent(f)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
}
