package stdout

import (
	"fmt"

	"github.com/unprofession-al/hookr/internal/sink"
)

func init() {
	sink.Register("stdout", Setup)
}

func Setup(c sink.Config) (sink.Sink, error) {
	return StdOut{}, nil
}

type StdOut struct{}

func (s StdOut) Send(message string) error {
	fmt.Println(message)
	return nil
}
