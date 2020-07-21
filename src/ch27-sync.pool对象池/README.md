## sync.Pool对象获取
+ 尝试从私有对象获取
+ 私有对象不存在，尝试从当前Processor的共享池获取
+ 如果当前Processor共享池也是空的，那么就尝试去其他的Processor共享池获取
+ 如果所有池子都是空的，最后就用用户指定的New函数产生一个新的对象返回

## sync.Pool对象的放回
+ 如果私有对象不存在则保存为私有对象
+ 如果私有对象存在，放入当前Processor子池的共享池中

## sync.Pool对象的生命周期
+ GC会清除sync.pool缓存的对象
+ 对象的缓存有效期为下一次GC之前

## sync.Pool总结
+ 适合于通过复用，降低复杂对象的创建和GC代价
+ 协程安全，<font color=red>会有锁的开销</font>
+ <font color=red>生命周期受GC影响，不适合做连接池等需自己管理生命周期的资源的池化</font>