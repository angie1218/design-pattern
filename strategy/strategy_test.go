package strategy

func ExamplePayByCash() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
}

func ExamplePayByBank() {
	payment := NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
}
