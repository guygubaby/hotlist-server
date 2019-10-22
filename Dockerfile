FROM alpine:latest
MAINTAINER 1907004005@qq.com
RUN apk update \
    && apk add --no-cache bash
WORKDIR /usr/app
COPY hotlist-server /
ENTRYPOINT ["./usr/app/hotlist-server"]
