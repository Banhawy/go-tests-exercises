package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times by default", func(t *testing.T) {
		got := Repeat("5")
		want := "55555"

		assertCorrectMessage(t, got, want)
	})

	t.Run("repeat x times given x as 2nd parameter", func(t *testing.T) {
		got := Repeat("5", 2)
		want := "55"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("5")
	}
}

func ExampleRepeat() {
	repeatX := Repeat("x", 3)
	fmt.Println(repeatX)
	// Output: xxx
}
