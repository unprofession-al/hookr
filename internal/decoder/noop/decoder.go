package noop

import (
	"github.com/unprofession-al/hookr/internal/decoder"
)

func init() {
	decoder.Register("noop", Setup)
}

func Setup(c decoder.Config) (decoder.Decoder, error) {
	return Noop{}, nil
}

type Noop struct{}

func (n Noop) Decode(message []byte) ([]byte, error) {
	return message, nil
}
