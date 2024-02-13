package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个 quit 通道，用于接收终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthzHandler)

	// 启动一个新的 Go 协程来等待接收 quit 通道中的信号，并在收到信号时执行清理操作
	go func() {
		<-quit
		log.Println("Server is shutting down...")

		// 在收到终止信号后执行清理操作，例如关闭数据库连接等

		log.Println("Server stopped gracefully")
		os.Exit(0)
	}()

	log.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. 将 request 中的 header 写入 response header
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// 2. 读取当前系统环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	// 3. 记录访问日志
	log.Printf("Client IP: %s, HTTP Status Code: %d\n", r.RemoteAddr, http.StatusOK)

	// 返回响应
	fmt.Fprintln(w, "Hello, World!")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// 4. 访问 /healthz 时返回200
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
