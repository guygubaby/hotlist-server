FROM alpine:latest
MAINTAINER 1907004005@qq.com
RUN apk update &&\
    apk add curl bash
WORKDIR /usr/app
COPY hotlist-server /usr/bin/
ENTRYPOINT ["hotlist-server"]
