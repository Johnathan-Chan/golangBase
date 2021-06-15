package string

import (
	"fmt"
	"testing"
)

func TestLinkString(t *testing.T) {
	data := []string{"test", "ing", "test", "ing", "test", "ing","test", "ing", "test", "ing"}

	fmt.Println(LinkStringPlus(data))
	fmt.Println(LinkStringSprint(data))
	fmt.Println(LinkStringJoin(data))
	fmt.Println(LinkStringBuffer(data))
	fmt.Println(LinkStringBuilder(data))
}

var data []string

func BenchmarkLinkStringPlus(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinkStringPlus(data)
	}
}

func BenchmarkLinkStringSprint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinkStringSprint(data)
	}
}

func BenchmarkLinkStringJoin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinkStringJoin(data)
	}
}

func BenchmarkLinkStringBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinkStringBuffer(data)
	}
}

func BenchmarkLinkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinkStringBuilder(data)
	}
}

func TestMain (m *testing.M){
	data = 	[]string{"test", "ing", "test", "ing", "test", "ing","test", "ing", "test", "ing"}
	m.Run()
}

