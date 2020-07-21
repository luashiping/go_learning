## 内置单元测试框架
+ Fail, Error: 该测试失败、该测试继续，其他测试继续执行
+ FailNow, Fatal: 该测试失败、该测试中止，其他测试继续执行



+ 代码覆盖率
  ```bash
  go test -v -cover
  ```
+ 断言
  https://github.com/stretchr/testify