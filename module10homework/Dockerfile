# 使用 golang 作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到工作目录中
COPY . .

# 构建 Go 程序
RUN go build -o main .

# 设置容器的入口点（启动命令）
CMD ["./main"]
