package rfc3464

import "errors"

var (
	// ErrorNilMessage returned when message is nil
	ErrorNilMessage = errors.New("Message is nil")

	// ErrorInvalidContentTypeHeader returned when Content-Type header not valid
	//
	// Valid examples:
	// - multipart/report; report-type=delivery-status; boundary="RAA14128.773615765/CS.UTK.EDU"
	// - multipart/report; report-type="delivery-status"; boundary="RAA14128.773615765/CS.UTK.EDU"
	ErrorInvalidContentTypeHeader = errors.New("Invalid Content-Type header")

	// ErrorDSNPartNotFound retured when "message/delivery-status" part cannot be found in message body
	ErrorDSNPartNotFound = errors.New("DSN part not found in message body")
)
