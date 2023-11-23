all: build

build:
	GOOS=js GOARCH=wasm go build -o teste.wasm main.go

server:
	python3 -m http.server