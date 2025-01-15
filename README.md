# httpcode

Get info about http status codes right at your terminal.

## Installation

### Homebrew

```shell
$ brew tap faithfulnessalamu/tap
$ brew install httpcode
```

### Build from source

```shell
$ git clone https://github.com/faithfulnessalamu/httpcode
$ cd httpcode
$ go install
$ httpcode --help
```

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
