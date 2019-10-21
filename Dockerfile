FROM golang:1.13.2 as build
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR /go/release

COPY . .

RUN go build -o hotlist_server main.go

FROM scratch as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/hotlist_server /
EXPOSE 5000
CMD ["./hotlist_server"]


