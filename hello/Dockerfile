# 使用官方 Go 镜像作为基础镜像
FROM golang:1.23-alpine AS builder

# Dockerfile 中定义 ARG
ARG API_VERSION

# 使用 ARG 变量
ENV API_VERSION=${API_VERSION}

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum（从项目根目录）
COPY ../go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制子目录代码和必要的共享模块
COPY hello/ ./hello/

# 构建可执行文件（在子目录内）
WORKDIR /app/hello
RUN go build -o main .

# 使用镜像运行应用程序
FROM alpine:3.18

# 设置工作目录
WORKDIR /app

# 复制构建阶段生成的可执行文件到运行镜像
COPY --from=builder /app/hello/main .
RUN mkdir -p /app/etc
COPY hello/etc/ /app/etc/

# 暴露服务端口（根据实际服务设置）
EXPOSE 8888

# 设置启动命令
CMD ["./main"]
