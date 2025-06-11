package pointers

import (
	"fmt"
	"math"
	"testing"
)

func TestWalletDeposit(t *testing.T) {
	wallet := Wallet{1.0}
	wallet.Deposit(9.0)
	expected := Bitcoin(10.0)

	assertBalance(t, wallet, expected)
}

func TestWalletWithdraw(t *testing.T) {
	wallet := Wallet{100.0}
	msg, err := wallet.Withdraw(Bitcoin(90.0))
	expected := Bitcoin(10.0)

	assertBalance(t, wallet, expected)
	assertError(t, err, msg)
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	got := wallet.Balance()

	const epsilon = 0.0001
	if math.Abs(float64(got-expected)) > epsilon {
		t.Errorf("%#v got %.2f expected %.2f", wallet, got, expected)
	}
}

func assertError(t testing.TB, err error, msg string) {
	t.Helper()
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(msg)
	}
}
