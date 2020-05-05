package main

func nonRepeatStr(arr string)int{
	start,maxLen:=0,0
	lastOccur:=make(map[rune]int)
	for i,ch:=range []rune(arr){
		lastI,ok:=lastOccur[ch]
		if ok&&start<=lastI{
			start=lastI+1
		}
		lastOccur[ch]=i
		if maxLen<i-start+1{
			maxLen=i-start+1
		}
	}
	return maxLen
}


//func main(){
//	fmt.Println("success")
//	array:=" "
//	fmt.Println(nonRepeatStr(array))
//}
