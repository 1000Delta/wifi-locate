# BUILD STAGE
# ---
FROM golang:alpine as builder

WORKDIR /go/src/github.com/1000Delta/wifi-locate

RUN apk --no-cache add git && \
  apk --no-cache add make

COPY . .

# 减少重复构建依赖安装时间
RUN go env -w GOPROXY=goproxy.io,direct && \
  go env -w CGO_ENABLED=0 && \
  go env -w GOARCH=amd64 && \
  go env -w GOOS=linux && \
  make build

#
# GATEWAY IMAGE
# ---
FROM alpine:latest as gateway

WORKDIR /app

RUN apk add --no-cache ca-certificates

# 多 APP 修改环境变量即可
ENV APP_NAME=gateway
ENV CMD_DIR=cmd/${APP_NAME}
ENV CONFIG_DIR=config
ENV CONFIG_FILE=./config.yml

COPY --from=builder /go/src/github.com/1000Delta/wifi-locate/bin/${APP_NAME} .
COPY --from=builder /go/src/github.com/1000Delta/wifi-locate/$CMD_DIR/config .

# CMD [ "ls", "-a" ]
CMD [ "sh", "-c", "./${APP_NAME} serve --config ${CONFIG_FILE}" ]

#
# LOCATE IMAGE
# ---
FROM alpine:latest as locate

WORKDIR /app

RUN apk add --no-cache ca-certificates

# 多 APP 修改环境变量即可
ENV APP_NAME=locate
ENV CMD_DIR=cmd/${APP_NAME}
ENV CONFIG_DIR=config
ENV CONFIG_FILE=./config.yml

COPY --from=builder /go/src/github.com/1000Delta/wifi-locate/bin/${APP_NAME} .
COPY --from=builder /go/src/github.com/1000Delta/wifi-locate/$CMD_DIR/config .

# CMD [ "ls", "-a" ]
CMD [ "sh", "-c", "./${APP_NAME} serve --config ${CONFIG_FILE}" ]

