package account

import "sync"

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	ad := Account{
		saldo: initialDeposit,
		open:  true,
	}

	return &ad
}

func (account *Account) Balance() (int64, bool) {
	return account.saldo, account.open
}
func (account *Account) Close() (int64, bool) {

	account.Lock()
	defer account.Unlock()

	r := account.saldo
	o := account.open

	account.open = false
	account.saldo = 0

	return r, o
}
func (account *Account) Deposit(d int64) (int64, bool) {

	account.Lock()
	defer account.Unlock()

	s := account.saldo
	if d < 0 {
		if s+d < 0 {
			return 0, false
		}
	}
	account.saldo += d
	return account.saldo, account.open
}

type Account struct {
	sync.Mutex
	saldo int64
	open  bool
}
