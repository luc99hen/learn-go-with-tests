// Golang's orginal test for strings.Compare
// https://cs.opensource.google/go/go/+/refs/tags/go1.17.1:src/strings/compare_test.go

package strings

import (
	"strings"
	"testing"
)

func assertCorrectMessage(compared int, expected int, t testing.TB) {
	t.Helper()
	if compared != expected {
		t.Errorf("Error: expected: %d, got %d", expected, compared)
	}
}

func TestCompare(t *testing.T) {

	t.Run("less than", func(t *testing.T) {
		got := strings.Compare("aaa", "bbb")
		want := -1
		assertCorrectMessage(want, got, t)
	})

	t.Run("equal", func(t *testing.T) {
		got := strings.Compare("aaa", "aaa")
		want := 1
		assertCorrectMessage(want, got, t)
	})

}
