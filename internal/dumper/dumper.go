package dumper

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
)

func ProcessChannel(dumperInQueue chan uint64, writer *bufio.Writer, isLeadingZeros bool, logger logrus.Ext1FieldLogger) {

	var template string
	if isLeadingZeros {
		template = "%09d\n"
	} else {
		template = "%d\n"
	}

	logger.Debug(fmt.Sprintf("Dumper strated with template %s", template))

	for number := range dumperInQueue {
		logger.Trace(fmt.Sprintf("Dumper received: %d", number))

		_, writeError := writer.WriteString(fmt.Sprintf(template, number))

		if writeError != nil {
			logger.Fatal(fmt.Sprintf("%e", writeError))
		}
	}

	flushError := writer.Flush()

	if flushError != nil {
		logger.Fatal(fmt.Sprintf("%e", flushError))
	}
}
