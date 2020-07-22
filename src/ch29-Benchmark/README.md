## Benchmark
```go
func BenchmarkConcatStringByAdd(b *tesing.B){
    //与性能测试无关的代码
    b.ResetTimer()
    for i := 0;i < b.N; i++{
        //测试代码
    }
    b.StopTimer()
    //与性能测试无关的代码

}
```

```bash
$ go test -bench=. -benchmem
-bench=<相关benchmark测试>
-benchmem 打印内存分配，通过这些信息可以知道代码为什么快，为什么慢
Windows下使用go test命令行时，-bench=.应该改为-bench="."
```
