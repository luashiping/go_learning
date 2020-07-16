package loop_test

import "testing"

// Go语言仅支持循环关键字for
func TestWhileLoop(t *testing.T) {
	n := 0
	/* */
	for n < 5 {
		t.Log(n)
		n++
	}
}

// 无限循环
func TestWhile(t *testing.T) {
	n := 0
	for {
		t.Log(n)
		n++
	}
}
