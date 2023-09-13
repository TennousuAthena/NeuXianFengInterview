package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务器地址和端口
	serverAddr := "47.93.236.205:30000"

	// 学号和姓名（注意按照要求格式化）
	studentInfo := "20227291 李卓远\n"

	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return
	}
	defer conn.Close()

	// 发送学号和姓名
	_, err = conn.Write([]byte(studentInfo))
	if err != nil {
		fmt.Println("发送数据时出错:", err)
		return
	}

	// 接收并打印服务器返回的信息
	buffer := make([]byte, 1024)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("接收数据时出错:", err)
		return
	}

	response := string(buffer[:bytesRead])
	fmt.Println("服务器返回的信息:", response)
}
