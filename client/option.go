package client

import (
	"fmt"
	"strings"
)

type Option func(client *icClient) error

func WithEventSink(url string) Option  {
	return func(client *icClient) error {
		if client == nil {
			return fmt.Errorf("evnetsink option can not set nil client")
		}
		url = strings.TrimSpace(url)
		if url != "" {
			client.url = url
			return nil
		}
		return fmt.Errorf("evnetsink option can not set empty string")
	}
}
