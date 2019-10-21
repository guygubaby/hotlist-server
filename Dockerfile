FROM golang:1.13.2-alpine3.10 as build
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR /go/release

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o hotlist_server main.go

FROM scratch as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/hotlist_server /
EXPOSE 5000
CMD ["/hotlist_server"]
