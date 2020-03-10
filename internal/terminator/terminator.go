package terminator

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

var triggerTerminationChannel chan string

func init() {
	triggerTerminationChannel = make(chan string)
}

func WatchForTermination(terminationChannel chan struct{}, messagesQueue chan uint64, logger logrus.Ext1FieldLogger) {
	stoppedByAndReason := <-triggerTerminationChannel
	logger.Debug(fmt.Sprintf("[x] Application termination initialized %s", stoppedByAndReason))
	close(terminationChannel)
	logger.Debug("[x] Termination channel closed")
	time.Sleep(1 * time.Second)
	close(messagesQueue) // I know it looks like a work around, but I did not find better solution to process all messages from the queue and quit processor
}

func Terminate(by string, reason string) {
	triggerTerminationChannel <- fmt.Sprintf("by: %s reason: %s", by, reason)
}
