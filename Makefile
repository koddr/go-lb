.PHONY: clean test security build run

APP_NAME = lb
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

linter:
	golangci-lint run

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	BACKENDS_LIST="http://127.0.0.1:5051,http://127.0.0.1:5052" \
	LB_PORT=3030 \
	$(BUILD_DIR)/$(APP_NAME)

docker.run: docker.network docker.lb

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.lb.build:
	docker build -t go-lb .

docker.lb: docker.lb.build
	docker run --rm -d \
		--name dev-lb \
		--network dev-network \
		-p 5050:5050 \
		go-lb

docker.stop: docker.stop.lb

docker.stop.lb:
	docker stop dev-lb