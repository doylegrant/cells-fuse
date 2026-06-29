package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFmtSprintfByteRange(b *testing.B) {
	chunkOffset := int64(1048576)
	dataLen := int64(4096)

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("bytes=%d-%d", chunkOffset, chunkOffset+dataLen-1)
		}
	})

	b.Run("ConcatStrconv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = "bytes=" + strconv.FormatInt(chunkOffset, 10) + "-" + strconv.FormatInt(chunkOffset+dataLen-1, 10)
		}
	})
}
