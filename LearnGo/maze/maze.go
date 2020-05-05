package main

import (
	"fmt"
	"os"
)

func main() {
	maze:=readMaze("maze/in.txt")
	//maze:=make([][]int,6)
	//maze=[][]int{
	//	{0,1,0,0,0},
	//	{0,0,0,1,0},
	//	{0,1,0,1,0},
	//	{1,1,1,0,0},
	//	{0,1,0,0,1},
	//	{0,1,0,0,0},
	//}
	for _,row:=range maze{
		for _,val:=range row{
			fmt.Printf("%2d",val)
		}
		fmt.Println()
	}
	steps:=walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})
	for _,row:=range steps{
		for _,val:=range row{
			fmt.Printf("%2d   ",val)
		}
		fmt.Println()
	}
}
type point struct{
	i,j int
}
func (p point)add(r point)point{
	return point{p.i+r.i,p.j+r.j}
}
func (p point)at(grid [][]int)(int,bool){
	if p.i<0||p.i>=len(grid){
		return  0,false
	}
	if p.j<0||p.j>=len(grid[0]){
		return 0,false
	}
	return grid[p.i][p.j],true
}
var dirs=[4]point{
	{-1,0},{0,-1},{1,0},{0,1},
}
func walk(maze [][]int,start,end point)[][]int{
	step:=make([][]int,len(maze))
	for i:=range maze{
		step[i]=make([]int,len(maze[0]))
	}

	Q:=[]point{start}
	for len(Q)>0{
		cur:=Q[0]
		Q=Q[1:]
		if cur==end{//发现终点退出
			break
		}


		for _,dir:=range dirs{
			next:=cur.add(dir)

			//maze at next is 0
			//and step at next is 0
			//and next!=start
			val,ok:=next.at(maze)
			if !ok||val==1{
				continue
			}
			val,ok=next.at(step)
			if !ok||val!=0{
				continue
			}
			if next==start{
				continue
			}
			curStep,_:=cur.at(step)
			step[next.i][next.j]=curStep+1
			Q=append(Q,next)
		}
	}
	return step
}
func readMaze(filename string)[][]int{
	file,err:=os.Open(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	var  row,col int
	fmt.Fscanf(file,"%d %d",&row,&col)

	maze:=make([][]int,row)
	for i:=range maze{
		maze[i]=make([]int,col)
		for j:=range maze[1] {
			//fmt.Fscanf(file, "%d", &maze[i][j])

			fmt.Fscanf(file,"%d",&maze[i][j])

		}
	}
	return maze
}