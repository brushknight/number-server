package bufio_based

import (
	"bufio"
	"fmt"
	logger2 "number-server/infrastructure/logger"
)

type Dumper struct {
	template string
	logger   logger2.LoggerInterface
}

func (d *Dumper) ProcessChannel(dumperInQueue chan uint64, writer *bufio.Writer) {
	for number := range dumperInQueue {
		_, writeError := writer.WriteString(fmt.Sprintf(d.template, number))

		if writeError != nil {
			d.logger.Critical(fmt.Sprintf("%e", writeError))
		}
	}

	flushError := writer.Flush()

	if flushError != nil {
		d.logger.Critical(fmt.Sprintf("%e", flushError))
	}
}

func NewDumper(isLeadingZeros bool, logger logger2.LoggerInterface) *Dumper {

	var template string
	if isLeadingZeros {
		template = "%09d\n"
	} else {
		template = "%d\n"
	}
	return &Dumper{template: template, logger: logger}
}
