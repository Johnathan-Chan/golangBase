package _map

import "fmt"


// 错误的遍历map并赋值地址的方式
func RangeMapError(){

	slice := []int{1, 2, 3, 4}
	m := make(map[int]*int)

	for k, v := range slice{
		m[k] = &v
	}


	for k, v := range m {
		fmt.Println(k, " -> ", v, " -> ", *v)
	}
}

// 正确的遍历map并赋值地址的方式
func RangeMapSure(){

	slice := []int{1, 2, 3, 4}
	m := make(map[int]*int)

	for k, v := range slice{
		_v := v
		m[k] = &_v
	}


	for k, v := range m {
		fmt.Println(k, " -> ", v, " -> ", *v)
	}
}