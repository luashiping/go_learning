## channel的关闭
+ 向关闭的channel发送数据，会导致panic
+ v, ok <-ch; ok为bool值，true表示正常接受，false表示通道关闭
+ 所有的channel接受者都会在channel关闭时，立刻从阻塞等待中返回且上诉ok值为false。这个广播机制常被利用，进行向多个订阅者发送信号。
  如：退出信号