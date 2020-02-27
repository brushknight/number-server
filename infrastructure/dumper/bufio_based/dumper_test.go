package bufio_based

import (
	"bufio"
	"number-server/infrastructure/logger/mock"
	"testing"
)

type MockWriter struct {
	InternalSlice []byte
}

func (w *MockWriter) Write(p []byte) (n int, err error) {
	w.InternalSlice = append(w.InternalSlice, p...)

	return len(p), nil
}

func TestDumper_StartListening_WithLeadingZeros(t *testing.T) {

	dumperInQueue := make(chan uint64, 1000)
	mockWriter := new(MockWriter)
	loggerMock := mock.NewMockLogger()
	writer := bufio.NewWriterSize(mockWriter, 4096)
	var numbersNumber uint64 = 2

	dumper := NewDumper(true, loggerMock)

	go func() {
		defer close(dumperInQueue)
		var i uint64 = 1
		for i = 1; i < numbersNumber; i++ {
			dumperInQueue <- i
		}
	}()

	dumper.ProcessChannel(dumperInQueue, writer)

	got := string(mockWriter.InternalSlice)

	want := "000000001\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDumper_StartListening_WithoutLeadingZeros(t *testing.T) {

	dumperInQueue := make(chan uint64, 1000)
	mockWriter := new(MockWriter)
	loggerMock := mock.NewMockLogger()
	writer := bufio.NewWriterSize(mockWriter, 4096)
	var numbersNumber uint64 = 2

	dumper := NewDumper(false, loggerMock)

	go func() {
		defer close(dumperInQueue)
		var i uint64 = 1
		for i = 1; i < numbersNumber; i++ {
			dumperInQueue <- i
		}
	}()

	dumper.ProcessChannel(dumperInQueue, writer)
	writer.Flush()

	got := string(mockWriter.InternalSlice)

	want := "1\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDumper_StartListening_MultipleWithLeadingZeros(t *testing.T) {

	dumperInQueue := make(chan uint64, 1000)
	mockWriter := new(MockWriter)
	loggerMock := mock.NewMockLogger()
	writer := bufio.NewWriterSize(mockWriter, 4096)
	var numbersNumber uint64 = 10

	dumper := NewDumper(true, loggerMock)

	go func() {
		defer close(dumperInQueue)
		var i uint64 = 1
		for i = 1; i < numbersNumber; i++ {
			dumperInQueue <- i
		}
	}()

	dumper.ProcessChannel(dumperInQueue, writer)
	writer.Flush()

	got := string(mockWriter.InternalSlice)

	want := "000000001\n000000002\n000000003\n000000004\n000000005\n000000006\n000000007\n000000008\n000000009\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDumper_StartListening_MultipleWithoutLeadingZeros(t *testing.T) {

	dumperInQueue := make(chan uint64, 1000)
	mockWriter := new(MockWriter)
	loggerMock := mock.NewMockLogger()
	writer := bufio.NewWriterSize(mockWriter, 4096)
	var numbersNumber uint64 = 10

	dumper := NewDumper(false, loggerMock)

	go func() {
		defer close(dumperInQueue)
		var i uint64 = 1
		for i = 1; i < numbersNumber; i++ {
			dumperInQueue <- i
		}
	}()

	dumper.ProcessChannel(dumperInQueue, writer)
	writer.Flush()

	got := string(mockWriter.InternalSlice)
	want := "1\n2\n3\n4\n5\n6\n7\n8\n9\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDumper_StartListening_WithLeadingZerosBigNumber(t *testing.T) {

	dumperInQueue := make(chan uint64, 1000)
	mockWriter := new(MockWriter)
	loggerMock := mock.NewMockLogger()
	writer := bufio.NewWriterSize(mockWriter, 4096)

	dumper := NewDumper(true, loggerMock)

	go func() {
		defer close(dumperInQueue)
		dumperInQueue <- 1423123
	}()

	dumper.ProcessChannel(dumperInQueue, writer)
	writer.Flush()

	got := string(mockWriter.InternalSlice)
	want := "001423123\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
