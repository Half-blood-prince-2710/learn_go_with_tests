package main

import "testing"


func TestHello(t *testing.T){
	got:= Hello("Manish")
	want:= "Hello, Manish"

	if got != want {
		t.Errorf("got %q want %q",got,want)
	}
}