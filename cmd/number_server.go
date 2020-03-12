package cmd

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"number-server/cmd/tcp"
	"number-server/internal/dumper"
	handler2 "number-server/internal/handler"
	processor2 "number-server/internal/processor"
	"number-server/internal/reporter"
	storage2 "number-server/internal/storage"
	"number-server/pkg"
	"os"
	"sync"
	"time"
)

func CreateServerAndStart(maxClientsCount int, reporterTimeout int, interfaceToStart string, env string, logFilePath string, isLeadingZeros bool) {

	if maxClientsCount < 1 {
		panic(fmt.Errorf("-max-clients cannot be less than 1, got %d", maxClientsCount))
	}

	if reporterTimeout < 1 {
		panic(fmt.Errorf("-reporter-timeout cannot be less than 1, got %d", reporterTimeout))
	}

	if env != "dev" && env != "prd" {
		panic(fmt.Errorf("-env should be dev or prd, got %s", env))
	}

	numbersQueue := make(chan uint64, 1000*100)
	dumperQueue := make(chan uint64, 1000*100)
	reportsQueue := make(chan reporter.ReportDTO, 10)
	triggerTerminationChannel := make(chan string)
	terminationChannel := make(chan struct{})
	reportTriggerTicker := time.NewTicker(time.Duration(reporterTimeout) * time.Second)

	logger := logrus.New()

	if env == "dev" {
		logger.SetLevel(logrus.TraceLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	wgServer := sync.WaitGroup{}

	storage := storage2.NewNumberStorage()
	terminator := pkg.NewTerminator()

	go terminator.WatchForTermination(terminationChannel, numbersQueue, logger)
	logger.Debug("[✔] Application terminator created")

	handler := handler2.NewMessageHandler(numbersQueue)
	logger.Debug("[✔] Message handler created")

	wgServer.Add(1)
	go func() {
		defer wgServer.Done()

		reporter.ProcessReportsChannel(reportsQueue, os.Stdout, env)
	}()

	wgServer.Add(1)
	go func() {
		defer wgServer.Done()

		file := createAndOpenDumperFile(logFilePath, logger)
		defer file.Close()

		fileWriter := bufio.NewWriter(file)

		dumper.ProcessChannel(dumperQueue, fileWriter, isLeadingZeros, logger)
	}()

	wgServer.Add(1)
	processor := processor2.NewProcessor(storage, logger)
	logger.Debug("[✔] Message processor created")

	go func() {
		defer wgServer.Done()
		defer reportTriggerTicker.Stop()
		processor.ProcessChannel(numbersQueue, dumperQueue, reportTriggerTicker.C, reportsQueue, triggerTerminationChannel)
	}()
	logger.Debug("[✔] Message processor started")

	server := tcp.NewServer(interfaceToStart, int64(maxClientsCount), triggerTerminationChannel, terminationChannel, logger, terminator)
	logger.Debug("[✔] Server created")

	go server.StartListening(handler, triggerTerminationChannel)

	wgServer.Wait()
}

func createAndOpenDumperFile(dumperFilePath string, logger logrus.Ext1FieldLogger) *os.File {
	var createdFile, err = os.Create(dumperFilePath)

	if err != nil {
		logger.Fatal(fmt.Sprintf("%e", err))
	}
	createdFile.Close()

	file, err := os.OpenFile(dumperFilePath, os.O_RDWR, 0644)
	if err != nil {
		logger.Fatal(fmt.Sprintf("%e", err))
	}

	return file
}
