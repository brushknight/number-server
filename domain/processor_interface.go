package domain

import "time"

type ProcessorInterface interface {
	ProcessChannel(numbersQueue chan uint64, dumperQueue chan uint64, reportTriggerChannel chan time.Time, reportsQueue chan ReportDTO)
}
