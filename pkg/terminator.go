package pkg

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Terminator struct {
	triggerTerminationChannel chan string
}

func (t *Terminator) WatchForTermination(terminationChannel chan struct{}, messagesQueue chan uint64, logger logrus.Ext1FieldLogger) {
	stoppedByAndReason := <-t.triggerTerminationChannel
	logger.Debug(fmt.Sprintf("[x] Application termination initialized %s", stoppedByAndReason))
	close(terminationChannel)
	logger.Debug("[x] Termination channel closed")
	time.Sleep(1 * time.Second)
	close(messagesQueue) // I know it looks like a work around, but I did not find better solution to process all messages from the queue and quit processor
}

func (t *Terminator) Terminate(by string, reason string) {
	t.triggerTerminationChannel <- fmt.Sprintf("by: %s reason: %s", by, reason)
}

func NewTerminator() *Terminator {
	triggerTerminationChannel := make(chan string)
	return &Terminator{triggerTerminationChannel: triggerTerminationChannel}
}
