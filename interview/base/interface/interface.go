package _interface

import "fmt"

// 返回nil和返回空值的区别
func ReturnValue(s string) []string {

	if s == "" {
		return nil
	}

	return []string{}
}

// 接口类型才可进行断言
func GetType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}
}
