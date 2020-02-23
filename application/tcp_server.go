package application

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync/atomic"
)

type TcpServer struct {
	interfaceToListen         string
	maxClientsCount           int64
	triggerTerminationChannel chan string
	terminationChannel        chan struct{}
	logger                    Logger
}

func (s *TcpServer) StartListening(handler MessageHandlerInterface) {

	l, err := net.Listen("tcp4", s.interfaceToListen)
	if err != nil {
		s.logger.Error(fmt.Sprintf("%e", err))
		return
	}
	// @TODO unhandled error
	defer l.Close()

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
			continue
		}

		if clientsCounter >= s.maxClientsCount {
			// @TODO unhandled error
			c.Close()
			s.logger.Debug(fmt.Sprintf("Maximum of %d clients reached: %d", s.maxClientsCount, clientsCounter))

			continue
		}

		atomic.AddInt64(&clientsCounter, 1)
		s.logger.Debug(fmt.Sprintf("Clients count: %d", clientsCounter))

		go s.handleConnection(c, &clientsCounter, handler)
	}
}

func (s *TcpServer) handleConnection(c net.Conn, clientsCounter *int64, handler MessageHandlerInterface) {

	// @TODO unhandled error
	defer c.Close()
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
			break
		}

		message := strings.TrimSpace(rowData)
		if message == "terminate" {
			handler.Terminate()
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

func NewTcpServer(interfaceToListen string, maxClientsCount int64, triggerTerminationChannel chan string, terminationChannel chan struct{}, logger Logger) *TcpServer {
	return &TcpServer{
		interfaceToListen:         interfaceToListen,
		maxClientsCount:           maxClientsCount,
		triggerTerminationChannel: triggerTerminationChannel,
		terminationChannel:        terminationChannel,
		logger:                    logger}
}
