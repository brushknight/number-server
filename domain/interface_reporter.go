package domain

import "time"

type ReporterInterface interface {
	GetFrequencyMs() time.Duration
	MakeAReport(uniqNumbersProcesses uint64, duplicateNumbersProcesses uint64, totalUniqNumbersProcesses uint64, totalNumbersProcesses uint64)
}
