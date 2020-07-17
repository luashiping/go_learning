## 字符串
#### 与其他编程语言的主要差异：
1. string是数据类型，零值不是空而是空字符串,不是引用或指针类型
2. string是只读的byte slice, len函数返回它包含的byte数
3. string的byte数组可以存放任何数据

## Unicode UTF8
1. Unicode是一种字符集(code point)
2. UTF8是unicode的存储实现(转换为字节序列的规则)

## 常用字符串函数
1. strings 包(https://golang.org/pkg/strings/)
2. strconv 包(https://golang.org/pkg/strconv/)