## GC友好的代码

### 避免内存分配和复制
+ 复杂对象尽量传递引用
  + 数组的传递
  + 结构体传递

### 打开GC日志
只要在程序之前加上环境变量GODEBUG=gctrace=1
如:
```bash
GODEBUG=gctrace=1 go test -bench=.
GODEBUG=gctrace=1 go run main.go
```
日志详细信息参考：https://godoc.org/runtime


### go tool trace
**普通程序输出trace信息**
```go
package main
import (
  "os"
  "runtime/trace"
)

func main() {
  f, err := os.Create("trace.out")
  if err != nil {
    panic(err)
  }
  defer f.Close

  err = trace.Start(f)
  if err != nil {
    panic(err)
  }
  defer trace.Stop()
}
```

**测试程序输出trace信息**
```bash
go test -trace=trace.out
```

**可视化trace信息**
```bash
go tool trace trace.out
```

### 避免内存分配和复制
+ 初始化至合适的大小
  + 自动扩容是有代价的
+ 复用内存