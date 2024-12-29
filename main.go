package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

// 获取本机局域网 IP
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}

func main() {
	// 设置静态文件服务
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// 获取本机局域网 IP
	localIP := getLocalIP()
	port := "8080"
	addr := fmt.Sprintf("%s:%s", localIP, port)

	fmt.Printf("服务器启动在: http://%s\n", addr)

	// 启动服务器
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
