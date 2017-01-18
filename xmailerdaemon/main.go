package xmailerdaemon

import (
	"net/mail"
	"strings"
)

// IsDSN checks that message is valid Delivery Status Notification (DSN)
func IsDSN(message *mail.Message) bool {
	if message == nil {
		return false
	}

	hdr := message.Header.Get("X-Mailer-Daemon-Recipients")
	hdr = strings.TrimSpace(hdr)

	reasonHdr := message.Header.Get("X-Mailer-Daemon-Error")
	reasonHdr = strings.TrimSpace(reasonHdr)

	if hdr == "" || reasonHdr == "" {
		return false
	}

	return true
}

// Parse parses Delivery Status Notification (DSN) from mail message and returs failed recipients list
func Parse(message *mail.Message) ([]Result, error) {
	if message == nil {
		return nil, ErrorNilMessage
	}

	hdr := message.Header.Get("X-Mailer-Daemon-Recipients")
	hdr = strings.TrimSpace(hdr)

	reasonHdr := message.Header.Get("X-Mailer-Daemon-Error")
	reasonHdr = strings.TrimSpace(reasonHdr)

	if hdr == "" || reasonHdr == "" {
		return nil, ErrorDSNNotFound
	}

	recipients := strings.FieldsFunc(hdr, func(r rune) bool {
		return r == ';' || r == ','
	})

	var ret []Result

	// trim spaces from recipients
	for i := range recipients {
		ret = append(ret, Result{
			Address: strings.TrimSpace(recipients[i]),
			Reason:  reasonHdr,
		})
	}

	return ret, nil
}
