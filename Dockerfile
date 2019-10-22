FROM alpine:latest
MAINTAINER 1907004005@qq.com
RUN apk update \
    && apk add --no-cache bash
COPY hotlist_server /usr/bin/
ENTRYPOINT ["hotlist_server"]
