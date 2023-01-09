## panic
1. paninc用于不可以恢复的错误
2. panic退出前回执行defer指定的内容


## panic vs. os.Exit
1. os.Exit退出时不会调用defer指定的函数
2. os.Exit退出时不输出当前调用栈信息

## 当心！recover成为恶魔
1. 形成僵尸服务进程，导致health check失效
2. “Let it Crash”往往是我们恢复不确定性错误的最好方法