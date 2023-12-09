FROM debian:stretch-slim

#工作目录
WORKDIR /go-blog

# 从builder镜像中把配置文件拷贝到当前目录
COPY ./conf/config.ini /go-blog/conf/config.ini
COPY ./runtime/logs /go-blog/runtime/logs
COPY ./main /go-blog/main
RUN chmod -R 777 /go-blog


# 需要运行的命令
ENTRYPOINT ["/go-blog/main"]
