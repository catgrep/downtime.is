.PHONY: server proxy publish deploy

test:
	go test -count=1 -v -failfast ./...

server:
	docker compose --progress=plain up server --remove-orphans --build

push:
	./push-to-ghcr.sh

deploy:
	./deploy-site.sh
