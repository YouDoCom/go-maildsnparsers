package xfailedrecipients

import "errors"

var (
	// ErrorNilMessage returned when message is nil
	ErrorNilMessage = errors.New("Message is nil")

	// ErrorDSNNotFound retured when DSN cannot be found in message
	ErrorDSNNotFound = errors.New("DSN not found in message")
)
