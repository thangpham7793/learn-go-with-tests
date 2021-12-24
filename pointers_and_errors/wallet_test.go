package pointers_and_errors

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		if got := wallet.Balance(); got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(20)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient fund", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		err := wallet.Withdraw(Bitcoin(40))

		assertError(t, err, InsufficientFundError)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("expect no error but got %s", got)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expect error but got none")
	}

	if got != want {
		t.Fatalf("expect %s but got %s", want, got)
	}
}
