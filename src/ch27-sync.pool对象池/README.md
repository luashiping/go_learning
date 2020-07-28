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

-------------------------------
##### 1.解释测试案例TestSyncPoolWp
输错结果为：
```
1
4
3
2
0
```
原因：
```
(1).源码面前没秘密, pool在get的时候是从切片的最后一个元素倒序依次取出
if x == nil { 
    l.Lock() 
    last := len(l.shared) - 1 
    if last &gt;= 0 {
        x = l.shared[last] 
        l.shared = l.shared[:last] 
    } 
    l.Unlock() 
    if x == nil { 
        x = p.getSlow()        }
    }
}
(2).至于第一个为什么是1，是因为put的时候起初private私有对象为nil，所以会先放到private中，接下去的2，3，4都会放到shared共享池中。在取出的时候会先从private私有对象中获取(获取到1)，取出1后私有对象变为nil，然后再从shared共享池中获取，共享池中获取的方法如(1)说明
```