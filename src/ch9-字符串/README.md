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


## Go字符串的组成
Go语言在看待Go字符串组成这个问题上，有两种视角。
+ 一种是**字节视角**，和所有其他支持字符串的主流语言一样，**Go语言中的字符串值也就是一个可空的字节序列，字节序列中的字节个数称为该字符串的长度**。**一个个的字节只是孤立数据，不表意**。比如下面的代码，输出了字符串中的每个字节，以及整个字符串的长度：
```go
var s = "中国人"
fmt.Printf("the length of s = %d\n", len(s)) // 9

for i := 0; i < len(s); i++ {
  fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
}
fmt.Printf("\n")
```

+ 如果要表意，就从字符串的另一个角度来看，也就是字符串是由一个可空的字符序列构成
```go
var s = "中国人"
fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3

for _, c := range s {
  fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
}
fmt.Printf("\n")
```
上面这段代码输出了字符串中的字符数量，也输出了这个字符串中的每个字符。**Go采用的是Unicode字符集，每个字符都是一个Unicode字符**，这里输出的以0x4e2d为例，它是汉字“中”在Unicode字符集群中的码点（Code point）

什么是**Unicode码点？** Unicode字符集中的每个字符，都被分配了统一且唯一的字符编号。所谓Unicode码点，就是指将Unicode字符集中的所有字符“排成一队”，字符在这个“队伍”中的位次，就是它在Unicode字符集中的码点。一个码点唯一对应一个字符。

### rune类型
Go使用rune这个类型来表示一个unicode码点。**rune本质上是int32类型的别名类型，它与int32类型是完全等价的**，在Go源码中我们可以看到它的定义：
```go
// $GOROOT/src/builtin.go
type rune = int32
```
**一个unicode码点唯一对应一个unicode字符**。所以可以说**一个rune实例就是一个unicode字符，一个Go字符串也可以被视为rune实例的集合。我们可以通过字符字面值来初始化一个rune变量。** 

+ 在Go中，字符字面值由多种表示法，最常见的是**通过单引号括起的字符字面值**，比如：
```go
'a'  // ASCII字符
'中' // Unicode字符集中的中文字符
'\n' // 换行字符
'\'' // 单引号字符
```
+ **还可以使用Unicode专用的转义字符\u或\U作为前缀，来表示一个Unicode字符。**这里我们要注意：\u后面接四个十六进制数。如果是用四个十六进制数无法表示的Unicode字符，我们可以使用\U，\U后面可以接八个十六进制数来表示一个Unicode字符：
```go
'\u4e2d'     // 字符：中
'\U00004e2d' // 字符：中
'\u0027'     // 单引号字符
```
由于表示码点的rune本质上就是一个整型数，所以我们还可**用整型值来直接作为字符字面值给rune变量赋值**，比如：
```go
'\x27'  // 使用十六进制表示的单引号字符
'\047'  // 使用八进制表示的单引号字符
```

### 字符串字面值
字符串是字符的集合，我们**把表示单个字符的单引号，换为表示多个字符组成的字符串的双引号**就可以了：
```go
"abc\n"
"中国人"
"\u4e2d\u56fd\u4eba" // 中国人
"\U00004e2d\U000056fd\U00004eba" // 中国人
"中\u56fd\u4eba" // 中国人，不同字符字面值形式混合在一起
"\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba" // 十六进制表示的字符串字面值：中国人
```
