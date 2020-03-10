package handler

type MessageHandlerInterface interface {
	ValidateAndParse(message string) (uint64, error)
	Handle(number uint64)
}
