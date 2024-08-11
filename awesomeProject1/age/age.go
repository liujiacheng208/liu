package main

import (
	"fmt"
	"net"
	"time"
)

func proces(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
func main() {
	fmt.Println("服务器")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err2 := listen.Accept()
	if err != nil {
		fmt.Println("err2", err2)
	} else {
		fmt.Printf("等待联接成功，con=%v,接收的信息%v\n", conn, conn.RemoteAddr().String())
	}
	go proces(conn)
	time.Sleep(time.Second * 20)
}
