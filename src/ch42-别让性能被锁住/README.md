## sync.Map
+ 适合读多写少，且key相对稳定的环境
+ 采用了空间换时间的方案，并且采用指针的方式间接实现值的映射，所以存储空间会比built-in map大
  
https://my.oschina.net/qiangmzsx/blog/1827059

## 别让性能被"锁"住
+ 减少锁的影响范围
+ 减少发生锁冲突的概率
  + sync.Map
  + ConcurrentMap
+ 避免锁的使用
  + LAMX Disruptor: https://martinfowler.com/articles/lmax.html
