package rfc3464

import (
	"mime"
	"net/mail"
)

func getMediaInfo(hdr mail.Header) (string, error) {
	ctype := hdr.Get("Content-Type")
	mediatype, params, err := mime.ParseMediaType(ctype)

	if err == nil {
		boundary := params["boundary"]
		//if mediatype == "multipart/report" && params["report-type"] == "delivery-status" && boundary != "" {
		if mediatype == "multipart/report" && boundary != "" {
			return boundary, nil
		}
	}

	return "", ErrorInvalidContentTypeHeader
}
