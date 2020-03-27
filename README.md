# A simple data package deal in golang.

# Usage
```golang
// server 
package main

import "github.com/eavesmy/packet_golang"
import "fmt"
import "net"
import "bufio"
import "encoding/json"
import "io"

type message struct {
	Msg string
}

func main(){
	ln, _ := net.Listen("tcp","8080")

	for {
		conn,_ := ln.Accept()
		go handler(conn)
	}
}

func handler(conn *net.Conn){
	reader := bufio.NewReader(conn)

	// create new unpack struct and set buffer size.
	unpack := packet_golang.NewUnpack(1024)
	for {
		data, _:= reader.ReadBytes('\n')

		if err := unpack.Deal(data); err == io.EOF {
			continue
		}
		
		msg := &message{}
		json.Unmarshal(unpack.Bytes(), msg)
		
		fmt.Println(msg) // {"hello world"}
	}
}

```    


```golang
// client
package main

import "net"
import "github.com/eavesmy/packet_golang"

type message struct {
	Msg string
}

func main(){
	msg := &{Msg: "helow world"}
	// create pack and set package id(uint32)
	pack := packet_golang.NewPacket(0,msg)
	
	conn,_ := net.Dial("tcp",":8080')
	conn.Write(pack)

}
```

# Struct
```
[ pid ] uint32    4 bit    
[ len ] uint16    2 bit
[ body ] []byte   ...
[ \n ]  bit       1 bit
```
Is simple.

# Install
```go get github.com/eavesmy/packet_golang```
