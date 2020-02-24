package stdout

import (
	"fmt"
	"number-server/infrastructure/logger/mock"
	"testing"
	"time"
)

func TestReporter_MakeAReport(t *testing.T) {

	logger := mock.NewLogger()

	reporter := NewReporter(10*time.Millisecond, logger, "prd")

	var uniqNumbersProcesses uint64 = 50
	var duplicateNumbersProcesses uint64 = 2
	var totalUniqNumbersProcesses uint64 = 567231

	reporter.MakeAReport(uniqNumbersProcesses, duplicateNumbersProcesses, totalUniqNumbersProcesses, 1000000)

	expectedMessage := fmt.Sprintf("[INFO] Received %d unique numbers, %d duplicates. Unique total: %d", uniqNumbersProcesses, duplicateNumbersProcesses, totalUniqNumbersProcesses)

	if logger.LastMessage != expectedMessage {
		t.Errorf("Expected message : %s, got: %s", expectedMessage, logger.LastMessage)
	}

}
