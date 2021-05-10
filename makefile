SVC_DIR:=cmd
BIN_DIR:=bin
SVC_LIST:=$$(for f in $$(ls ${SVC_DIR}); do echo $$f; done)

build:
	@for svc in $(SVC_LIST); do echo 'build' $$svc; CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./${BIN_DIR}/$$svc ./${SVC_DIR}/$$svc; done

.PHONY: close
close:
	@for svc in $(SVC_LIST); do echo 'close' $$svc; kill $$(ps -A | grep $$svc | awk '{print $$1}'); done

list:
	@echo ${SVC_LIST}

run:
	docker-compose up