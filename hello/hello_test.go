package hello_test

import (
	"testing"

	"github.com/yleoer/go-demo/hello"
)

func TestGreet(t *testing.T) {
	result := hello.Greet()
	const want = "Hello GitHub Actions"
	if result != want {
		t.Errorf("Greet() got %s, expected %s", result, want)
	}
}
