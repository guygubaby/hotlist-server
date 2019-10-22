FROM alpine:latest
MAINTAINER 1907004005@qq.com
RUN apk update \
    && apk add --no-cache bash \
    && cp -rf /usr/share/zoneinfo/Hongkong /etc/localtime
WORKDIR /usr/app
COPY hotlist_server /usr/bin/
ENTRYPOINT ["hotlist_server"]
