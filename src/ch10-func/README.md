## 函数是一等公民
#### 与其他编程语言的主要差异
1. 可以有多个返回值
2. 所有参数都是值传递：slice，map，channel会有传引用传递
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

#### 函数和方法
在Go语言中，方法和函数是两个概念，但又非常相似，不同点在于方法必须有一个接收者，这个接收者是一个类型，这样方法就和这个类型绑定在一起，成为这个类型的方法

-------------------------------
##### 1.解释测试案例TestDeferWp
输错结果为：
```
outside
Here
Inside
```
原因：
```
需要先把func入栈，故先执行return前面的outside
```