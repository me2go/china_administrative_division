package china_administrative_division

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/kr/pretty"
)

func TestSearch(t *testing.T) {
	segs, ok := Search("650109")
	if ok {
		pretty.Logln(segs)
	}
}

func BenchmarkSearch(b *testing.B) {
	b.StopTimer()
	str := "0123456789"
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j := 0; j < 6; j++ {
			k := rand.Intn(10)
			buf.WriteByte(str[k])
		}
		code := buf.String()
		b.StartTimer()
		Search(code)
		b.StopTimer()
	}
}
