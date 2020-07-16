package type_test

import (
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// b = a        // Go语言不允许隐式类型转换
	b = int64(a) // 显示类型转换
	var c MyInt
	// c = b // 别名和原有类型也不能进行隐式类型转换
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a // 通过取指符获得变量的指针
	// aPtr = aPtr + 1 // 不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) // 打印变量类型
}

func TestString(t *testing.T) {
	var s string // string是值类型，取默认的初始化值为空字符串，而不是nil
	t.Log("*" + s + "*")
	t.Log(len(s))
}

func TestTypePredefine(t *testing.T) {
	// 类型的预定义值
	t.Log(math.MaxInt32)
	t.Log(math.MaxInt64)
	t.Log(math.MaxFloat64)
	t.Log(math.MaxUint32)
	// t.Log(math.MaxUint64)

}
