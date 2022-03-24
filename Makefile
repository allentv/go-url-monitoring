run:
	go run main.go

test:
	go test ./... -count=1 -race -cover

build:
	docker build -t go-url-monitoring .
