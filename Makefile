.PHONY: build

build: clean
	mkdir build && \
	GOOS=darwin GOARCH=arm64 go build -o ./build/httpcode
	cd ./build/ && \
	tar -czvf httpcode-mac-arm64.tar.gz httpcode

clean: 
	rm -rf ./build
