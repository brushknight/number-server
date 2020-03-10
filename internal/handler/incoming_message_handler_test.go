package handler

import (
	"fmt"
	"testing"
)

func TestIncomingMessageHandler_ValidateAndParse_Positive(t *testing.T) {

	numbersQueue := make(chan uint64)

	handler := NewMessageHandler(numbersQueue)

	var expectedNumber uint64 = 123456
	message := fmt.Sprintf("%09d", expectedNumber)

	number, err := handler.ValidateAndParse(message)

	if err != nil {
		t.Errorf("Unexpected error occurred: %e.", err)
	}

	if number != expectedNumber {
		t.Errorf("Number from parser doesn't match expected. Expected %d, got %d.", expectedNumber, number)
	}
}

func TestIncomingMessageHandler_ValidateAndParse_Negative_Length(t *testing.T) {

	numbersQueue := make(chan uint64)

	handler := NewMessageHandler(numbersQueue)

	var expectedNumber uint64 = 123456
	message := fmt.Sprintf("d%06d", expectedNumber)

	_, err := handler.ValidateAndParse(message)

	if err.Error() != fmt.Sprintf("message should be 9 symbols length, got %d message: %s", len(message), message) {
		t.Errorf("Unexpected error occurred: %e.", err)
	}
}

func TestIncomingMessageHandler_ValidateAndParse_Negative_NaN(t *testing.T) {

	numbersQueue := make(chan uint64)

	handler := NewMessageHandler(numbersQueue)

	var expectedNumber uint64 = 123456
	message := fmt.Sprintf("O%08d", expectedNumber)

	_, err := handler.ValidateAndParse(message)

	if err == nil {
		t.Errorf("Error on number parsing should be thrown")
	}
}
