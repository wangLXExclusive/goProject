package main

import (
	"LearnGo/algorithm_test/heapSort"
	"fmt"
	"math"
	"math/rand"
)

//归并排序
func merge(arr []int, p int, q int, r int) []int { //合并两个排序好部分数组的算法

	n1, n2 := q-p+1, r-q
	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = arr[p+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[q+j+1]
	}
	L[n1], R[n2] = math.MaxInt64, math.MaxInt64
	i, j := 0, 0

	for k := p; k <= r; k++ {
		if L[i] < R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
	}
	return arr
}
func mergeSort(arr []int, p int, r int) []int {
	if p < r {
		q := (p + r) / 2
		mergeSort(arr[:], p, q)
		mergeSort(arr[:], q+1, r)
		merge(arr[:], p, q, r)
	}
	return arr
}

//插入排序
func insertSort(array []int) []int {
	for j := 2; j < len(array); j++ {
		key := array[j]
		i := j - 1
		for i > 0 && array[i] > key {
			array[i+1] = array[i]
			i = i - 1
		}
		array[i+1] = key
	}
	return array
}

//最大子数组算法区域
func findMaxCrossingSubarray(arr []int, low int, mid int, high int) (int, int, int) { //找到跨中间最大子数组
	left_sum, right_sum, sum, left_index, right_index := math.MinInt64, math.MaxInt64, 0, mid, mid+1
	for i := mid; i >= low; i-- {
		sum += arr[i]
		if sum > left_sum {
			left_sum = sum
			left_index = i
		}
	}
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum += arr[j]
		if sum > right_sum {
			right_sum = sum
			right_index = j
		}
	}
	return left_index, right_index, left_sum + right_sum
}
func findMaximunSubarray(array []int, low int, high int) (int, int, int) { //找到最大子数组返回左右坐标以及值
	if low == high {
		return low, low, array[low] //只有一个值的时候
	} else {
		mid := (low + high) / 2
		left_low, left_high, left_sum := findMaximunSubarray(array, low, mid)        //找左边
		right_low, right_high, right_sum := findMaximunSubarray(array, mid+1, high)  //找右边
		mid_low, mid_high, mid_sum := findMaxCrossingSubarray(array, low, mid, high) //找跨中间的
		if left_sum >= right_sum && left_sum >= mid_sum {                            //判断这三个哪个大
			return left_low, left_high, left_sum
		} else if right_sum >= left_sum && right_sum >= mid_sum {
			return right_low, right_high, right_sum
		} else {
			return mid_low, mid_high, mid_sum
		}
	}
}
						
//矩阵乘法算法区域
type Matrix [][]int
func squareMatrixMultiply(a Matrix, b Matrix) Matrix {
	n := len(a)
	var c Matrix //因为没有指定空间，所以要用append分配空间不能直接用c[i][j]赋值
	//var c [2][2]int
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			tmp = append(tmp, 0) //同理先分配空间
			for k := 0; k < n; k++ {
				tmp[j] += a[i][k] * b[k][j]
			}
		}
		c = append(c, tmp) //同理
	}
	return c
}
//堆排序算法区域
func heapSort(arr []int)[]int{
	h:=myHeap.HEAP{}
	h.InitHeap(arr)
	for i:=0;!h.IsEmpty();i++{
		arr[i]=h.Head()
		h.Pop()
	}
	return arr
}

//快速排序算法区域
func partition(arr []int,p int,r int)int{
	x,i:=arr[r],p-1
	for j:=p;j<=r-1;j++{					//在循环体内每一轮迭代开始时，对于下标k，有：
		if arr[j]<=x {						//1.若p<=k<=i,则arr[k]<=x
			i++								//2.若i+1<=k<=j-1,则arr[k]>x
			arr[i], arr[j] = arr[j], arr[i] //3.若k=r,则arr[k]=x
		}
	}
	arr[i+1],arr[r]=arr[r],arr[i+1]
	return i+1
}
func randInt(min int,max int)int{
	//if min >= max || min == 0 || max == 0 {
	//	return max
	//}
	return rand.Intn(max-min) + min
}
func randomizedPartition(arr []int,p int,r int)int{
	i:=randInt(p,r)
	arr[i],arr[r]=arr[r],arr[i]
	return partition(arr,p,r)
}
func quickSort(arr []int,p int, r int){
	if p<r{
		q:=randomizedPartition(arr,p,r)
		quickSort(arr,p,q-1)
		quickSort(arr,q+1,r)
	}
}

