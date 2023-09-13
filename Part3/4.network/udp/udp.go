package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务器地址和端口
	serverAddr := "47.93.236.205:30001"

	// 学号和姓名（注意按照要求格式化）
	studentInfo := "20227291 李卓远\n"

	// 解析服务器地址
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("无法解析服务器地址:", err)
		return
	}

	// 创建UDP连接
	conn, err := net.DialUDP("udp", nil, udpAddr)
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

	fmt.Println("数据已发送到服务器:", studentInfo)
}
