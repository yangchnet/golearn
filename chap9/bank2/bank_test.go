package bank2_test

import (
	"golearn/chap9/bank2"
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank2.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank2.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

}
