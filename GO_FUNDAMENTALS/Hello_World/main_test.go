package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying helllo to people", func(t *testing.T) {
		got := Hello("Manish","")
		want := "Hello, Manish"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func (t *testing.T)  {
		got := Hello("","")
		want:= "Hello, world"
		assertCorrectMessage(t,got,want)
		
	})
	 t.Run("in Spanish",func(t *testing.T) {
		got := Hello("Manish","Spanish")
		want := "Hola, Manish"

		assertCorrectMessage(t,got, want)
	 })
	 t.Run("in French", func(t *testing.T) {
		got:= Hello("manish","French")
		want  := "Bonjour, manish"

		assertCorrectMessage(t,got,want)
	 })
}


func assertCorrectMessage(t testing.TB, got string, want string){
	t.Helper() // will tell this func is helper and when we got error ot will show error where it got error not here at declaration
	if got != want {
			t.Errorf("got %q want %q", got, want)
		}
}