package application

type MessageHandlerInterface interface {
	ValidateAndParse(message string) (uint64, error)
	Handle(number uint64)
	Terminate()
}
