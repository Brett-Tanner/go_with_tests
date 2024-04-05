package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("depositing funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	},
	)

	t.Run("withdrawing funds", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(5)
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(100)

		assertError(t, err, errInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got unexpected error")
	}
}
