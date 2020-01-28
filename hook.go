package main

import (
	"fmt"

	"github.com/unprofession-al/hookr/internal/sink"
	_ "github.com/unprofession-al/hookr/internal/sink/twilio_sms"

	"github.com/unprofession-al/hookr/internal/decoder"
	_ "github.com/unprofession-al/hookr/internal/decoder/noop"
	_ "github.com/unprofession-al/hookr/internal/decoder/pingdom"
)

type Hooks map[string]*Hook

func (h Hooks) Prepare() []error {
	var errs []error
	for name, hook := range h {
		d, err := decoder.New(hook.Decoder)
		if err != nil {
			errs = append(errs, fmt.Errorf("cound not prepare decoder for hook '%s', error was '%s'", name, err.Error()))
		}
		s, err := sink.New(hook.Sink)
		if err != nil {
			errs = append(errs, fmt.Errorf("cound not prepare sink for hook '%s', error was '%s'", name, err.Error()))
		}
		hook.decoder = d
		hook.sink = s
	}
	return errs
}

type Hook struct {
	Decoder decoder.Config `yaml:"decoder"`
	Sink    sink.Config    `yaml:"sink"`

	decoder decoder.Decoder
	sink    sink.Sink
}

func (h *Hook) Process(in []byte) error {
	msg, err := h.decoder.Decode(in)
	if err != nil {
		return err
	}

	err = h.sink.Send(string(msg))
	return err
}
