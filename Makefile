.PHONY: all linux windows mac

all: linux windows mac

linux: clean
	GOOS=linux GOARCH=amd64 go build -o ./build/httpcode-linux-x86_64

windows: clean
	GOOS=windows GOARCH=amd64 go build -o ./build/httpcode-win-x86_64.exe

mac: clean
	GOOS=darwin GOARCH=amd64 go build -o ./build/httpcode-mac-x86_64

clean: 
	rm -rf ./build
	mkdir build
