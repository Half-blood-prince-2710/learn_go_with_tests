package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB,got error, want string){
		t.Helper()
		if got == nil {
			t.Errorf("wanting a error but didnt get one")
		}
		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("testing insuffieciend friends",func(t *testing.T) {
		initialFunds := Bitcoin(20)
		wallet := Wallet{initialFunds}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t,err,"insufficient funds")
		assertBalance(t,wallet,initialFunds)
		
	})
}