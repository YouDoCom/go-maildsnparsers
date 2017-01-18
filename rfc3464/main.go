package rfc3464

import (
	"bufio"
	"io"
	"mime"
	"mime/multipart"
	"net/mail"
	"net/textproto"
)

// Parse parses RFC3464 Delivery Status Notification (DSN) from mail message
func Parse(message *mail.Message) (*DSN, error) {
	if message == nil {
		return nil, ErrorNilMessage
	}

	boundary, err := getMediaInfo(message.Header)

	if err != nil {
		return nil, err
	}

	r, err := findReport(boundary, message.Body)

	if err != nil {
		return nil, err
	}

	return parseReport(r)
}

func findReport(boundary string, reader io.Reader) (io.Reader, error) {
	r := multipart.NewReader(reader, boundary)

	for {
		p, err := r.NextPart()

		if err != nil {
			if err == io.EOF {
				return nil, ErrorDSNPartNotFound
			}

			return nil, err
		}

		contentHeader := p.Header.Get("Content-Type")
		mediatype, _, err := mime.ParseMediaType(contentHeader)
		if err == nil && mediatype == "message/delivery-status" {
			return p, nil
		}
	}
}

func parseReport(reader io.Reader) (*DSN, error) {
	r := textproto.NewReader(bufio.NewReader(reader))
	hdr, err := r.ReadMIMEHeader()

	if err != nil {
		return nil, err
	}

	dsn := DSN{}
	dsn.fillFromHeader(hdr)

	for {
		hdr, err = r.ReadMIMEHeader()

		if hdr != nil && len(hdr) > 0 {
			record := RecipientRecord{}
			record.fillFromHeader(hdr)

			dsn.Recipients = append(dsn.Recipients, record)
		}

		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return &dsn, err
		}
	}
}

// IsDSN checks that message is valid RFC3464 Delivery Status Notification (DSN)
func IsDSN(message *mail.Message) bool {
	if message == nil {
		return false
	}

	_, err := getMediaInfo(message.Header)
	return err == nil
}
