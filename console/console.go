package console

type IConsole interface {
	ReadLine() string
	ReadStringArray() []string
	ReadIntArray() ([] int, error)
}

type Console struct {
	IConsole
}

func New() *Console  {
	return &Console{}
}