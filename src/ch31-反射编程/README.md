## 利用反射编写灵活的代码

按名字访问结构的成员
```go
reflect.ValueOf(*e).FieldByName("Name")
```

按名字访问结构的方法
```go
reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
```

## Struct Tag
```go
type BasicInfo struct{
    Name string `json:"name"`---->// Struct tag
    Age  int    `json:"age"`---->// Struct tag 
}
```