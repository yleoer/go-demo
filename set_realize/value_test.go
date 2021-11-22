package set_realize

import "testing"

const num = int(1 << 24)

func BenchmarkSetWithBoolValue(b *testing.B) {
	set := make(map[int]bool)
	for i := 0; i < b.N; i++ {
		set[i] = true
	}
}

func BenchmarkSetWithInterfaceValue(b *testing.B) {
	set := make(map[int]interface{})
	for i := 0; i < b.N; i++ {
		set[i] = struct{}{}
	}
}

func BenchmarkSetWithIntValue(b *testing.B) {
	set := make(map[int]int)
	for i := 0; i < b.N; i++ {
		set[i] = 0
	}
}

func BenchmarkSetWithStructValue(b *testing.B) {
	set := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		set[i] = struct{}{}
	}
}
