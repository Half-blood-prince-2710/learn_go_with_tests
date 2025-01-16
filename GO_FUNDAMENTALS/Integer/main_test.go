package integer

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdd(t *testing.T) {
	t.Run("sum of two number",func(t *testing.T) {
		got := Add(7,-6)
		want := 1

		assertError(t,got,want)
	})
}


func assertError(t testing.TB, got int, want int){
	t.Helper()
	if got!=want {
		t.Errorf("got %d want %d ",got,want)
	}

}