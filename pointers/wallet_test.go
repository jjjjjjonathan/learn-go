package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("deposits", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdrawals", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdrawing insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
