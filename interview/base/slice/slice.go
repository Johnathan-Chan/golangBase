package slice

import "fmt"

// 获取切片时同时设置切边容量[start:end:cap]
func SetSliceCap(){
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t)
}

// nil切片和空切片
func NilSlice()  {
	//var s1 []int
	var s2 = []int{}
	if s2 == nil {
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}
}
