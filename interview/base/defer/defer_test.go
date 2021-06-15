package _defer

import (
	"fmt"
	"testing"
)

func TestDeferCall(t *testing.T) {
	DeferCall()
}

func TestDeferAppend(t *testing.T) {
	resBefore := DeferBefore()
	fmt.Println(resBefore)

	resAfter := DeferAfter()
	fmt.Println(resAfter)
}

func TestDeferValue(t *testing.T) {
	DeferValue()
}

func TestDeferPlace(t *testing.T) {
	DeferPlace()
}