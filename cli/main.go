package main

import (
	"fmt"
	"log"
)

/*

Purpose:
- Convert integer values, represented as strings, from command-line input
- Convert from and to: hex, decimal and binary

Implementation Choices:
- Command-line flag that describes what we're getting - OR - some type of prefix
  to denote the input type (`0b`, `0x`)
- Should have explicit flag to define output type

Inputs:
1. Value converting from (represented as `0x`, `0b` or `0`) (Positional)
2. Type we want to convert to (hex, decimal, binary) (Flags)

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
		f        Args
		valueInt int64
		valueStr string
		err      error
	)

	if f, err = ParseArgs(); err != nil {
		log.Fatalf("Failed to parse command-line inputs: %s.", err)
	}

	if valueInt, err = ParseInt(f.value); err != nil {
		log.Fatalf(
			"Failed to parse provided input value ['%s'] as an integer: %s.",
			f.value,
			err,
		)
	}

	if valueStr = FormatInt(valueInt, f); len(valueStr) == 0 {
    
	}

	fmt.Println(valueStr)
}
