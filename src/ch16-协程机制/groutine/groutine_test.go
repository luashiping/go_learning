package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		// go func() {
		// 	fmt.Println(i)
		// }()
	}
	time.Sleep(time.Millisecond * 50)
}

func TestGroutine2(t *testing.T) {
	go fmt.Println("宝仔love")
	fmt.Println("我是 main goroutine")
	time.Sleep(time.Second)
}

// channel解决goroutine之间的通信问题
func TestChannel(t *testing.T) {
	ch := make(chan string)
	go func() {
		fmt.Println("飞雪无情")
		// time.Sleep(time.Second * 5)
		// 向通道ch发送值
		ch <- "goroutine 完成"
	}()
	fmt.Println("我是 main goroutine")
	// time.Sleep(time.Second * 10)
	//从通道ch接收值，如果ch中没有值，则阻塞等待到ch中有值可以接收为止
	v := <-ch
	fmt.Println("接收到的chan中的值为: ", v)
}

func TestCacheChannel(t *testing.T) {
	cacheCh := make(chan int, 5)
	res1 := <-cacheCh
	fmt.Println(res1)
	cacheCh <- 2
	cacheCh <- 3
	cacheCh <- 4
	fmt.Println("cache容量:", cap(cacheCh), "元素个数为:", len(cacheCh))
	cacheCh <- 5
	fmt.Println("cache容量:", cap(cacheCh), "元素个数为:", len(cacheCh))
	cacheCh <- 6
	fmt.Println("cache容量:", cap(cacheCh), "元素个数为:", len(cacheCh))
	// res1 := <-cacheCh
	// res2 := <-cacheCh
	fmt.Println("cache容量:", cap(cacheCh), "元素个数为:", len(cacheCh))
	fmt.Println(res1)
}
