## Go的错误机制
#### 与其他编程语言的主要差异：
1. 没有异常机制
2. error类型实现了error接口
type error interface{
    Error() string
}
3. 可以通过erros.New来快速创建错误实例
errors.New("xxx")