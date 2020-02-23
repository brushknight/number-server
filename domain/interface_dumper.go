package domain

type DumperInterface interface {
	ProcessChannel(c chan uint64)
}
