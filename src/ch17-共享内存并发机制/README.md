### 同步原语

#### sync.Mutex
互斥锁，顾名思义，指的是在同一时刻只有一个协程执行某段代码，其他协程执行完毕后才能继续执行。
互斥锁的使用非常简单，它只有两个方法 Lock 和 Unlock，代表加锁和解锁。当一个协程获得 Mutex 锁后，其他协程只能等到 Mutex 锁释放后才能再次获得锁。
Mutex 的 Lock 和 Unlock 方法总是成对出现，而且要确保 Lock 获得锁后，一定执行 UnLock 释放锁，所以在函数或者方法中会采用 defer 语句释放锁，如下面的代码所示：
```go
func add(i int) {
   mutex.Lock()
   defer mutex.Unlock()
   sum += i
}
```
这样可以确保锁一定会被释放，不会被遗忘。

#### sync.RWMutex

#### sync.WaitGroup
