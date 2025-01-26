BINARY_NAME=prime-backoff
DOCKER_IMAGE=prime-backoff

.PHONY: build test run clean docker-build docker-run integration-test

build:
	go build -o ${BINARY_NAME} .

test:
	go test -v ./...

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm -f ${BINARY_NAME}

docker-build:
	docker build -t ${DOCKER_IMAGE} .

docker-run: docker-build
	docker run --rm ${DOCKER_IMAGE}

integration-test:
	docker-compose up --build --abort-on-container-exit