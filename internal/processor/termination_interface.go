package processor

type TerminationInterface interface {
	Terminate(by string, reason string)
}
