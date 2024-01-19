# 使用Go官方镜像
FROM golang:1.18-alpine

# 设置工作目录
WORKDIR /app

## 复制Go模块依赖文件
#COPY go.mod ./
#COPY go.sum ./

## 下载Go模块依赖
#RUN go mod download

# 复制Go源代码
COPY *.go ./

# 构建Go应用
RUN go build -o /microplate-reader

# 设置运行时的命令
CMD [ "/microplate-reader" ]
