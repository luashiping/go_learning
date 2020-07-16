package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	var (
		a int = 1
		b     = 1
		c int = 1
	)
	// 赋值可以进行自动类型推断
	// a := 1
	// b := 1
	t.Log("hehe", c)
	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// 交换两个变量
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// 借助一个中间变量
	// tmp := a
	// a = b
	// b = tmp
	// Go语言里，一个赋值语句中可以对多个变量进行同时赋值
	a, b = b, a
	t.Log(a, b)
}
