package domain

import (
	"fmt"
	"time"
)

type Processor struct {
	storage                                 NumberStorageInterface
	logger                                  LoggerInterface
	triggerTerminationChannel               chan string
	totalNumbersProcesses                   uint64
	uniqNumbersProcessedFromLastReport      uint64
	duplicateNumbersProcessesFromLastReport uint64
}

func (p *Processor) ProcessChannel(numbersQueue chan uint64, dumperQueue chan uint64, reportTriggerChannel <-chan time.Time, reportsQueue chan ReportDTO) {
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
					p.terminate()
				}
				if !status {
					p.logger.Error(fmt.Sprintf("number was not inserted %d", number))
					p.terminate()
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

func (p *Processor) doReport(reportsQueue chan ReportDTO) {
	reportsQueue <- CreateReportDTO(p.storage.GetLength(), p.uniqNumbersProcessedFromLastReport, p.duplicateNumbersProcessesFromLastReport, p.totalNumbersProcesses)
	p.uniqNumbersProcessedFromLastReport = 0
	p.duplicateNumbersProcessesFromLastReport = 0
}

func (p *Processor) terminate() {
	p.logger.Debug(fmt.Sprintf("[x] Termination inited"))
	p.triggerTerminationChannel <- "number processor"
}

func NewProcessor(storage NumberStorageInterface, triggerTerminationChannel chan string, logger LoggerInterface) *Processor {
	return &Processor{storage: storage, triggerTerminationChannel: triggerTerminationChannel, logger: logger}
}
