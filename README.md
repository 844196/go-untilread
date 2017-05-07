# go-untilread

untilread provide read function of until separator

## Usage

```go
package main

import (
	"fmt"
	"strings"

	"github.com/844196/go-untilread"
)

func Example() {
	src := "こんにちは、ジャパリパーク。さようなら、世界。"
	sep := "。"
	ir := strings.NewReader(src)

	untilread.Do(ir, sep, func(s string) error {
		fmt.Println(s)
		return nil
	})
	// Output:
	// こんにちは、ジャパリパーク。
	// さようなら、世界。
}
```

## Installation

```console
$ go get github.com/844196/go-untilread
```
