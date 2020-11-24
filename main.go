package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	flag "github.com/spf13/pflag"
)

var (
	errNoCode      = errors.New("no code given as argument")
	errInvalidCode = errors.New("invalid code")
	errOutOfRange  = errors.New("code is out of range")
)

var verbose = flag.BoolP("verbose", "v", false, "display all available info about the status code")

var logerr = log.New(os.Stderr, "", 0)

func main() {
	flag.Parse()

	if len(os.Args) <= 1 {
		logerr.Fatal(errNoCode)
	}

	code, err := strconv.Atoi(os.Args[1])
	if err != nil {
		logerr.Fatal(errInvalidCode)
	}

	if !isInRange(code) {
		logerr.Fatal(errOutOfRange)
	}

	if *verbose {
		printVerbose(code)
	} else {
		printReason(code)
	}
}

func printVerbose(code int) {
	httpCode, err := getDetails(code)
	if err != nil {
		logerr.Fatal(explainCorpusError(err))
	}
	fmt.Printf(`%d\n%s\n%s\n%s\n`, httpCode.code, httpCode.reasonPhrase, httpCode.description, httpCode.moreinfoLink)
}

func printReason(code int) {
	reasonPhrase, err := getReasonPhrase(code)
	if err != nil {
		logerr.Fatal(explainCorpusError(err))
	}
	fmt.Println(reasonPhrase)
}

func explainCorpusError(err error) error {
	if err == errCodeNotFound {
		return errors.New("this is a non-standard response, possibly custom to the server's software.")
	}
	return err
}

//isInRange returns true if code is in the range 100-599
func isInRange(code int) bool {
	return 100 <= code && code <= 599
}
