# go-untilread [![](https://circleci.com/gh/844196/go-untilread.svg?style=shield&circle-token=d04f9961cb32d3df4ee68c4ff404b2765d9a22ce)](https://circleci.com/gh/844196/go-untilread) [![](https://goreportcard.com/badge/github.com/844196/go-untilread)](https://goreportcard.com/report/github.com/844196/go-untilread) [![](https://godoc.org/github.com/844196/go-untilread?status.svg)](https://godoc.org/github.com/844196/go-untilread)

untilread provide read function of until separator

## Usage

```go
package main

import (
	"fmt"
	"strings"

	"github.com/844196/go-untilread"
)

func main() {
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
