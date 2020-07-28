package object_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	// v2, _ := pool.Get().(int)
	// fmt.Println(v2)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(101)
	pool.Put(102)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
来自蔡超老师的<<Go小测试: 你真了解sync.Pool>>
原文链接:https://mp.weixin.qq.com/s/VF8fpREG5oK-fw1DWldFTQ
*/
func TestSyncPoolWp(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)

	for i := 0; i < 5; i++ {
		fmt.Println(p.Get())
	}
}
