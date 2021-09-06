package money

func NewDollar(amount int) Dollar {
	return Dollar{amount: amount}
}

type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) {
	d.amount = d.amount * 2
}
