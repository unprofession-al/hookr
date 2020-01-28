package pingdom

import (
	"encoding/json"
	"fmt"

	"github.com/unprofession-al/hookr/internal/decoder"
)

func init() {
	decoder.Register("pingdom", Setup)
}

func Setup(c decoder.Config) (decoder.Decoder, error) {
	return Pingdom{}, nil
}

type Pingdom struct{}

func (n Pingdom) Decode(message []byte) ([]byte, error) {
	var d Message
	err := json.Unmarshal(message, &d)
	out := fmt.Sprintf("pingdom alert: check %s went from %s to %s", d.CheckName, d.PreviousState, d.CurrentState)
	return []byte(out), err
}

type Message struct {
	CheckID     int    `json:"check_id"`
	CheckName   string `json:"check_name"`
	CheckType   string `json:"check_type"`
	CheckParams struct {
		Hostname  string `json:"hostname"`
		BasicAuth bool   `json:"basic_auth"`
		Ipv6      bool   `json:"ipv6"`
		Port      int    `json:"port"`
	} `json:"check_params"`
	PreviousState string `json:"previous_state"`
	CurrentState  string `json:"current_state"`
}
