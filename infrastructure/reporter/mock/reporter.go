package mock

import (
	"time"
)

type Reporter struct {
	UniqNumbersProcesses      uint64
	DuplicateNumbersProcesses uint64
	TotalUniqNumbersProcesses uint64
	TotalNumbersProcesses     uint64
	CalledTimes               uint
	frequencyMs               time.Duration
}

func (r *Reporter) MakeAReport(uniqNumbersProcesses uint64, duplicateNumbersProcesses uint64, totalUniqNumbersProcesses uint64, totalNumbersProcesses uint64) {

	r.UniqNumbersProcesses = uniqNumbersProcesses
	r.DuplicateNumbersProcesses = duplicateNumbersProcesses
	r.TotalUniqNumbersProcesses = totalUniqNumbersProcesses
	r.TotalNumbersProcesses = totalNumbersProcesses
	r.CalledTimes++
}

func (r *Reporter) GetFrequencyMs() time.Duration {
	return r.frequencyMs
}

func NewMockReporter(frequencyMs time.Duration) *Reporter {
	return &Reporter{frequencyMs: frequencyMs}
}
