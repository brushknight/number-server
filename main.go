package main

import (
	"flag"
	"number-server/cmd"
)

func main() {
	maxClientsCount := flag.Int("max-clients", 5, "a number of concurrent clients")
	reporterTimeout := flag.Int("reporter-timeout", 10, "how often to write report into stdout (seconds)")
	interfaceToStart := flag.String("interface", "0.0.0.0:4000", "an interface to startup")
	env := flag.String("env", "prd", "env, allowed: prd, dev ")
	logFilePath := flag.String("log-file", "./numbers.log", "path to log a file")
	isLeadingZeros := flag.Bool("leading-zeros", true, "add leading zeros to ./numbers.log or not")

	flag.Parse()

	cmd.CreateServerAndStart(*maxClientsCount, *reporterTimeout, *interfaceToStart, *env, *logFilePath, *isLeadingZeros)
}
