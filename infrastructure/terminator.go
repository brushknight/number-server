package infrastructure

import (
	"fmt"
	"number-server/infrastructure/logger/stdout"
	"time"
)

func Terminator(triggerTerminationChannel chan string, terminationChannel chan struct{}, messagesQueue chan uint64, logger stdout.Logger) {
	stoppedBy := <-triggerTerminationChannel
	logger.Debug(fmt.Sprintf("[x] Application termination initialized by: %s", stoppedBy))
	close(terminationChannel)
	logger.Debug("[x] Termination channel closed")
	time.Sleep(1 * time.Second)
	close(messagesQueue) // I know it looks like a work around, but I did not find better solution to process all messages from the queue and quit processor
}
