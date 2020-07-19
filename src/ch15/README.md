## package
1. **基本复用模块单元**
   以首字母大写来表明可被外代码访问
2. **代码的package名称可以和所在的目录名称不一致**
3. **同一目录的Go代码的package名称要保持一致**
4. **通过go get来获取远程依赖**
   + go get -u 强制从网络更新远程依赖
5. **注意代码在GitHub上的组织形式，以适应go get**
   + 直接以代码路径开始，不要有src

示例：https://github.com/easierway/concurrent_map

## init方法
+ 在main被执行之前，所有依赖的package init方法都会被执行
+ 不同包的init函数按照包导入的依赖关系决定执行顺序
+ 每个包可以有多个init函数
+ 包的每个源文件也可以有多个init函数


## Go未解决的依赖问题
1. 同一环境下，不同项目使用同一包的不同版本
2. 无法管理对包的特定版本依赖

## Vendor路径
随着Go 1.5 release版本的发布，vendor目录被添加到除了GOPATH和GOROOT之外的依赖目录查找的解决方案。

##### 查找依赖包路径的解决方案如下：
1. 当前包下的vendor目录
2. 向上级目录查找，直到找到src下的vendor目录
3. 在GOPATH下面查找依赖包
4. 在GOROOT目录下查找

## 常用的依赖管理工具
+ godep http://github.com/tools/godep
+ glide http://github.com/Masterminds/glide
+ dep   http://github.com/golang/dep
