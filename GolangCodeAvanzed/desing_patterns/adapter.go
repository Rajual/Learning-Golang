package main

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {

}
