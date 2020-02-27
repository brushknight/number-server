package mock

import "bufio"

type Dumper struct {
	NumbersDumped []uint64
}

func (d *Dumper) ProcessChannel(dumperInQueue chan uint64, writer *bufio.Writer) {
	for number := range dumperInQueue {
		d.NumbersDumped = append(d.NumbersDumped, number)
	}
}

func NewMockDumper() *Dumper {
	return &Dumper{}
}
