package file

import (
	"fmt"
	"number-server/infrastructure/logger/stdout"
	"os"
	"sync"
)

type Dumper struct {
	filePath string
	wgServer *sync.WaitGroup
	logger   stdout.Logger
}

func (d *Dumper) ProcessChannel(c chan uint64) {

	defer d.wgServer.Done()

	var dumpChunk []uint64

	var file, err = os.OpenFile(d.filePath, os.O_RDWR, 0644)
	if err != nil {
		d.logger.Critical(fmt.Sprintf("%e", err))
	}
	defer file.Close()

	for {
		for i := 0; i < 10000; i++ {
			number, ok := <-c

			if !ok {
				d.logger.Debug("[x] Terminating dumper")
				return
			}

			dumpChunk = append(dumpChunk, number)
		}

		d.writeDownToFile(dumpChunk, file)
		dumpChunk = nil
	}
}

func (d *Dumper) writeDownToFile(dumpChunk []uint64, file *os.File) {
	var dumpChunkString string

	for i := 0; i < len(dumpChunk); i++ {
		dumpChunkString += fmt.Sprintf("%d\n", dumpChunk[i])
	}

	_, errWrite := file.WriteString(dumpChunkString)
	if errWrite != nil {
		d.logger.Critical(fmt.Sprintf("%e", errWrite))
	}

	err := file.Sync()
	if err != nil {
		d.logger.Critical(fmt.Sprintf("%e", err))
	}
}

func NewDumper(filePath string, wgServer *sync.WaitGroup, logger stdout.Logger) *Dumper {
	var file, err = os.Create(filePath)
	if err != nil {
		logger.Critical(fmt.Sprintf("%e", err))
	}
	defer file.Close()

	logger.Debug(fmt.Sprintf("[✔] File Successfully created and erased - %s", filePath))

	return &Dumper{filePath: filePath, wgServer: wgServer, logger: logger}
}
