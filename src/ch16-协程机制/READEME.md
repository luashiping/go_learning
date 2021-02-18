## 协程（Goroutine）
Go语言中没有线程的概念，只有协程序，也成为Goroutine。相比线程来说，协程更加轻量，一个程序可以随意启动成千上万个goroutine。
goroutine被Go runtime所调度，这一点和线程不一样。也就是说Go语言的并发是有Go自己所调度的，自己决定同时执行多少个goroutine，什么时候执行哪几个。这些对于我们开发者来说完全透明，只需要在编码的时候告诉Go语言要启动几个goroutine，至于如何调度，我们不关心。

## Channel
在Go语言中声明一个channel：
```go
ch := make(chan string)
```
其中chan是一个关键字，表示channel类型。后面的string表示channel里的数据是string类型。通过channel的声明可以看到，chan是一个**集合类型**
一个channel的操作只有两种：**发送和接收**
1. 接收：获取chan中的值，操作符为<-chan
2. 发送：向chan发送值，把值放在chan中，操作符为chan<-
```
小技巧：这里发送和接收的操作符都是<-,只不过位置不同。接收的<-操作符在chan的左侧，发送的操作符在chan的右侧。
```

##### 无缓冲channel
无缓冲channel的容量为0，不能存储任何数据。所以无缓冲的channel只起到数据传输的作用，数据并不会在channel中做任何停留。这也意味着，无缓冲channel的发送和接收操作是同时进行的，它也可以称为同步channel。

##### 有缓冲channel
有缓冲的队列类似一个可阻塞的队列，内部的元素先进先出。通过make函数的第二个参数可以指定channel容量的大小，进而创建一个有缓冲channel。
一个有缓冲channel具备以下特点：
1. 有缓冲channel的内部有一个缓冲队列；
2. 发送操作是队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个goroutine执行，接收操作释放队列的空间；
3. 接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个goroutine执行，发送操作插入新的元素；

##### 关闭channel
