package money

func NewDollar(amount int) Dollar {
	return Dollar{}
}

type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) {

}
