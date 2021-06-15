package src

import (
	"fmt"
	"testing"
)

func hello(num ...int) {
	num[0] = 18
}

func TestHello(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

func hello2(num []int){
	num[0] = 18
}

func TestHello2(t *testing.T) {
	i := []int{5, 6, 7}
	hello2(i)
	fmt.Println(i[0])
}
