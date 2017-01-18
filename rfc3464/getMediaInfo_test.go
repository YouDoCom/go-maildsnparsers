package rfc3464

import (
	"net/mail"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMediaInfo(t *testing.T) {
	type fixture struct {
		value         mail.Header
		expected      string
		expectedError error
	}

	fixtures := []fixture{
		fixture{
			value: mail.Header{
				"Content-Type": []string{"multipart/report; report-type=delivery-status; boundary=\"B525417BF12.1476910811/mail02.sample.com\""},
			},
			expected: `B525417BF12.1476910811/mail02.sample.com`,
		},
		fixture{
			value: mail.Header{
				"Content-Type": []string{`multipart/report; report-type=delivery-status; boundary=abcc`},
			},
			expected: `abcc`,
		},
		fixture{
			value: mail.Header{
				"Content-Type": []string{`multipart/woot; report-type=delivery-status; boundary="B525417BF12.1476910811/mail02.sample.com"`},
			},
			expectedError: ErrorInvalidContentTypeHeader,
		},
		fixture{
			value: mail.Header{
				"Content-Type": []string{`multipart/report; report-type=delivery-x; boundary=abcc`},
			},
			expected: `abcc`,
		},
		fixture{
			value: mail.Header{
				"Content-Type": []string{`multipart/report; report-type=delivery-status;`},
			},
			expectedError: ErrorInvalidContentTypeHeader,
		},
		fixture{
			value: mail.Header{
				"Woot-Type": []string{`multipart/woot; report-type=delivery-status; boundary="B525417BF12.1476910811/mail02.sample.com"`},
			},
			expectedError: ErrorInvalidContentTypeHeader,
		},
		fixture{
			value: mail.Header{
				"Content-Type": []string{`woot/"`},
			},
			expectedError: ErrorInvalidContentTypeHeader,
		},
	}

	for _, f := range fixtures {
		got, err := getMediaInfo(f.value)

		if assert.Equal(t, f.expectedError, err, "Fixture: %#v", f) {
			continue
		}

		assert.Equal(t, f.expected, got, "Fixture: %#v", f)
	}
}
