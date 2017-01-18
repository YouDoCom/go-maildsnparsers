package rfc3464

import (
	"net/mail"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Parse(b *testing.B) {
	value := `Date: Thu, 7 Jul 1994 17:16:05 -0400
From: Mail Delivery Subsystem <MAILER-DAEMON@CS.UTK.EDU>
Message-Id: <199407072116.RAA14128@CS.UTK.EDU>
Subject: Returned mail: Cannot send message for 5 days
To: <owner-info-mime@cs.utk.edu>
MIME-Version: 1.0
Content-Type: multipart/report; report-type=delivery-status;
 boundary="RAA14128.773615765/CS.UTK.EDU"

--RAA14128.773615765/CS.UTK.EDU

The original message was received at Sat, 2 Jul 1994 17:10:28 -0400
from root@localhost

    ----- The following addresses had delivery problems -----
<louisl@larry.slip.umd.edu>  (unrecoverable error)

----- Transcript of session follows -----
<louisl@larry.slip.umd.edu>... Deferred: Connection timed out
            with larry.slip.umd.edu.
Message could not be delivered for 5 days
Message will be deleted from queue

--RAA14128.773615765/CS.UTK.EDU
content-type: message/delivery-status

Original-Envelope-Id: my-envelope
Reporting-MTA: dns; cs.utk.edu
DSN-Gateway: dns; cs.utk.edu2
Received-From-MTA: dns; cs.xx.ss
Arrival-Date: Thu, 7 Jul 1994 17:15:49 -0401
X-Postfix-Queue-ID: 3354017BFA8
X-Postfix-Sender: rfc822; noreply@mail.xx.com

Original-Recipient: rfc822;louisl@larry.slip.umd.edu
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

--RAA14128.773615765/CS.UTK.EDU
content-type: message/rfc822

[original message goes here]

--RAA14128.773615765/CS.UTK.EDU--
`

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(msg)
	}
}

