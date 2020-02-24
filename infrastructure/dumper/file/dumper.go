package file

import (
	"bufio"
	"fmt"
	"number-server/infrastructure/logger"
	"os"
	"sync"
)

type Dumper struct {
	filePath       string
	wgServer       *sync.WaitGroup
	logger         logger.LoggerInterface
	writerTemplate string
}

func (d *Dumper) ProcessChannel(c chan uint64) {

	defer d.wgServer.Done()

	var file, err = os.OpenFile(d.filePath, os.O_RDWR, 0644)
	if err != nil {
		d.logger.Critical(fmt.Sprintf("%e", err))
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	defer d.syncWriterStream(w)

	for {
		for i := 0; i < 10000; i++ {
			number, ok := <-c

			if !ok {
				d.logger.Debug("[x] Terminating dumper")
				return
			}

			d.writeNumberViaWriterStream(number, w)
		}

		d.syncWriterStream(w)
	}
}

func (d *Dumper) writeNumberViaWriterStream(number uint64, writer *bufio.Writer) int {
	numberOfBytesWritten, _ := writer.WriteString(fmt.Sprintf(d.writerTemplate, number))
	return numberOfBytesWritten
}

func (d *Dumper) syncWriterStream(writer *bufio.Writer) bool {
	errWriter := writer.Flush()
	if errWriter != nil {
		d.logger.Critical(fmt.Sprintf("%e", errWriter))
		return false
	}

	return true
}

func NewDumper(filePath string, wgServer *sync.WaitGroup, logger logger.LoggerInterface, isLeadingZeros bool) *Dumper {
	var file, err = os.Create(filePath)

	if err != nil {
		logger.Critical(fmt.Sprintf("%e", err))
	}
	defer file.Close()

	logger.Debug(fmt.Sprintf("[✔] File Successfully created and erased - %s", filePath))

	var writerTemplate string
	if isLeadingZeros {
		writerTemplate = "%09d\n"
	} else {
		writerTemplate = "%d\n"
	}

	return &Dumper{filePath: filePath, wgServer: wgServer, logger: logger, writerTemplate: writerTemplate}
}
