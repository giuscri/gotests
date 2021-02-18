package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit bitcoins", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("withdraw bitcoins", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)

		if err != nil {
			t.Errorf("failed to withdraw")
		}

		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("withdraw more bitcoins than in the wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(100)

		if err == nil {
			t.Errorf("withdrawing more btc than in the wallet should return an error")
		}

		got := wallet.Balance()
		want := Bitcoin(20)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
