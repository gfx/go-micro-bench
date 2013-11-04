package micro_bench

import (
	"bytes"
	"os"
	"testing"
)

func BenchmarkStrCatWithBytes(b *testing.B) {
	var buffer bytes.Buffer

	for i := 0; i < b.N; i++ {
		buffer.WriteString("foobar\n")
	}

	file, err := os.OpenFile("/dev/null", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	file.WriteString(buffer.String())
}

func BenchmarkStrCatWithString(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s += "foo\n"
	}

	file, err := os.OpenFile("/dev/null", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	file.WriteString(s)
}
