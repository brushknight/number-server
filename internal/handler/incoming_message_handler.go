package handler

import (
	"errors"
	"fmt"
	"strconv"
)

type IncomingMessageHandler struct {
	numbersQueue chan uint64
}

func (h *IncomingMessageHandler) Handle(number uint64) {
	h.numbersQueue <- number
}

func (h *IncomingMessageHandler) ValidateAndParse(message string) (uint64, error) {

	messageLength := len(message)

	if messageLength != 9 {
		return 0, errors.New(fmt.Sprintf("message should be 9 symbols length, got %d message: %s", messageLength, message))
	}

	number, err := strconv.ParseUint(message, 10, 64)

	if err != nil {
		return 0, err
	}

	return number, nil
}

func NewMessageHandler(numbersQueue chan uint64) *IncomingMessageHandler {
	return &IncomingMessageHandler{numbersQueue: numbersQueue}
}
