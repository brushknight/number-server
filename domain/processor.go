package domain

import (
	"fmt"
	"number-server/infrastructure/storage/memory"
	"sync"
	"time"
)

type Processor struct {
	numbersQueue chan uint64
	storage      NumberStorageInterface
	reporter     ReporterInterface
	dumper       DumperInterface
	wgServer     *sync.WaitGroup
	logger       LoggerInterface
}

func (p *Processor) StartProcessing() {

	defer p.wgServer.Done()

	ticker := time.NewTicker(p.reporter.GetFrequencyMs() * time.Millisecond)
	defer ticker.Stop()

	dumperQueue := make(chan uint64)
	defer close(dumperQueue)

	var uniqNumbersProcesses uint64 = 0
	var duplicateNumbersProcesses uint64 = 0
	var totalNumbersProcesses uint64 = 0

	go p.dumper.ProcessChannel(dumperQueue)

	for {
		select {
		case number, ok := <-p.numbersQueue:

			if !ok {
				p.logger.Debug("[x] Terminating messages processor")
				return
			}

			totalNumbersProcesses++

			if p.storage.IsNumberExists(number) {
				duplicateNumbersProcesses++
			} else {
				status, err := p.storage.AddNumber(number)

				if err != nil {
					p.logger.Error(fmt.Sprintf("%e\n", err))
				}
				if !status {
					p.logger.Error(fmt.Sprintf("number was not inserted %d\n", number))
				}

				dumperQueue <- number
				uniqNumbersProcesses++
			}

		case <-ticker.C:
			p.reporter.MakeAReport(uniqNumbersProcesses, duplicateNumbersProcesses, p.storage.GetLength(), totalNumbersProcesses)
			uniqNumbersProcesses = 0
			duplicateNumbersProcesses = 0
		}
	}
}

func NewProcessor(numbersQueue chan uint64, reporter ReporterInterface, dumper DumperInterface, wgServer *sync.WaitGroup, logger LoggerInterface) *Processor {

	storage := memory.NewNumberStorage()
	return &Processor{numbersQueue: numbersQueue, storage: storage, reporter: reporter, dumper: dumper, wgServer: wgServer, logger: logger}
}
