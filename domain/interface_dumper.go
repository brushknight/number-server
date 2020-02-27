package domain

import "bufio"

type DumperInterface interface {
	ProcessChannel(inChannel chan uint64, writer *bufio.Writer)
}
