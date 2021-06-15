package _defer

import "fmt"

// defer以栈的方式访问
// defer func() {}() 中使用recover可以恢复panic
func DeferCall() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()
	panic("append error")
}

// 一下两个方法是defer对显示返回和隐式返回的影响
func DeferBefore() int {
	a, b := 1, 2

	sum := a + b

	defer func() {
		sum++
	}()

	return sum
}

func DeferAfter() (sum int) {
	a, b := 1, 2

	sum = a + b

	defer func() {
		sum++
	}()

	return
}

// defer语句会保存变量的copy值，后续变量改变时不会影响这个值
// 但是如果传入的是一个地址就会影响
func DeferValue() {
	var hello func(int)

	hello = func(i int) {
		fmt.Println(i)
	}

	i := 5
	defer hello(i)

	i = i + 10
}

// defer的执行顺序
func DeferPlace() {

	var calc func(string, int, int) int

	calc = func(index string, a, b int) int {
		ret := a + b
		fmt.Println(index, a, b, ret)
		return ret
	}

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
