package operator_test

import (
	"strconv"
	"testing"
)

// const

func TestCompareArray(t *testing.T) {
	// 用 == 比较数组：相同维数且含有相同个数元素的数组才可以比较
	// 每个元素都相同的才相等
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 2}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)
	t.Log(a == d)
}

func TestBitwiseAndOperator(t *testing.T) {
	var a int64 = 14
	// var b int64 = 3
	v1 := strconv.FormatInt(a, 2)
	t.Log(v1)

	// 按位与: 两个相应的二进位都为1,该位的结果值才为1,否则为0
	v := a & 1 // 按位与判断奇偶数
	if v == 0 {
		t.Log("a is even number")
	} else {
		t.Log("a is odd number")
	}
	t.Log(v)
}
