package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	})

	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests in seconds",
		Buckets: prometheus.ExponentialBuckets(0.01, 2, 15),
	})	
)

func init() {
	// 注册 Prometheus 指标
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
}

func main() {
	// 创建一个 quit 通道，用于接收终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthzHandler)
	http.Handle("/metrics", promhttp.Handler())

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
	startTime := time.Now()

	// 添加 0-2 秒的随机延时
	randomDelay := time.Duration(rand.Intn(3)) * time.Second
	time.Sleep(randomDelay)

	// 记录请求次数
	requestsTotal.Inc()

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

	// 计算请求耗时并记录到指标中
	duration := time.Since(startTime).Seconds()
	requestDuration.Observe(duration)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// 4. 访问 /healthz 时返回200
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
