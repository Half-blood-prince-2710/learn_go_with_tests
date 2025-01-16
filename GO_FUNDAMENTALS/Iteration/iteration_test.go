package iteration

import "testing"


func TestIteration(t *testing.T){
	t.Run("repeating char 5 times", func(t *testing.T) {
		got:= Repeat("a")
		want:= "aaaaa"

		assertError(t,got,want)
	})
}


func assertError(t testing.TB, got string, want string){
	t.Helper()
	if got!=want {
		t.Errorf("got %q want %q ",got,want)
	}

}