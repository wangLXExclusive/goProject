package main

import (
	"os"
)

func main() {

	//***********创建文件************
	f, _ := os.Create("1.txt")
	//f.Read()读数据
	//f.Write()写数据
	//f.Close()
	//f.Sync()同步数据
	f.WriteString("hello")

	//***********************读文件********
	// 	file, err := os.Open("test.txt")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	buf := make([]byte, 1000)
	// 	ln, err := file.Read(buf)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(string(buf[:ln]))

	//***************创建文件夹***********
	// os.Mkdir("dir", 0777)        //创建目录
	// os.MkdirAll("a/b/c/d", 0666) //创建多级目录
	// //0777   0开头 ->八进制
	// //7   7   7
	// //rwx rwx rwx     读写执行
	// //权限 拥有者 我们班 其他班
}
