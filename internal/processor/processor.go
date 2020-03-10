package processor

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"number-server/internal/reporter"
	storage2 "number-server/internal/storage"
	"number-server/internal/terminator"
	"time"
)

type Processor struct {
	storage                                 storage2.NumberStorageInterface
	logger                                  logrus.Ext1FieldLogger
	totalNumbersProcesses                   uint64
	uniqNumbersProcessedFromLastReport      uint64
	duplicateNumbersProcessesFromLastReport uint64
}

func (p *Processor) ProcessChannel(numbersQueue chan uint64, dumperQueue chan uint64, reportTriggerChannel <-chan time.Time, reportsQueue chan reporter.ReportDTO, triggerTerminationChannel chan string) {
	defer close(dumperQueue)
	defer close(reportsQueue)

	for {
		select {
		case number, ok := <-numbersQueue:
			p.logger.Trace(fmt.Sprintf("Processor: Get number %d", number))

			if !ok {
				p.logger.Debug("[x] Terminating messages processor")
				return
			}

			if p.storage.IsNumberExists(number) {
				p.duplicateNumbersProcessesFromLastReport++
			} else {
				status, err := p.storage.AddNumber(number)

				if err != nil {
					p.logger.Error(fmt.Sprintf("%e", err))
					p.terminate(triggerTerminationChannel, fmt.Sprintf("%e", err))
				}
				if !status {
					p.logger.Error(fmt.Sprintf("number was not inserted %d", number))
					p.terminate(triggerTerminationChannel, fmt.Sprintf("number was not inserted %d", number))
				}

				dumperQueue <- number
				p.logger.Trace(fmt.Sprintf("Processor: Sent to dumer queue %d", number))
				p.uniqNumbersProcessedFromLastReport++
			}

			p.totalNumbersProcesses++
		case _, ok := <-reportTriggerChannel:
			if !ok {
				p.logger.Debug("[x] Terminating messages processor")
				return
			}
			p.doReport(reportsQueue)
		}
	}
}

func (p *Processor) doReport(reportsQueue chan reporter.ReportDTO) {
	reportsQueue <- reporter.CreateReportDTO(p.storage.GetLength(), p.uniqNumbersProcessedFromLastReport, p.duplicateNumbersProcessesFromLastReport, p.totalNumbersProcesses, 0)
	p.uniqNumbersProcessedFromLastReport = 0
	p.duplicateNumbersProcessesFromLastReport = 0
}

func (p *Processor) terminate(triggerTerminationChannel chan string, reason string) {
	terminator.Terminate(triggerTerminationChannel, "Processor", reason)
}

func NewProcessor(storage storage2.NumberStorageInterface, logger logrus.Ext1FieldLogger) *Processor {
	return &Processor{storage: storage, logger: logger}
}
