package cmd

type TerminationInterface interface {
	Terminate(by string, reason string)
}
