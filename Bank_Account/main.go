package main

import "fmt"

//Initialize Account Type with Functions

type Account interface {
	add(value int) int
	decrease(value int) int
	get() int
}

//Initialize Accountbalance with balance as a varialbe

type AccountBalance struct {
	balance int
}

//The Account Function add with a pointer to the Accountbalance to reference the Value

func (b *AccountBalance) add(value int) int {
	b.balance += value
	return b.balance
}

// The Account Function decrease with a pointer to the Accountbalance to reference the Value
func (b *AccountBalance) decrease(value int) int {
	b.balance -= value
	return b.balance
}

// The Account Function add with a pointer to the Accountbalance to reference the Value
func (b AccountBalance) get() int {
	return b.balance
}

func main() {
	var balance Account = &AccountBalance{balance: 0}
	fmt.Println(balance.get())
}
