## 单例模式(懒汉式，线程安全)

```go
type SingletonObj struct{}

var once sync.Once
var obj *SingletonObj

func GetSingletonObj() *SingletonObj{
    once.Do(func(){
        fmt.Println("Create Single obj.")
        obj = &SingletonObj{}
    })
    return obj
}
```