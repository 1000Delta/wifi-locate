# BUILD STAGE
# ---
FROM golang:1.16-alpine as builder

WORKDIR /go/src/github.com/1000Delta/wifi-locate

COPY . .

RUN apk --no-cache add git && \
apk --no-cache add make

# 减少重复构建依赖安装时间
RUN go env -w GOPROXY=goproxy.io,direct && go mod tidy

RUN make build

# GATEWAY IMAGE
# ---
FROM alpine:latest as gateway

WORKDIR /app

ENV BUILD_DIR=cmd/gateway
ENV CONFIG_DIR=config
ENV CONFIG_FILE=./config.yml

COPY --from=builder /go/src/github.com/1000Delta/wifi-locate/main .
COPY --from=builder /go/src/github.com/1000Delta/wifi-locate/$BUILD_DIR/config .

# CMD [ "ls", "-a" ]
CMD [ "./main", "serve", "--config ${CONFIG_FILE}" ]

#
# LOCATE IMAGE

