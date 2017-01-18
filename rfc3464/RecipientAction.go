package rfc3464

import (
	"strings"
)

/*
RecipientAction represents type for Action field of Recipient
*/
type RecipientAction string

/*
IsFailed indicates that the message could not be delivered to the
recipient.  The Reporting MTA has abandoned any attempts
to deliver the message to this recipient.  No further
notifications should be expected.
*/
func (a RecipientAction) IsFailed() bool {
	return strings.ToLower(string(a)) == "failed"
}

/*
IsDelayed indicates that the Reporting MTA has so far been unable
to deliver or relay the message, but it will continue to
attempt to do so.  Additional notification messages may
be issued as the message is further delayed or
successfully delivered, or if delivery attempts are later
abandoned.
*/
func (a RecipientAction) IsDelayed() bool {
	return strings.ToLower(string(a)) == "delayed"
}

/*
IsDelivered indicates that the message was successfully delivered to
the recipient address specified by the sender, which
includes "delivery" to a mailing list exploder.  It does
not indicate that the message has been read.  This is a
terminal state and no further DSN for this recipient
should be expected.
*/
func (a RecipientAction) IsDelivered() bool {
	return strings.ToLower(string(a)) == "delivered"
}

/*
IsRelayed indicates that the message has been relayed or gatewayed
into an environment that does not accept responsibility
for generating DSNs upon successful delivery.  This
action-value SHOULD NOT be used unless the sender has
requested notification of successful delivery for this
recipient.
*/
func (a RecipientAction) IsRelayed() bool {
	return strings.ToLower(string(a)) == "relayed"
}

/*
IsExpanded indicates that the message has been successfully
delivered to the recipient address as specified by the
sender, and forwarded by the Reporting-MTA beyond that
destination to multiple additional recipient addresses.
An action-value of "expanded" differs from "delivered" in
that "expanded" is not a terminal state.  Further
"failed" and/or "delayed" notifications may be provided.
*/
func (a RecipientAction) IsExpanded() bool {
	return strings.ToLower(string(a)) == "expanded"
}
