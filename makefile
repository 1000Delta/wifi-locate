.PHONY: gateway
gateway:
	cd gateway && go run . &

.PHONY: locate
locate:
	cd svc-locate && go run . &

SERVICE:=$$(for f in $$(ls); do if [ -d $$f ]; then echo $$f; fi; done)

.PHONY: close
close:
	@for svc in $(SERVICE); do echo 'close' $$svc; kill $$(ps -A | grep $$svc | awk '{print $$1}'); done

list:
	echo ${SERVICE}

run: gateway locate