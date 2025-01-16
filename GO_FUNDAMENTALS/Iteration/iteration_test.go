package iteration

import "testing"


func TestIteration(t *testing.T){
	t.Run("repeating char 5 times", func(t *testing.T) {
		got:= Repeat("a",5)
		want:= "aaaaa"

		assertError(t,got,want)
	})
	t.Run("repeat 0 times", func(t *testing.T) {
		got:= Repeat("a",0)
		want:=""
		assertError(t,got ,want)
	})
	t.Run("repeat 1 times", func(t *testing.T) {
		got:= Repeat("a",1)
		want:="a"
		assertError(t,got ,want)
	})
	t.Run("repeat 10 times", func(t *testing.T) {
		got:= Repeat("a",10)
		want:="aaaaaaaaaa"
		assertError(t,got ,want)
	})
	t.Run("count is negative", func(t *testing.T) {
		got:= Repeat("a",-2)
		want:= "Repeat count should be greater than equal to 0"
		assertError(t,got ,want)
	})
}


func assertError(t testing.TB, got string, want string){
	t.Helper()
	if got!=want {
		t.Errorf("got %q want %q ",got,want)
	}

}

func BenchmarkRepeat(b *testing.B) {
	b.Run("benchmark repeat 1000 times", func(b *testing.B) {
		for i:=0;i<b.N;i++ {
			Repeat("a", 1000)
		}
	})
	b.Run("benchmark repeat 5000 times", func(b *testing.B) {
		for i:=0;i<b.N;i++ {
			Repeat("a", 5000)
		}
	})
	b.Run("benchmark repeat 10000 times", func(b *testing.B) {
		for i:=0;i<b.N;i++ {
			Repeat("a", 10000)
		}
	})

}