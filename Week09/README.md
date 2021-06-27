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
自己不太懂，看的别人的代码
fix length
#### client

```go
func client_tcp_fix_length(conn net.Conn) {
	sendByte := make([]byte, 1024)
	sendMsg := "{\"test01\":1,\"test02\",2}"
	for i := 0; i < 1000; i++ {
		tempByte := []byte(sendMsg)
		for j := 0; j < len(tempByte) && j < 1024; j++ {
			sendByte[j] = tempByte[j]
		}
		_, err := conn.Write(sendByte)
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("send over once")
	}
}
```



#### server

```go
func server_tcp_fix_length(conn net.Conn) {
	fmt.Println("server, fix length")
	const (
		BYTE_LENGTH = 1024
	)
	for {
		var buf = make([]byte, BYTE_LENGTH)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("client data :", string(buf))
	}
}
```

