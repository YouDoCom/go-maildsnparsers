package xfailedrecipients

import (
	"net/mail"
	"strings"
)

// IsDSN checks that message is valid X-Failed-Recipients Delivery Status Notification (DSN)
func IsDSN(message *mail.Message) bool {
	if message == nil {
		return false
	}

	hdr := message.Header.Get("X-Failed-Recipients")
	hdr = strings.TrimSpace(hdr)

	if hdr == "" {
		return false
	}

	return true
}

// Parse parses X-Failed-Recipients Delivery Status Notification (DSN) from mail message and returs failed recipients list
func Parse(message *mail.Message) ([]string, error) {
	if message == nil {
		return nil, ErrorNilMessage
	}

	hdr := message.Header.Get("X-Failed-Recipients")
	hdr = strings.TrimSpace(hdr)

	if hdr == "" {
		return nil, ErrorDSNNotFound
	}

	recipients := strings.FieldsFunc(hdr, func(r rune) bool {
		return r == ';' || r == ','
	})

	// trim spaces from recipients
	for i := range recipients {
		recipients[i] = strings.TrimSpace(recipients[i])
	}

	return recipients, nil
}
