package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	t.Run("Deposit...", func(t *testing.T) {
		wallet.Deposit(12.0)

		got := wallet.Balance()

		want := Bitcoin(12.0)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})

	t.Run("Withdraw...", func(t *testing.T) {
		err := wallet.Withdraw(Bitcoin(50.0))
		got := wallet.Balance()
		want := Bitcoin(2)

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}

		if err == nil {
			t.Error("Wanted an error but didnt get one")
		}
	})
}
