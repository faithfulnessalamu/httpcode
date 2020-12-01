package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	flag "github.com/spf13/pflag"
)

var (
	errNoCode      = errors.New("no code given as argument")
	errInvalidCode = errors.New("invalid code")
	errOutOfRange  = errors.New("code is out of range")
	errNonStandard = errors.New("this is a non-standard response, possibly custom to the server's software")
)

var verbose = flag.BoolP("verbose", "v", false, "display all available info about the status code")

var logerr = log.New(os.Stderr, "", 0)

func main() {
	flag.Parse()

	err := run(os.Args[1:])

	if err != nil {
		logerr.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return printAllCodes()
	}

	code, err := strconv.Atoi(args[len(args)-1])

	if err != nil {
		return errInvalidCode
	}

	if !isInRange(code) {
		return errOutOfRange
	}

	if *verbose {
		return printVerbose(code)
	}

	return printReason(code)
}

func printAllCodes() error {
	var codes []int

	for k := range getAll() {
		codes = append(codes, k)
	}

	sort.Ints(codes)

	for _, code := range codes {
		if err := printVerbose(code); err != nil {
			return err
		}

		fmt.Print("\n")
	}

	return nil
}

func printVerbose(code int) error {
	httpCode, err := getDetails(code)

	if err != nil {
		return explainCorpusError(err)
	}

	fmt.Printf("%d\n%s\n%s\n%s\n", httpCode.code, httpCode.reasonPhrase, httpCode.description, httpCode.moreinfoLink)

	return nil
}

func printReason(code int) error {
	reasonPhrase, err := getReasonPhrase(code)

	if err != nil {
		return explainCorpusError(err)
	}

	fmt.Println(reasonPhrase)
	return nil
}

func explainCorpusError(err error) error {
	if err == errCodeNotFound {
		return errNonStandard
	}

	return err
}

//isInRange returns true if code is in the range [100-599]
func isInRange(code int) bool {
	return 100 <= code && code <= 599
}
