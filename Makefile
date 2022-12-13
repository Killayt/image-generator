.PHONY: build
build:
	rm -Rf build && mkdir build && go build -o build/img_generator -v ./cmd

.PHONY: run
run: 
	go run cmd/main.go