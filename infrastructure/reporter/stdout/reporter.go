package stdout

import (
	"fmt"
	"number-server/infrastructure/logger"
	"time"
)

type Reporter struct {
	frequencyMs time.Duration
	logger      logger.LoggerInterface
	env         string
}

func (r *Reporter) MakeAReport(uniqNumbersProcesses uint64, duplicateNumbersProcesses uint64, totalUniqNumbersProcesses uint64, totalNumbersProcesses uint64) {

	if r.env == "prd" {
		r.logger.Info(fmt.Sprintf("Received %d unique numbers, %d duplicates. Unique total: %d", uniqNumbersProcesses, duplicateNumbersProcesses, totalUniqNumbersProcesses))
	} else {
		rps := (duplicateNumbersProcesses + uniqNumbersProcesses) / uint64(r.frequencyMs/1000)
		r.logger.Info(fmt.Sprintf("Received %d unique numbers, %d duplicates. Total %d Uniq: %d. Rps %d", uniqNumbersProcesses, duplicateNumbersProcesses, totalNumbersProcesses, totalUniqNumbersProcesses, rps))
	}
}

func (r *Reporter) GetFrequencyMs() time.Duration {
	return r.frequencyMs
}

func NewReporter(frequencyMs time.Duration, logger logger.LoggerInterface, env string) *Reporter {
	return &Reporter{frequencyMs: frequencyMs, logger: logger, env: env}
}
