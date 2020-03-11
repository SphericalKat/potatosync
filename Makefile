.PHONY: build
build:
	mkdir bin || echo bin already exists
	@echo building artifacts
	@GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/potatosync . && echo DONE

.PHONY: dep
dep:
	go mod vendor
	go mod verify

.PHONY: img-build
img-build:
	docker image build -t atechnohazard/potatosync .

img-push:
	docker push atechnohazard/potatosync:latest

.PHONY: docs
docs:
	apidoc -i controllers -o docs

.PHONY: setup-buildx
setup-buildx:
	update-binfmts --enable
	docker buildx create --driver docker-container --use
	docker buildx inspect --bootstrap
	docker buildx ls
	docker login -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD"

.PHONY: img-buildx
img-buildx:
	docker buildx build --platform linux/arm,linux/amd64,linux/x86 --progress plain --pull -t "basmakes/potatosync" --push .
