FROM golang:1.18-rc

LABEL maintainer="杨圆建 <624508914@qq.com>"

# 安装必要的软件包和依赖包
USER root
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    git \
    vim

# 添加环境变量
RUN go env -w GOPROXY=https://goproxy.cn/,direct

# 安装 air 工具（热加载）
RUN go install github.com/cosmtrek/air@v1.15.1

# 安装 goctl
USER root
RUN go install github.com/zeromicro/go-zero/tools/goctl@master

# $GOPATH/bin 添加到环境变量中
ENV PATH $GOPATH/bin:$PATH

# 自动安装 go-zero 环境依赖
USER root
RUN goctl env check -i -f
# 赋权执行权限
RUN chmod +x /go/bin/*

# 清理垃圾
USER root
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

# 设置工作目录
WORKDIR /usr/src/app

EXPOSE 8000
EXPOSE 9000