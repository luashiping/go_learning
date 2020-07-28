## golang mac交叉编译linux运行二进制可执行程序
```bash
GOOS=linux GOARCH=amd64 go build hello_world.go
```