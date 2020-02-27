package domain

import (
	"fmt"
	"time"
)

type Processor struct {
	numbersQueue                            chan uint64
	storage                                 NumberStorageInterface
	reporter                                ReporterInterface
	dumper                                  DumperInterface
	triggerTerminationChannel               chan string
	logger                                  LoggerInterface
	totalNumbersProcesses                   uint64
	uniqNumbersProcessedFromLastReport      uint64
	duplicateNumbersProcessesFromLastReport uint64
}

func (p *Processor) StartProcessing(dumperQueue chan uint64) {

	ticker := time.NewTicker(p.reporter.GetFrequencyMs() * time.Millisecond)
	defer ticker.Stop()

	defer close(dumperQueue)

	for {
		select {
		case number, ok := <-p.numbersQueue:

			if !ok {
				p.logger.Debug("[x] Terminating messages processor")
				return
			}

			p.processNumber(number, dumperQueue)
		case <-ticker.C:
			p.DoReport()
		}
	}
}

func (p *Processor) DoReport() {
	p.reporter.MakeAReport(p.uniqNumbersProcessedFromLastReport, p.duplicateNumbersProcessesFromLastReport, p.storage.GetLength(), p.totalNumbersProcesses)
	p.uniqNumbersProcessedFromLastReport = 0
	p.duplicateNumbersProcessesFromLastReport = 0
}

func (p *Processor) processNumber(number uint64, dumperQueue chan uint64) {
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
		p.uniqNumbersProcessedFromLastReport++
	}

	p.totalNumbersProcesses++
}

func (p *Processor) terminate() {
	p.logger.Debug(fmt.Sprintf("[x] Termination inited"))
	p.triggerTerminationChannel <- "number processor"
}

func NewProcessor(numbersQueue chan uint64, storage NumberStorageInterface, reporter ReporterInterface,  triggerTerminationChannel chan string, logger LoggerInterface) *Processor {
	return &Processor{numbersQueue: numbersQueue, storage: storage, reporter: reporter, triggerTerminationChannel: triggerTerminationChannel, logger: logger}
}
