package string

import (
	"bytes"
	"fmt"
	"strings"
)

// 字符串的连接方式 速度 Builder > join > buffer > Sprintf > +
// 使用 “+”
func LinkStringPlus(input []string) string{

	res := ""

	for _, v := range input{
		res += v
	}

	return res
}

// 使用 fmt.Sprint
func LinkStringSprint(input []string) string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", input[0], input[1], input[2], input[3],
		input[4], input[5], input[6], input[7], input[8], input[9])
}

// // 使用join
func LinkStringJoin(input []string) string {
	return strings.Join(input, "")
}

// 使用Buffer
func LinkStringBuffer(input []string) string {
	var bf bytes.Buffer

	n := len(input) - 1
	for _, v := range input{
		n += len(v)
	}

	bf.Grow(n)


	for _, v := range input {
		bf.WriteString(v)
	}

	return bf.String()
}

// 使用Builder
func LinkStringBuilder(input []string) string {
	var bf strings.Builder

	for _, v := range input {
		bf.WriteString(v)
	}

	return bf.String()
}