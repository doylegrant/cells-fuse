package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFmtSprintf(b *testing.B) {
	path := "/some/long/path/to/a/file.txt"
	chunkIndex := int64(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s\x00%d", path, chunkIndex)
	}
}

func BenchmarkStrconv(b *testing.B) {
	path := "/some/long/path/to/a/file.txt"
	chunkIndex := int64(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = path + "\x00" + strconv.FormatInt(chunkIndex, 10)
	}
}

func BenchmarkByteRangeFmt(b *testing.B) {
	chunkOffset := int64(1048576)
	length := 4194304
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("bytes=%d-%d", chunkOffset, chunkOffset+int64(length)-1)
	}
}

func BenchmarkByteRangeStrconv(b *testing.B) {
	chunkOffset := int64(1048576)
	length := 4194304
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = "bytes=" + strconv.FormatInt(chunkOffset, 10) + "-" + strconv.FormatInt(chunkOffset+int64(length)-1, 10)
	}
}
