package main

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Class struct {
	Id   int64
	Name string
	Desc string
}


func main()  {
	db,err:=sqlx.Open(`mysql`, `root:123456@tcp(127.0.0.1:3306)/news?charset=utf8&parseTime=true`)
	log.Println(db,err)

	//查询
	//get 查询一个	(count())/一条 struct
	//select 一个集合   []
	//非查询
	//db.Exec()   //执行 insert update delete

	mod:=Class{}
	//log.Println("----",mod)
	db.Get(&mod,`select * from class`)
	log.Println(mod)
}