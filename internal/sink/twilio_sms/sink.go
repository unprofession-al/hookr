package twilio_sms

import (
	"fmt"
	"strings"

	"github.com/sfreiberg/gotwilio"
	"github.com/unprofession-al/hookr/internal/sink"
)

func init() {
	sink.Register("twilio_sms", Setup)
}

func Setup(c sink.Config) (sink.Sink, error) {
	tokens := strings.Split(c.Connection, ":")
	if len(tokens) != 4 {
		return nil, fmt.Errorf("malformend connection string, must be '[accountSid]:[accountSid]:[fromPhone]:[toPhone]'")
	}

	twilio := TwilioSMS{
		client: gotwilio.NewTwilioClient(tokens[0], tokens[1]),
		from:   tokens[2],
		to:     tokens[3],
	}

	return twilio, nil
}

type TwilioSMS struct {
	client *gotwilio.Twilio
	from   string
	to     string
}

func (t TwilioSMS) Send(message string) error {
	_, _, err := t.client.SendSMS(t.from, t.to, message, "", "")
	return err
}
