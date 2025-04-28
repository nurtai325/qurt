VERSION=1.0.0
NAME=qurt
PORT=8080

all:
	make build
	make run

build:
	docker build -t $(NAME):$(VERSION) .

run:
	docker run --rm -p $(PORT):80 --name $(NAME) $(NAME):$(VERSION)

.PHONY: all build run
