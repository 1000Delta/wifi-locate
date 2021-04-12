module github.com/1000Delta/wifi-locate/gateway

go 1.16

require (
  github.com/gin-gonic/gin v1.6.3
  github.com/1000Delta/wifi-locate/svc-locate v1.0.0
)

replace (
  github.com/1000Delta/wifi-locate/svc-locate => ../svc-locate
)
