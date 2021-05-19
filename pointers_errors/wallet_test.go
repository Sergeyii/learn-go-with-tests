package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet.Balance(), Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		//Снять со счёта
		gotErr := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet.Balance(), Bitcoin(10))
		assertNoError(t, gotErr)
	})

	t.Run("Withdraw insufficient funds(проверка на списание средств больше хранящегося)", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		gotErr := wallet.Withdraw(50)
		assertBalance(t, wallet.Balance(), Bitcoin(10))
		assertError(t, gotErr, "Нельзя снимать денег больше чем имеется!")
	})
}

func assertBalance(t testing.TB, got Bitcoin, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got=%s, want=%s", got, want)
	}
}

func assertNoError(t testing.TB, gotErr error) {
	t.Helper()
	if gotErr != nil {
		t.Fatalf("got an error but didn't want one")
	}
}

func assertError(t testing.TB, gotErr error, wantErr string) {
	t.Helper()
	if gotErr == nil {
		t.Fatalf("wanted an error but didn't get one")
	}

	if gotErr.Error() != wantErr {
		t.Errorf("got=%q, want=%q", gotErr, wantErr)
	}
}
