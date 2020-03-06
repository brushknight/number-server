package stdout

import (
	"bytes"
	"fmt"
	"number-server/domain"
	"testing"
)

func TestReporter_ProcessChannel(t *testing.T) {

	env := "prd"

	reportsQueue := make(chan domain.ReportDTO, 10)

	var uniqNumbersProcesses uint64 = 50
	var duplicateNumbersProcesses uint64 = 2
	var totalUniqNumbersProcesses uint64 = 567231

	reportsQueue <- domain.CreateReportDTO(totalUniqNumbersProcesses, uniqNumbersProcesses, duplicateNumbersProcesses, totalUniqNumbersProcesses+duplicateNumbersProcesses, totalUniqNumbersProcesses+duplicateNumbersProcesses/10)

	close(reportsQueue)

	buffer := &bytes.Buffer{}

	ProcessReportsChannel(reportsQueue, buffer, env)

	expectedMessage := fmt.Sprintf("Received %d unique numbers, %d duplicates. Unique total: %d\n", uniqNumbersProcesses, duplicateNumbersProcesses, totalUniqNumbersProcesses)

	got := buffer.String()

	if got != expectedMessage {
		t.Errorf("Expected message : %s, got: %s", expectedMessage, got)
	}

}
