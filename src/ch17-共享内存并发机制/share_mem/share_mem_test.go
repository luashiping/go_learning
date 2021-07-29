package share_mem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			// counter在不同的协程里去做自增，导致并发竞争条件，传统意义上来说就是非线程安全
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// 共享内存锁保护
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

//
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}

func TestSummation(t *testing.T) {
	// 声明一个互斥锁mutex
	var mutex sync.Mutex

	var wg sync.WaitGroup
	var sum int
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(a int) {
			defer mutex.Unlock()
			mutex.Lock()
			// sum += a被加锁保护
			sum += a
			wg.Done()
		}(12)
	}
	wg.Wait()
	fmt.Println("和为:", sum)
}
