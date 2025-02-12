# 使用Go官方镜像作为构建环境
FROM golang:1.23.6 AS builder

# 设置工作目录
WORKDIR /app

# 复制Go模块文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/news-backend ./cmd/main.go

# 使用Alpine镜像作为运行环境
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/news-backend .

# 暴露端口
EXPOSE 8099

# 设置环境变量
ENV PORT=8099
ENV DB_DSN=proxy:proxyPass@tcp(localhost:3306)/news?charset=utf8mb4%26parseTime=True%26loc=Local

# 运行应用
CMD ["./news-backend"]