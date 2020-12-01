# httpcode

Get info about http status codes right at your terminal.

## Download

See [releases](https://github.com/thealamu/httpcode/releases) page to download binaries.

## Usage

```shell
$ httpcode
100
Continue
This interim response indicates that everything so far is OK and that the client should continue the request, or ignore the response if the request is already finished.
https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/100

101
Switching Protocols
This code is sent in response to an Upgrade request header from the client, and indicates the protocol the server is switching to.
https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/101

102
Processing
This code indicates that the server has received and is processing the request, but no response is available yet.
...

$ httpcode 418
I'm a teapot

$ httpcode -v 418
418
I'm a teapot
The server refuses the attempt to brew coffee with a teapot.
https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418
```

## Build :hammer:

You can install with Go:

```shell
$ go get github.com/thealamu/httpcode
```

or build yourself:

```shell
$ git clone https://github.com/thealamu/httpcode
$ cd httpcode
$ go install
$ httpcode --help
```
