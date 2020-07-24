## 更快的json解析
+ EasyJSON采用代码生成而非反射
  ![avatar](性能评测.png)
+ 安装
   ```bash
   $ go get -u github.com/mailru/easyjson/...
   ```
+ 使用
  ```bash
  $GOPATH/bin/easyjson -all <结构定义.go>
  ```

