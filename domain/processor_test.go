package domain

import (
	"bufio"
	mock3 "number-server/infrastructure/dumper/mock"
	mock2 "number-server/infrastructure/logger/mock"
	"number-server/infrastructure/reporter/mock"
	"number-server/infrastructure/storage/memory"
	"testing"
	"time"
)

type MockWriter struct {
	InternalSlice []byte
}

func (w *MockWriter) Write(p []byte) (n int, err error) {
	w.InternalSlice = append(w.InternalSlice, p...)

	return len(p), nil
}

func TestProcessor_DoReport(t *testing.T) {

	reporterMock := mock.NewMockReporter(10 * time.Second)
	loggerMock := mock2.NewMockLogger()
	storage := memory.NewNumberStorage()
	numbersQueue := make(chan uint64, 100)
	dumperQueue := make(chan uint64, 100)
	triggerTerminationChannel := make(chan string)

	processor := NewProcessor(numbersQueue, storage, reporterMock, triggerTerminationChannel, loggerMock)

	go func() {
		numbersQueue <- 123
		close(numbersQueue)
	}()

	processor.StartProcessing(dumperQueue)

	if processor.uniqNumbersProcessedFromLastReport != 1 {
		t.Errorf("Before report amount of uniq numbers processes from last report should be 1, got %d", processor.uniqNumbersProcessedFromLastReport)
	}

	if processor.duplicateNumbersProcessesFromLastReport != 0 {
		t.Errorf("Before report amount of duplicate numbers processes from last report should be 0, got %d", processor.duplicateNumbersProcessesFromLastReport)
	}

	processor.DoReport()

	if reporterMock.CalledTimes != 1 {
		t.Errorf("After report reporter mast be called 1 times, got %d", reporterMock.CalledTimes)
	}

	if processor.uniqNumbersProcessedFromLastReport != 0 {
		t.Errorf("After report amount of uniq numbers processes from last report should be 0, got %d", processor.uniqNumbersProcessedFromLastReport)
	}

	if processor.duplicateNumbersProcessesFromLastReport != 0 {
		t.Errorf("After report amount of duplicate numbers processes from last report should be 0, got %d", processor.duplicateNumbersProcessesFromLastReport)
	}
}

func TestProcessor_StartProcessing(t *testing.T) {
	reporterMock := mock.NewMockReporter(10 * time.Second)
	loggerMock := mock2.NewMockLogger()
	dumperMock := mock3.NewMockDumper()
	storage := memory.NewNumberStorage()
	numbersQueue := make(chan uint64, 100)
	dumperQueue := make(chan uint64, 100)
	triggerTerminationChannel := make(chan string)
	mockWriter := new(MockWriter)
	writer := bufio.NewWriterSize(mockWriter, 4096)


	go func() {
		numbersQueue <- 123456789
		numbersQueue <- 223456789
		numbersQueue <- 323456789
		numbersQueue <- 423456789
		numbersQueue <- 123456789
		numbersQueue <- 223456789
		numbersQueue <- 323456789
		numbersQueue <- 423456789
		numbersQueue <- 523456789
		close(numbersQueue)
	}()

	processor := NewProcessor(numbersQueue, storage, reporterMock, triggerTerminationChannel, loggerMock)
	processor.StartProcessing(dumperQueue)
	dumperMock.ProcessChannel(dumperQueue, writer)

	numberShouldBeInTheStorage(storage,123456789, t)
	numberShouldBeInTheStorage(storage,223456789, t)
	numberShouldBeInTheStorage(storage,323456789, t)
	numberShouldBeInTheStorage(storage,423456789, t)
	numberShouldBeInTheStorage(storage,523456789, t)

	if storage.GetLength() != 5 {
		t.Errorf("Storage length should be 5, got %d", storage.GetLength())
	}

	if len(dumperMock.NumbersDumped) != 5 {
		t.Errorf("Dumper should dump 5 numbers, got %d", len(dumperMock.NumbersDumped))
	}

	if processor.uniqNumbersProcessedFromLastReport != 5 {
		t.Errorf("Before report amount of uniq numbers processes from last report should be 5, got %d", processor.uniqNumbersProcessedFromLastReport)
	}

	if processor.duplicateNumbersProcessesFromLastReport != 4 {
		t.Errorf("Before report amount of duplicate numbers processes from last report should be 4, got %d", processor.duplicateNumbersProcessesFromLastReport)
	}
}

func numberShouldBeInTheStorage(storage NumberStorageInterface, number uint64, t *testing.T) {
	if !storage.IsNumberExists(number) {
		t.Errorf("Number %d should be in the storage and it is not", number)
	}
}
