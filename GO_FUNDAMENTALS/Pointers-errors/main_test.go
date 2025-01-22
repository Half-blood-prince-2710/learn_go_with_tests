package main

import "testing"


func TestWallet(t *testing.T){
	wallet := Wallet{}

	wallet.Deposit(20)

	got:=wallet.Balance()
	want:= 20

	if got!=want {
		t.Errorf("got %d want %d",got,want)
	}
}