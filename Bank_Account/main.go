package main

import "fmt"

type Account interface {
	add(value int) int
	decrease(value int) int
	get() int
}

type AccountBalance struct {
	balance int
}

func (b *AccountBalance) add(value int) int {
	b.balance += value
	return b.balance
}

func (b *AccountBalance) decrease(value int) int {
	b.balance -= value
	return b.balance
}

func (b *AccountBalance) get() int {
	return b.balance
}

func main() {
	var balance Account = &AccountBalance{balance: 0}
	fmt.Println(balance.get())
}