//钢条切割问题
func max(a int,b int)int{
	if a<b{
		return  b
	}else{
		return a
	}
}
func cutRod(p []int,n int)int{//递归解法（缺陷是反复计算相同的子问题，时间复杂度为2^n）
	if n==0{
		return 0
	}
	q:=math.MinInt64
	for i:=1;i<=n;i++{
		q=max(q,p[i]+cutRod(p,n-i))
	}
	return q
}
func memoizedCutRodAux(p []int,n int,r []int)int{//带备忘的自顶向下方法
	if r[n]>=0{//因为r已经被初始化为负无穷，然后收入为非负值，所以大于0表示有值
		return r[n]
	}
	//没有记录值得话，进行朴素递归
	if n==0{
		return  0
	}
	q:=math.MinInt64
	for i:=1;i<=n;i++{
		q=max(q,p[i]+memoizedCutRodAux(p,n-i,r))
	}
	//记录n是q的值
	r[n]=q
	return q
}
func memoizedCutRod(p []int,n int)int{//带备忘的自顶向下方法
	r:=make([]int,n+1)
	for i:=0;i<n+1;i++{
		r[i]=math.MinInt64
	}
	return memoizedCutRodAux(p,n,r)
}
func bottomUpCutRod(p []int,n int)int{//自底向上的动态规划
	r:=make([]int,n+1)
	var q int
	r[0]=0
	for j:=1;j<=n;j++{
		for i:=1;i<=j;i++{
			q = max(q, p[i]+r[j-i])
		}
		r[j]=q
	}
	return r[n]
}
//主函数
func main() {
	//***********分治算法问题（最大子数组问题）*****************
	//arr := [...]int{-1, 5, -9, 6, 4, 7, -8, 5, 4, -6, -9, 4, 9, 9, -5}
	////arr := [...]int{-1, -9, -5, -3, -4, -7}
	//left, right, sum := findMaximunSubarray(arr[:], 0, len(arr)-1)
	//fmt.Println(left, right, sum)
	//fmt.Println(INT_MIN)
	//************分治算法问题（矩阵乘法）*********************
	//a := Matrix{
	//	{2, 3, 4},
	//	{4, 5, 6},
	//	{3, 4, 5},
	//}
	//b := Matrix{
	//	{2, 3, 4},
	//	{4, 5, 6},
	//	{3, 4, 5},
	//}
	//c := squareMatrixMultiply(a, b)
	//fmt.Println(c)
	// //*********测试排序算法部分**********
	//array := [...]int{1, 5, 8, 6, 7, 1, 3, 4, 9, 10}
	// //fmt.Println("排序前：", array[:])  //测试插入算法
	// //fmt.Println("排序后：", insertSort(array[:]))
	// fmt.Println("排序前：", array[:])//测试归并
	// fmt.Println("排序后：", mergeSort(array[:], 0, len(array)-1))
	// fmt.Println(INT_MAX)
	//*****************堆测试****************
	//arr1:=[...]int{1, 5, 8, 6, 7, 1, 3, 4, 9, 10,11,0,-8}
	//fmt.Println("排序后：",heapSort(arr1[:]))
	//***********快速排序测试区域***************
	//arr := [...]int{1, 2, 8, 4, 6, 3, 9, 5, 4, 8, 2, 0}
	//quickSort(arr[:], 0, len(arr)-1)
	//fmt.Println(arr)

	//钢条切割问题区域
	arr:=[...]int{0,1,5,8,9,10,17,17,20,24,30}//0没有数值

	fmt.Println(cutRod(arr[:],5))
	fmt.Println(memoizedCutRod(arr[:],5))
	fmt.Println(bottomUpCutRod(arr[:],5))
}