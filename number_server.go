package main

import (
	"bufio"
	"flag"
	"fmt"
	"number-server/application"
	"number-server/domain"
	"number-server/infrastructure"
	"number-server/infrastructure/dumper/bufio_based"
	loggerStdout "number-server/infrastructure/logger/stdout"
	reporterStdout "number-server/infrastructure/reporter/stdout"
	"number-server/infrastructure/storage/memory"
	"os"
	"sync"
	"time"
)

func main() {

	maxClientsCount := flag.Int("max-clients", 5, "a number of concurrent clients")
	reporterTimeout := flag.Int("reporter-timeout", 10, "how often to write report into stdout (seconds)")
	interfaceToStart := flag.String("interface", "0.0.0.0:4000", "an interface to startup")
	env := flag.String("env", "prd", "env, allowed: prd, dev ")
	logFilePath := flag.String("log-file", "./numbers.log", "path to log a file")
	isLeadingZeros := flag.Bool("leading-zeros", true, "add leading zeros to ./numbers.log or not")

	flag.Parse()

	if *maxClientsCount < 1 {
		panic(fmt.Errorf("-max-clients cannot be less than 1, got %d", *maxClientsCount))
	}

	if *reporterTimeout < 1 {
		panic(fmt.Errorf("-reporter-timeout cannot be less than 1, got %d", *reporterTimeout))
	}

	if *env != "dev" && *env != "prd" {
		panic(fmt.Errorf("-env should be dev or prd, got %s", *env))
	}

	numbersQueue := make(chan uint64, 1000*100)
	dumperQueue := make(chan uint64, 1000*100)
	triggerTerminationChannel := make(chan string)
	terminationChannel := make(chan struct{})

	wgServer := sync.WaitGroup{}

	storage := memory.NewNumberStorage()

	logger := loggerStdout.NewLogger(*env)

	go infrastructure.Terminator(triggerTerminationChannel, terminationChannel, numbersQueue, logger)
	logger.Debug("[✔] Application terminator created")

	handler := domain.NewMessageHandler(numbersQueue, triggerTerminationChannel)
	logger.Debug("[✔] Message handler created")

	reporter := reporterStdout.NewReporter(time.Duration(*reporterTimeout*1000), logger, *env)
	logger.Debug("[✔] Reporter created")

	wgServer.Add(1)
	dumper := bufio_based.NewDumper(*isLeadingZeros, logger)
	logger.Debug("[✔] Dumper created")

	go func() {
		defer wgServer.Done()

		file := createAndOpenDumperFile(*logFilePath, logger)
		defer file.Close()

		fileWriter := bufio.NewWriter(file)

		dumper.ProcessChannel(dumperQueue, fileWriter)
	}()
	logger.Debug("[✔] Dumper started")

	wgServer.Add(1)
	processor := domain.NewProcessor(numbersQueue, storage, reporter, triggerTerminationChannel, logger)
	logger.Debug("[✔] Message processor created")

	go func() {
		defer wgServer.Done()

		processor.StartProcessing(dumperQueue)
	}()
	logger.Debug("[✔] Message processor started")

	server := application.NewTcpServer(*interfaceToStart, int64(*maxClientsCount), triggerTerminationChannel, terminationChannel, logger)
	logger.Debug("[✔] Server created")

	go server.StartListening(handler)

	wgServer.Wait()

	processor.DoReport()
}

func createAndOpenDumperFile(dumperFilePath string, logger domain.LoggerInterface) *os.File {
	var createdFile, err = os.Create(dumperFilePath)

	if err != nil {
		logger.Critical(fmt.Sprintf("%e", err))
	}
	createdFile.Close()

	file, err := os.OpenFile(dumperFilePath, os.O_RDWR, 0644)
	if err != nil {
		logger.Critical(fmt.Sprintf("%e", err))
	}

	return file
}
