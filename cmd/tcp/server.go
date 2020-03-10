package tcp

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"number-server/cmd/handler"
	"number-server/internal/terminator"
	"strings"
	"sync/atomic"
)

type Server struct {
	logger                    logrus.Ext1FieldLogger
	interfaceToListen         string
	maxClientsCount           int64
	triggerTerminationChannel chan string
	terminationChannel        chan struct{}
}

func (s *Server) StartListening(handler handler.MessageHandlerInterface, triggerTerminationChannel chan string) {

	l, err := net.Listen("tcp4", s.interfaceToListen)
	if err != nil {
		s.logger.Error(fmt.Sprintf("%e", err))
		terminator.Terminate(triggerTerminationChannel, "Server", fmt.Sprintf("%e", err))
		return
	}
	defer func() {
		err := l.Close()
		if err != nil {
			s.logger.Error(fmt.Sprintf("%e", err))
			terminator.Terminate(triggerTerminationChannel, "Server", fmt.Sprintf("%e", err))
		}
	}()

	s.logger.Debug(fmt.Sprintf("[✔] Server started serving on %s, max %d clients", s.interfaceToListen, s.maxClientsCount))

	var clientsCounter int64

	for {
		select {
		case <-s.terminationChannel:
			s.logger.Debug(fmt.Sprintf("App termination, stopping server"))
			return
		default:

		}
		c, err := l.Accept()

		if err != nil {
			s.logger.Error(fmt.Sprintf("%e", err))
			terminator.Terminate(triggerTerminationChannel, "Server", fmt.Sprintf("%e", err))
			continue
		}

		if clientsCounter >= s.maxClientsCount {
			err := c.Close()
			if err != nil {
				s.logger.Error(fmt.Sprintf("%e", err))
				terminator.Terminate(triggerTerminationChannel, "Server", fmt.Sprintf("%e", err))
			}

			s.logger.Debug(fmt.Sprintf("Maximum of %d clients reached: %d", s.maxClientsCount, clientsCounter))

			continue
		}

		atomic.AddInt64(&clientsCounter, 1)
		s.logger.Debug(fmt.Sprintf("Clients count: %d", clientsCounter))

		go s.handleConnection(handler, c, &clientsCounter, triggerTerminationChannel)
	}
}

func (s *Server) handleConnection(handler handler.MessageHandlerInterface, c net.Conn, clientsCounter *int64, triggerTerminationChannel chan string) {

	defer func() {
		err := c.Close()
		if err != nil {
			s.logger.Error(fmt.Sprintf("%e", err))
			terminator.Terminate(triggerTerminationChannel, "Server", fmt.Sprintf("%e", err))
		}
	}()
	defer atomic.AddInt64(clientsCounter, -1)

	channelReader := bufio.NewReader(c)
	for {
		select {
		case <-s.terminationChannel:
			s.logger.Debug(fmt.Sprintf("[x] App termination, client disconnecting"))
			return
		default:
		}

		rowData, err := channelReader.ReadString('\n')

		if err != nil {

			if err.Error() == "EOF" {
				// connection closed, stop execution of handler
				break
			}

			s.logger.Error(fmt.Sprintf("Error: %e", err))
			terminator.Terminate(triggerTerminationChannel, "Server", fmt.Sprintf("%e", err))
			break
		}

		message := strings.TrimSpace(rowData)
		if message == "terminate" {
			terminator.Terminate(triggerTerminationChannel, "Server", "termination sequence received")
			return
		}

		number, err := handler.ValidateAndParse(message)

		if err != nil {
			s.logger.Error(fmt.Sprintf("error occurred while parsing message: %e", err))
			return
		}

		select {
		case <-s.terminationChannel:
			s.logger.Debug(fmt.Sprintf("[x] App termination, client disconnecting"))
			return
		default:
			handler.Handle(number)
		}
	}
}

func NewServer(interfaceToListen string, maxClientsCount int64, triggerTerminationChannel chan string, terminationChannel chan struct{}, logger logrus.Ext1FieldLogger) *Server {
	return &Server{
		interfaceToListen:         interfaceToListen,
		maxClientsCount:           maxClientsCount,
		triggerTerminationChannel: triggerTerminationChannel,
		terminationChannel:        terminationChannel,
		logger:                    logger}
}
