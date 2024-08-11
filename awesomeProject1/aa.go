package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("statc")
	conm, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	fmt.Println("连接成功", conm)
	reader := bufio.NewReader(os.Stdin)
	str, err2 := reader.ReadString('\n')
	if err2 != nil {
		fmt.Println(err2)
	}
	n, err3 := conm.Write([]byte(str))
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Printf("发送成功，%v", n)
}
