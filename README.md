

## Instructions

Simulate a bank account supporting opening/closing, withdrawals, and deposits
of money. Watch out for concurrent transactions!

A bank account can be accessed in multiple ways. Clients can make
deposits and withdrawals using the internet, mobile phones, etc. Shops
can charge against the account.

Create an account that can be accessed from multiple threads/processes
(terminology depends on your programming language).

It should be possible to close an account; operations against a closed
account must fail.



## Implementation Notes

If Open is given a negative initial deposit, it must return nil.
Deposit must handle a negative amount as a withdrawal. Withdrawals must
not succeed if they result in a negative balance.
If any Account method is called on an closed account, it must not modify
the account and must return ok = false.

The tests will execute some operations concurrently. You should strive
to ensure that operations on the Account leave it in a consistent state.
For example: multiple goroutines may be depositing and withdrawing money
simultaneously, two withdrawals occurring concurrently should not be able
to bring the balance into the negative.
