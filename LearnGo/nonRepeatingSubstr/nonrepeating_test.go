package main

import "testing"

func TestSubstr(t *testing.T)  {
	tests:=[]struct {
		a string
		ans int}{
		{"abcbacbb",3},
		{"pwwkew",3},
		{"",0},
		{"b",1},
		{"bbbbbb",1},
		{"abcabcabcd",4},
		{"这里是慕课网",6},
		{"黑化肥黑狐黑",4},
	}

	for _,tt:=range  tests{
		actual:=nonRepeatStr(tt.a)
		if actual!=tt.ans{
			t.Errorf("got %d for input %s;"+"expect %d",actual,tt.a,tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B){
	s,ans:="黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花",8
	for i:=0;i<b.N;i++ {
		actual := nonRepeatStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s;"+"expect %d", actual, s, ans)
		}
	}
}