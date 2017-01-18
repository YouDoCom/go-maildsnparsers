package rfc3464

import (
	"bufio"
	"net/textproto"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RecipientRecordParse(t *testing.T) {
	value := `Original-Recipient: rfc822;louisl@larry.slip.umd.edu
Final-Recipient: rfc822;louisl@larry.slip1.umd.edu
Remote-MTA: dns; sdcc13.ucsd.edu
Action: failed
Status: 4.0.0
Diagnostic-Code: smtp; 426 connection timed out
Last-Attempt-Date: Thu, 7 Jul 1994 17:15:49 -0400
Final-Log-ID: sdfskl1112
Will-Retry-Until: Thu, 7 Jul 2994 17:15:49 -0400
X-Postfix-Queue-ID: 3354017BFA8
X-Postfix-Sender: rfc822; noreply@mail.xx.com

 `

	rdr := textproto.NewReader(bufio.NewReader(strings.NewReader(value)))

	hdr, err := rdr.ReadMIMEHeader()

	if err != nil {
		t.Fatalf("Error reading mime header: %v", err)
	}

	record := RecipientRecord{}
	record.fillFromHeader(hdr)

	assert.Equal(t, "rfc822", record.OriginalRecipient.Type, "OriginalRecipient.Type")
	assert.Equal(t, "louisl@larry.slip.umd.edu", record.OriginalRecipient.Value, "OriginalRecipient.Value")

	assert.Equal(t, "rfc822", record.FinalRecipient.Type, "FinalRecipient.Type")
	assert.Equal(t, "louisl@larry.slip1.umd.edu", record.FinalRecipient.Value, "FinalRecipient.Value")

	assert.Equal(t, RecipientAction("failed"), record.Action, "Action")
	assert.True(t, record.Action.IsFailed(), "Action.IsFailed")

	assert.Equal(t, "4.0.0", record.Status, "Status")

	assert.Equal(t, "smtp", record.DiagnosticCode.Type, "DiagnosticCode.Type")
	assert.Equal(t, "426 connection timed out", record.DiagnosticCode.Value, "DiagnosticCode.Value")

	assert.Equal(t, "Thu, 7 Jul 1994 17:15:49 -0400", record.LastAttemptDate, "LastAttemptDate")

	assert.Equal(t, "dns", record.RemoteMTA.Type, "RemoteMTA.Type")
	assert.Equal(t, "sdcc13.ucsd.edu", record.RemoteMTA.Value, "RemoteMTA.Value")

	assert.Equal(t, "sdfskl1112", record.FinalLogID, "FinalLogID")

	assert.Equal(t, "Thu, 7 Jul 2994 17:15:49 -0400", record.WillRetryUntil, "WillRetryUntil")

	assert.Contains(t, record.Extensions, "X-Postfix-Queue-Id", "X-Postfix-Queue-Id")
	assert.Equal(t, "3354017BFA8", record.Extensions["X-Postfix-Queue-Id"], "X-Postfix-Queue-Id value")

	assert.Contains(t, record.Extensions, "X-Postfix-Sender", "X-Postfix-Sender")
	assert.Equal(t, "rfc822; noreply@mail.xx.com", record.Extensions["X-Postfix-Sender"], "X-Postfix-Sender value")
}
