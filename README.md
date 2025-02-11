# 投资新闻展示平台

## 简介
投资新闻展示平台是一个基于Gin框架的API后端服务，用于提供新闻相关的数据接口。前端使用Vue.js 2.0，结合Vuex、Vue Router、Element UI等技术栈，后端则使用Golang，并通过GORM操作MySQL数据库，Redis作为缓存。

## 技术栈
- **后端**: Go 1.23.6, Gin, GORM, go-redis, go-sql-driver/mysql, go-redis/redis-stack-go
- **前端**: Vue.js 2.0, Vuex, Vue Router, Element UI
- **数据库**: MySQL
- **缓存**: Redis

## 安装步骤

### 本地安装

1. **克隆项目**
   ```bash
   git clone https://github.com/teddymail/news-api.git
   cd news-api
   ```

2. **配置环境变量**
   - 复制 `.env.example` 文件并重命名为 `.env`。
   - 修改 `.env` 文件中的数据库地址和其他配置项。
     ```bash
     cp .env.example .env
     ```

3. **安装依赖**
   ```bash
   go mod tidy
   ```

4. **创建数据库表结构**
   ```bash
   go run cmd/migrations/main.go
   ```

5. **启动后端服务**
   ```bash
   go run cmd/main.go
   ```

### 使用Docker安装

1. **构建Docker镜像**
   ```bash
   docker build -t news-api:latest .
   ```

2. **运行Docker容器**
   ```bash
   docker run -d -p 8099:8099 --env-file .env --name news-api-container news-api:latest
   ```

## 环境变量配置

- `DB_DSN`: 数据库连接字符串，格式为 `username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local`
- `REDIS_ADDR`: Redis服务器地址
- `REDIS_PASSWORD`: Redis服务器密码
- `PORT`: 应用监听的端口号

## 贡献指南

欢迎任何形式的贡献！请参考 [CONTRIBUTING.md](CONTRIBUTING.md) 文件。

## 联系信息

- **GitHub**: [teddymail/news-api](https://github.com/teddymail/news-api)
- **Email**: iyuekang@gmail.com
- **Blog**: [https://yuekang.org.cn/](https://yuekang.org.cn/)