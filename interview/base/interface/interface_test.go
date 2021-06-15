package _interface

import (
	"fmt"
	"testing"
)

func TestReturnValue(t *testing.T) {
	data := ReturnValue("")
	if data == nil {
		fmt.Println("nil")
	}else{
		fmt.Println("not nil")
	}

	data1 := ReturnValue("res")
	if data1 == nil {
		fmt.Println("nil")
	}else{
		fmt.Println("not nil")
	}
}

func TestGetType(t *testing.T) {
	GetType(123)
}
