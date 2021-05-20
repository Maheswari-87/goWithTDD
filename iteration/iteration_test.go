package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q, want %q ", got, want)
	}
	//fmt.Printf("%q\n", strings.Split(" aaaaa ", ""))

}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
func ExampleRepeat() {
	got := Repeat("a", 6)
	fmt.Println(got)
	//output:aaaaaa

}
