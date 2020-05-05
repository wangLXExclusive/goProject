package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//斐波那契数列生成器
func fibonacci() intGen{
	a,b:=0,1
	return func()int{
		a,b=b,a+b
		return a
	}
}

type intGen func()int
func (g intGen)Read(p []byte) (n int, err error){
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
func printFileContent(reader io.Reader){
	scanner:=bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}
func writeFile(fileName string){
	file,err:=os.Create(fileName)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	writer:=bufio.NewWriter(file)
	defer writer.Flush()
	f:=fibonacci()
	for i:=0;i<20;i++{
		fmt.Fprintln(writer,f())
	}
}
func main()  {
	fmt.Println(1&^1)
	writeFile("abc.txt")
}
