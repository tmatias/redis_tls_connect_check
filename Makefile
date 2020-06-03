build:
	CGO_ENABLED=0 GO111MODULES=on go build -ldflags="-s -w" .
.PHONY: build

docker-build:
	docker build -t tmatias/redis_tls_connect_check .
.PHONY: docker-build

docker-run:
	docker run tmatias/redis_tls_connect_check
.PHONY: docker-run

docker-push:
	docker push tmatias/redis_tls_connect_check
.PHONY: docker-push

.DEFAULT_GOAL := build
