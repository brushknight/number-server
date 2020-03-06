package processor

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"number-server/internal/reporter"
	"number-server/internal/storage"
	"testing"
	"time"
)

func TestProcessor_ProcessChannel(t *testing.T) {

	// numbersInQueue is unbuffered because we need to fire and handle events in exact same order (synchronous)

	storageMock := storage.NewMockStorage()
	loggerMock := logrus.StandardLogger()
	numbersInQueue := make(chan uint64)
	numbersOutQueue := make(chan uint64, 100)
	reportsQueue := make(chan reporter.ReportDTO)
	triggerReportChannel := make(chan time.Time)

	processor := NewProcessor(storageMock, loggerMock)

	go func() {
		numbersInQueue <- 123456
		numbersInQueue <- 123457
		numbersInQueue <- 123458
		numbersInQueue <- 123459
		numbersInQueue <- 123456
		numbersInQueue <- 123457
		numbersInQueue <- 123458
		numbersInQueue <- 123460

		triggerReportChannel <- time.Time{}

		close(triggerReportChannel)
		close(numbersInQueue) // may cause error
	}()

	fmt.Println(&numbersInQueue)

	go processor.ProcessChannel(numbersInQueue, numbersOutQueue, triggerReportChannel, reportsQueue)

	var reports []reporter.ReportDTO

	for report := range reportsQueue {
		reports = append(reports, report)
	}

	resultMap := make(map[uint64]bool)

	for number := range numbersOutQueue {
		resultMap[number] = true
	}

	if len(reports) == 1 {

		report := reports[0]

		if report.AllNumbersTotal != 8 {
			t.Errorf("Report: Expected to have total = %d in report, got %d", 8, report.AllNumbersTotal)
		}

		if report.UniqNumbersFromLastReport != 5 {
			t.Errorf("Report: Expected to have uniqNumbersFromLastReport = %d in report, got %d", 5, report.UniqNumbersFromLastReport)
		}

		if report.UniqNumbersTotal != 5 {
			t.Errorf("Report: Expected to have uniqNumbersTotal = %d in report, got %d", 5, report.UniqNumbersTotal)
		}

		if report.DuplicateNumbersFromLastReport != 3 {
			t.Errorf("Report: Expected to have duplicateNumbersFromLastReport = %d in report, got %d", 3, report.DuplicateNumbersFromLastReport)
		}

	} else {
		t.Errorf("Report: Expected to have %d report, got %d", 1, len(reports))
		t.Errorf("Report: Unexpected issue")
	}

	if len(resultMap) != 5 {
		t.Errorf("Expected to have %d elements in the out queue, got %d", 5, len(resultMap))
	}

	if storageMock.MethodCalledTimes("IsNumberExists") != 8 {
		t.Errorf("Expected to call Storage method IsNumberExists %d times, got %d", 9, storageMock.MethodCalledTimes("IsNumberExists"))
	}

	if storageMock.MethodCalledTimes("AddNumber") != 5 {
		t.Errorf("Expected to call Storage method AddNumber %d times, got %d", 5, storageMock.MethodCalledTimes("AddNumber"))
	}

	assertMethodAddNumberWasCalledWithValueTimes(storageMock, 123456, 1, t)
	assertMethodAddNumberWasCalledWithValueTimes(storageMock, 123457, 1, t)
	assertMethodAddNumberWasCalledWithValueTimes(storageMock, 123458, 1, t)
	assertMethodAddNumberWasCalledWithValueTimes(storageMock, 123459, 1, t)
	assertMethodAddNumberWasCalledWithValueTimes(storageMock, 123460, 1, t)

	assertMethodIsNumberExistsWasCalledWithValueTimes(storageMock, 123456, 2, t)
	assertMethodIsNumberExistsWasCalledWithValueTimes(storageMock, 123457, 2, t)
	assertMethodIsNumberExistsWasCalledWithValueTimes(storageMock, 123458, 2, t)
	assertMethodIsNumberExistsWasCalledWithValueTimes(storageMock, 123459, 1, t)
	assertMethodIsNumberExistsWasCalledWithValueTimes(storageMock, 123460, 1, t)

	assertNumberInTheMap(resultMap, 123456, t)
	assertNumberInTheMap(resultMap, 123457, t)
	assertNumberInTheMap(resultMap, 123458, t)
	assertNumberInTheMap(resultMap, 123459, t)
	assertNumberInTheMap(resultMap, 123460, t)
}

func assertMethodAddNumberWasCalledWithValueTimes(checkingMap *storage.NumberStorageMock, number uint64, times uint64, t *testing.T) {
	if checkingMap.MethodCalledTimesWithValue("AddNumber", number) != times {
		t.Errorf("Storage method AddNumber expected to be called %d times with number %d, got %d", times, number, checkingMap.MethodCalledTimesWithValue("AddNumber", number))
	}
}

func assertMethodIsNumberExistsWasCalledWithValueTimes(checkingMap *storage.NumberStorageMock, number uint64, times uint64, t *testing.T) {
	if checkingMap.MethodCalledTimesWithValue("IsNumberExists", number) != times {
		t.Errorf("Storage method IsNumberExists expected to be called %d times with number %d, got %d", times, number, checkingMap.MethodCalledTimesWithValue("IsNumberExists", number))
	}
}

func assertNumberInTheMap(checkingMap map[uint64]bool, number uint64, t *testing.T) {
	if !isNumberInTheMap(checkingMap, number) {
		t.Errorf("Number %d expected to be in the out queue, but it is not", number)
	}
}

func isNumberInTheMap(checkingMap map[uint64]bool, number uint64) bool {
	if val, ok := checkingMap[number]; ok {
		return val
	}

	return false
}
