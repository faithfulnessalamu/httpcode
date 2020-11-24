# httpcode

Get info about http status codes right at your terminal

## Usage

```shell
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
