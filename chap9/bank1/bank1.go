package bank1

/**
@Description: 不要使用共享数据来通信，使用通信来共享数据。
一个提供对一个指定的变量通过channel来请求的goroutine叫做这个变量的 monitor goroutine
@Date: 2/17/2021
@Author: lichang
*/

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

// balance Monitor
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {

	go teller()
}
