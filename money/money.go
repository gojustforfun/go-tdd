package money

func NewDollar(amount int) dollar {
	return dollar{amount: amount}
}

type dollar struct {
	amount int
}

func (d *dollar) Times(multiplier int) dollar {
	return NewDollar(d.amount * multiplier)
}
