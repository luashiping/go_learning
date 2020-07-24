## 准备工作
+ 安装graphviz
  ```bash
  brew install graphviz
  ```


## Go支持的多种Profile
```bash
go help testflag
## 查看cpu profile，prof为编译出来的二进制程序
go tool pprof prof cpu.prof
```
https://golang.org/src/runtime/pprof/pprof.go

```
