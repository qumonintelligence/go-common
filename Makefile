.PHONY: checkstyle run

all: build

run:
	go run ./main.go --host 0.0.0.0 --port 8080