func Test_Parse_simple(t *testing.T) {
	value := `Date: Thu, 7 Jul 1994 17:16:05 -0400
From: Mail Delivery Subsystem <MAILER-DAEMON@CS.UTK.EDU>
Message-Id: <199407072116.RAA14128@CS.UTK.EDU>
Subject: Returned mail: Cannot send message for 5 days
To: <owner-info-mime@cs.utk.edu>
MIME-Version: 1.0
Content-Type: multipart/report; report-type=delivery-status;
 boundary="RAA14128.773615765/CS.UTK.EDU"

--RAA14128.773615765/CS.UTK.EDU

The original message was received at Sat, 2 Jul 1994 17:10:28 -0400
from root@localhost

    ----- The following addresses had delivery problems -----
<louisl@larry.slip.umd.edu>  (unrecoverable error)

----- Transcript of session follows -----
<louisl@larry.slip.umd.edu>... Deferred: Connection timed out
            with larry.slip.umd.edu.
Message could not be delivered for 5 days
Message will be deleted from queue

--RAA14128.773615765/CS.UTK.EDU
content-type: message/delivery-status

Original-Envelope-Id: my-envelope
Reporting-MTA: dns; cs.utk.edu
DSN-Gateway: dns; cs.utk.edu2
Received-From-MTA: dns; cs.xx.ss
Arrival-Date: Thu, 7 Jul 1994 17:15:49 -0401
X-Postfix-Queue-ID: 3354017BFA8
X-Postfix-Sender: rfc822; noreply@mail.xx.com

Original-Recipient: rfc822;louisl@larry.slip.umd.edu
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

--RAA14128.773615765/CS.UTK.EDU
content-type: message/rfc822

[original message goes here]

--RAA14128.773615765/CS.UTK.EDU--
`

	msg, _ := mail.ReadMessage(strings.NewReader(value))
	dsn, err := Parse(msg)

	assert.NoError(t, err)

	assert.Equal(t, "my-envelope", dsn.OriginalEnvelopeID)

	assert.Equal(t, "dns", dsn.ReportingMTA.Type)
	assert.Equal(t, "cs.utk.edu", dsn.ReportingMTA.Value)

	assert.Equal(t, "dns", dsn.DsnGateway.Type)
	assert.Equal(t, "cs.utk.edu2", dsn.DsnGateway.Value)

	assert.Equal(t, "dns", dsn.ReceivedFromMTA.Type)
	assert.Equal(t, "cs.xx.ss", dsn.ReceivedFromMTA.Value)

	assert.Equal(t, "Thu, 7 Jul 1994 17:15:49 -0401", dsn.ArrivalDate)

	assert.Contains(t, dsn.Extensions, "X-Postfix-Queue-Id")
	assert.Equal(t, "3354017BFA8", dsn.Extensions["X-Postfix-Queue-Id"])

	assert.Contains(t, dsn.Extensions, "X-Postfix-Sender")
	assert.Equal(t, "rfc822; noreply@mail.xx.com", dsn.Extensions["X-Postfix-Sender"])

	assert.Len(t, dsn.Recipients, 1)

	// Check record
	record := dsn.Recipients[0]

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

func Test_Parse_NilMessage(t *testing.T) {
	_, err := Parse(nil)

	assert.Equal(t, ErrorNilMessage, err)
}

func Test_IsDSN_Valid(t *testing.T) {
	value := `Date: Thu, 7 Jul 1994 17:16:05 -0400
From: Mail Delivery Subsystem <MAILER-DAEMON@CS.UTK.EDU>
Message-Id: <199407072116.RAA14128@CS.UTK.EDU>
Subject: Returned mail: Cannot send message for 5 days
To: <owner-info-mime@cs.utk.edu>
MIME-Version: 1.0
Content-Type: multipart/report; report-type=delivery-status;
 boundary="RAA14128.773615765/CS.UTK.EDU"

--RAA14128.773615765/CS.UTK.EDU

The original message was received at Sat, 2 Jul 1994 17:10:28 -0400
from root@localhost

    ----- The following addresses had delivery problems -----
<louisl@larry.slip.umd.edu>  (unrecoverable error)

----- Transcript of session follows -----
<louisl@larry.slip.umd.edu>... Deferred: Connection timed out
            with larry.slip.umd.edu.
Message could not be delivered for 5 days
Message will be deleted from queue

--RAA14128.773615765/CS.UTK.EDU
content-type: message/delivery-status

Original-Envelope-Id: my-envelope
Reporting-MTA: dns; cs.utk.edu
DSN-Gateway: dns; cs.utk.edu2
Received-From-MTA: dns; cs.xx.ss
Arrival-Date: Thu, 7 Jul 1994 17:15:49 -0401
X-Postfix-Queue-ID: 3354017BFA8
X-Postfix-Sender: rfc822; noreply@mail.xx.com

Original-Recipient: rfc822;louisl@larry.slip.umd.edu
Final-Recipient: rfc822;louisl@larry.slip.umd.edu
Action: failed
Status: 4.0.0
Diagnostic-Code: smtp; 426 connection timed out
Last-Attempt-Date: Thu, 7 Jul 1994 17:15:49 -0400

--RAA14128.773615765/CS.UTK.EDU
content-type: message/rfc822

[original message goes here]

--RAA14128.773615765/CS.UTK.EDU--
`

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	assert.True(t, IsDSN(msg), "IsDSN valid")
}

func Test_IsDSN_InValid(t *testing.T) {
	value := `X-Failed-Recipients: igorportoviy@mail.ru
X-Mailer-Daemon-Recipients: igorportoviy@mail.ru
Auto-Submitted: auto-replied
From: mailer-daemon@corp.mail.ru
To: noreply@xxx.ss.dd
Content-Transfer-Encoding: 8bit
Content-Type: text/plain; charset=utf-8
Subject: Mail failure.
Message-Id: <E1cDsQ7-0000By-U3@mx169.mail.ru>
Date: Mon, 05 Dec 2016 15:30:59 +0300
`

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	assert.False(t, IsDSN(msg), "IsDSN invalid")
}

func Test_Parse_InvalidCotentType(t *testing.T) {
	value := `Date: Thu, 7 Jul 1994 17:16:05 -0400
From: Mail Delivery Subsystem <MAILER-DAEMON@CS.UTK.EDU>
Message-Id: <199407072116.RAA14128@CS.UTK.EDU>
Subject: Returned mail: Cannot send message for 5 days
To: <owner-info-mime@cs.utk.edu>
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8

This message was created automatically by mail delivery software.

A message that you sent could not be delivered to one or more of its
recipients. This is a permanent error. The following address(es) failed:

  user@xxx.ss
    messages count limit 350000, msg count in mailbox is 350000
`

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	_, err := Parse(msg)
	assert.EqualError(t, err, ErrorInvalidContentTypeHeader.Error())
}

func Test_Parse_InvalidDSNPartNotFound(t *testing.T) {
	value := `Date: Thu, 7 Jul 1994 17:16:05 -0400
From: Mail Delivery Subsystem <MAILER-DAEMON@CS.UTK.EDU>
Message-Id: <199407072116.RAA14128@CS.UTK.EDU>
Subject: Returned mail: Cannot send message for 5 days
To: <owner-info-mime@cs.utk.edu>
MIME-Version: 1.0
Content-Type: multipart/report; report-type=delivery-status;
 boundary="RAA14128.773615765/CS.UTK.EDU"

--RAA14128.773615765/CS.UTK.EDU

The original message was received at Sat, 2 Jul 1994 17:10:28 -0400
from root@localhost

    ----- The following addresses had delivery problems -----
<louisl@larry.slip.umd.edu>  (unrecoverable error)

----- Transcript of session follows -----
<louisl@larry.slip.umd.edu>... Deferred: Connection timed out
            with larry.slip.umd.edu.
Message could not be delivered for 5 days
Message will be deleted from queue

content-type: message/rfc822

[original message goes here]

--RAA14128.773615765/CS.UTK.EDU--
`

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	_, err := Parse(msg)
	assert.EqualError(t, err, ErrorDSNPartNotFound.Error())
}

func TestReaderError(t *testing.T) {
	value := `Date: Thu, 7 Jul 1994 17:16:05 -0400
From: Mail Delivery Subsystem <MAILER-DAEMON@CS.UTK.EDU>
Message-Id: <199407072116.RAA14128@CS.UTK.EDU>
Subject: Returned mail: Cannot send message for 5 days
To: <owner-info-mime@cs.utk.edu>
MIME-Version: 1.0
Content-Type: multipart/report; report-type=delivery-status;
 boundary="RAA14128.773615765/CS.UTK.EDU-woot"

--RAA14128.773615765/CS.UTK.EDU

The original message was received at Sat, 2 Jul 1994 17:10:28 -0400
from root@localhost

    ----- The following addresses had delivery problems -----
<louisl@larry.slip.umd.edu>  (unrecoverable error)

----- Transcript of session follows -----
<louisl@larry.slip.umd.edu>... Deferred: Connection timed out
            with larry.slip.umd.edu.
Message could not be delivered for 5 days
Message will be deleted from queue

--RAA14128.773615765/CS.UTK.EDU
content-type: message/delivery-status

Original-Envelope-Id: my-envelope
Reporting-MTA: dns; cs.utk.edu
DSN-Gateway: dns; cs.utk.edu2
Received-From-MTA: dns; cs.xx.ss
Arrival-Date: Thu, 7 Jul 1994 17:15:49 -0401
X-Postfix-Queue-ID: 3354017BFA8
X-Postfix-Sender: rfc822; noreply@mail.xx.com

Original-Recipient: rfc822;louisl@larry.slip.umd.edu
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

--RAA14128.773615765/CS.UTK.EDU
content-type: message/rfc822

[original message goes here]

--RAA14128.773615765/CS.UTK.EDU--
`

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	_, err := Parse(msg)
	assert.Error(t, err)
}
