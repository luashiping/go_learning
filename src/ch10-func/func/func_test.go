package func_test

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMutilValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op

}
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func TestFn(t *testing.T) {
	// a, _ := returnMutilValues()
	// t.Log(a)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))

	ret := timeSpent(slowFun)(10)
	t.Log(ret)
}

//可变参数，函数的参数数量可变
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}

	return ret
}
func TestVarParam(t *testing.T) {
	ret := Sum(1, 2, 3, 4)
	t.Log(ret)

	ret1 := Sum(1, 2, 3, 4, 5)
	t.Log(ret1)

}

/*
命名返回参数
不止函数的参数可以有变量名称，函数的返回值也可以，也就是说你可以为每个返回值都起一个名字，这个名字可以像参数一样在函数体内使用
*/
func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil

	return
}

func TestReturnName(t *testing.T) {
	res, _ := sum(1, 3)
	fmt.Println(res)
}

// 闭包
func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func TestClosure(t *testing.T) {
	cl := colsure()

	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
}

// 不同于函数的方法
type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}

func TestMethods(t *testing.T) {
	// age := Age(26)
	// age.String()
	age := Age(25)
	sm := Age.String
	sm(age)
}

func Clear() {
	fmt.Println("clear resources.....")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	panic("err")
	// fmt.Println("End")
}

/*
来自蔡超老师的<<Go小测试: 你真了解defer>>
原文链接:https://mp.weixin.qq.com/s/yfH0CBnUBmH0oxfC2evKBA
*/
func GetFn() func() {
	fmt.Println("outside")
	return func() {
		fmt.Println("Inside")
	}
}

func TestDeferWp(t *testing.T) {
	defer GetFn()()
	fmt.Println("Here")
}
