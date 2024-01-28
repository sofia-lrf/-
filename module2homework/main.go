package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthzHandler)
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
