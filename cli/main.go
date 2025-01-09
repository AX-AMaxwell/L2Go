package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*

Purpose:
- Convert integer values, represented as strings from command-line input
- Convert from and to: hex, decimal and binary

Implementation Choices:
- Command-line flag that describes what we're getting - OR - some type of prefix
  to denote the input type (`0b`, `0x`)
- Should have explicit flag to define output type

Inputs:
1. Value converting from (represented as `0x`, `0b` or `0`)
2. Type we want to convert to (hex, decimal, binary)

```
./cli --to-hex 0b0100
./cli -h 0b0001
./cli --to-bin 0x46
./cli -b 0x46
./cli --to-dec 0x46
./cli -d 0x46
```

*/

func main() {
	var (
		f     = parseArgs()
		value int64
		err   error
	)

	switch {
	case isBin(f.value):
		f.value = strings.TrimPrefix(f.value, "0b")
		value, err = strconv.ParseInt(f.value, 2, 64)

	case isHex(f.value):
		f.value = strings.TrimPrefix(f.value, "0x")
		value, err = strconv.ParseInt(f.value, 16, 64)

	default:
		value, err = strconv.ParseInt(f.value, 10, 64)
	}

	if err != nil {
		log.Fatalf("Oh nooooo: %s!", err)
	}

	fmt.Printf("Success ['%d']!\n", value)
}

func isBin(v string) bool {
	return strings.HasPrefix(v, "0b")
}

func isHex(v string) bool {
	return strings.HasPrefix(v, "0x")
}

type flags struct {
	toHex, toBin, toDec bool
	value               string
}

func parseArgs() flags {
	var f flags

	flag.BoolVar(&f.toHex, "to-hex", false, "")
	flag.BoolVar(&f.toHex, "h", false, "")

	flag.BoolVar(&f.toBin, "to-bin", false, "")
	flag.BoolVar(&f.toBin, "b", false, "")

	flag.BoolVar(&f.toDec, "to-dec", false, "")
	flag.BoolVar(&f.toDec, "d", false, "")

	flag.Parse()

	var args = flag.Args()
	if len(args) < 1 {
		fmt.Println("Oh no! We needed a value to convert!")
		os.Exit(1)
	}
	f.value = args[0]

	return f
}
