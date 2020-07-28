package func_test

import (
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
