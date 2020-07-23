## Pipe-Filter 架构
![avatar](pipe-filter架构.png)

## Pipe-Filter 模式
+ 非常适合数据处理及数据分析系统
+ Filter封装数据处理功能
+ 松耦合：Filter只跟数据(格式)耦合
+ Pipe用于连接Filter传递数据或者在异步处理过程中缓冲数据流
  **进程内部同步调用时，pipe演变为数据在方法调用间传递**


## 示例
![avatar](示例.png)