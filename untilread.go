// Package untilread provide read function of until separator
package untilread

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Do read io.Reader until separator chars, read string pass to anonymous function.
// Anonymous function of argument should return err or nil.
// if read failed or anonymous function return err, "Do" return err.
func Do(r io.Reader, sep string, f func(string) error) error {
	br := bufio.NewReader(r)
	var buf bytes.Buffer

	for {
		byt, err := br.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = buf.WriteByte(byt)
		if err != nil {
			return err
		}

		if strings.HasSuffix(buf.String(), sep) == false {
			continue
		}

		err = f(buf.String())
		if err != nil {
			return err
		}

		buf.Reset()
	}

	return nil
}
