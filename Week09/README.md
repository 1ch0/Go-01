1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

### 1. 解答
#### fix length
包大小固定，但容易出现浪费、包长不灵活的情况。

#### delimiter based
包使用特殊符号做为定界符，判断是不是一个请求，定界符可能会出现二义性。

#### length field based frame decoder
包使用固定长度内容加变长内容组成，包长灵活。

### 2. 解答
