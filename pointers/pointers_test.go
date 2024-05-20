package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, w *Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
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

	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t, &w, Bitcoin(10))
	})

	t.Run("Withdraw Sufficient Balance", func(t *testing.T) {
		startingBal := Bitcoin(10)
		w := Wallet{startingBal}
		err := w.Withdraw(Bitcoin(5))
		assertBalance(t, &w, Bitcoin(5))
		assertNoError(t, err)
	})

	t.Run("Withdraw Insuffient Balance", func(t *testing.T) {
		startingBal := Bitcoin(10)
		w := Wallet{startingBal}
		err := w.Withdraw(Bitcoin(15))
		assertError(t, err, ErrInsuffientBalance)
		assertBalance(t, &w, Bitcoin(10))
	})
}
