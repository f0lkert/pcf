package processor

import (
	"github.com/f0lkert/pcf/internal/sbi/consumer"
	"github.com/f0lkert/pcf/pkg/app"
)

type PCF interface {
	app.App
	Consumer() *consumer.Consumer
}

type Processor struct {
	PCF
}

func NewProcessor(pcf PCF) (*Processor, error) {
	p := &Processor{
		PCF: pcf,
	}

	return p, nil
}
