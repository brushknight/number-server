package mock

import "sync"

type Dumper struct {
	wgServer      *sync.WaitGroup
	NumbersDumped []uint64
}

func (d *Dumper) ProcessChannel(c chan uint64) {
	defer d.wgServer.Done()
	for number := range c {
		d.NumbersDumped = append(d.NumbersDumped, number)
	}
}

func NewMockDumper(wgServer *sync.WaitGroup) *Dumper {
	return &Dumper{wgServer: wgServer}
}
