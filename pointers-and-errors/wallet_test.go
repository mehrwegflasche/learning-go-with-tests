package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(100)
	got := wallet.Balance()
	want := 100
	if got != want {
		t.Errorf("got %d , wanted %d", got, want)
	}
}
