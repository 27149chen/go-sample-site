FROM golang:1.17.5

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /build

COPY go.mod .
COPY go.sum .

# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static'" -o goserver main.go


FROM alpine:3.15.0
MAINTAINER lou

WORKDIR /app

# https://wiki.alpinelinux.org/wiki/Setting_the_timezone
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo Asia/Shanghai  > /etc/timezone && \
    apk del tzdata

COPY --from=0 /build/goserver .

ENTRYPOINT ["/app/goserver"]
