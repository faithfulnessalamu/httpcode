.PHONY: build

build: clean
	mkdir build && \
	GOOS=darwin GOARCH=arm64 go build -o ./build/httpcode-mac-arm64
	cd ./build/ && \
	tar -czvf httpcode-mac-arm64.tar.gz httpcode-mac-arm64

clean: 
	rm -rf ./build
