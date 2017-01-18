package rfc3464

import (
	"net/textproto"
	"strings"
)

/*
RecipientRecord represents per-recipient DSN record

	A DSN contains information about attempts to deliver a message to one
	or more recipients.  The delivery information for any particular
	recipient is contained in a group of contiguous per-recipient fields.
	Each group of per-recipient fields is preceded by a blank line.

	The syntax for the group of per-recipient fields is as follows:

	per-recipient-fields =
			[ original-recipient-field CRLF ]
			final-recipient-field CRLF
			action-field CRLF
			status-field CRLF
			[ remote-mta-field CRLF ]
			[ diagnostic-code-field CRLF ]
			[ last-attempt-date-field CRLF ]
			[ final-log-id-field CRLF ]
			[ will-retry-until-field CRLF ]
			*( extension-field CRLF )
*/
type RecipientRecord struct {
	/*
		2.3.1 Original-Recipient field

		The Original-Recipient field indicates the original recipient address
		as specified by the sender of the message for which the DSN is being
		issued.

			original-recipient-field =
					"Original-Recipient" ":" address-type ";" generic-address

			generic-address = *text

		The address-type field indicates the type of the original recipient
		address.  If the message originated within the Internet, the
		address-type field will normally be "rfc822", and the address will be
		according to the syntax specified in [RFC822].  The value "unknown"
		should be used if the Reporting MTA cannot determine the type of the
		original recipient address from the message envelope.

		This field is optional.  It should be included only if the sender-
		specified recipient address was present in the message envelope, such
		as by the SMTP extensions defined in [RFC3461].  This address is the
		same as that provided by the sender and can be used to automatically
		correlate DSN reports and message transactions.
	*/
	OriginalRecipient TypeValueField

	/*
		2.3.2 Final-Recipient field

		The Final-Recipient field indicates the recipient for which this set
		of per-recipient fields applies.  This field MUST be present in each
		set of per-recipient data.

		The syntax of the field is as follows:

			final-recipient-field =
				"Final-Recipient" ":" address-type ";" generic-address

		The generic-address sub-field of the Final-Recipient field MUST
		contain the mailbox address of the recipient (from the transport
		envelope), as it was when the Reporting MTA accepted the message for
		delivery.

		The Final-Recipient address may differ from the address originally
		provided by the sender, because it may have been transformed during
		forwarding and gatewaying into a totally unrecognizable mess.
		However, in the absence of the optional Original-Recipient field, the
		Final-Recipient field and any returned content may be the only
		information available with which to correlate the DSN with a
		particular message submission.

		The address-type sub-field indicates the type of address expected by
		the reporting MTA in that context.  Recipient addresses obtained via
		SMTP will normally be of address-type "rfc822".

		NOTE: The Reporting MTA is not expected to ensure that the address
		actually conforms to the syntax conventions of the address-type.
		Instead, it MUST report exactly the address received in the envelope,
		unless that address contains characters such as CR or LF which are
		not allowed in a DSN field.

		Since mailbox addresses (including those used in the Internet) may be
		case sensitive, the case of alphabetic characters in the address MUST
		be preserved.

	*/
	FinalRecipient TypeValueField

	/*
		2.3.3 Action field

		The Action field indicates the action performed by the Reporting-MTA
		as a result of its attempt to deliver the message to this recipient
		address.  This field MUST be present for each recipient named in the
		DSN.

		The syntax for the action-field is:

			action-field = "Action" ":" action-value

			action-value =
				"failed" / "delayed" / "delivered" / "relayed" / "expanded"

		The action-value may be spelled in any combination of upper and lower
		case characters.

		"failed"    indicates that the message could not be delivered to the
					recipient.  The Reporting MTA has abandoned any attempts
					to deliver the message to this recipient.  No further
					notifications should be expected.

		"delayed"   indicates that the Reporting MTA has so far been unable
					to deliver or relay the message, but it will continue to
					attempt to do so.  Additional notification messages may
					be issued as the message is further delayed or
					successfully delivered, or if delivery attempts are later
					abandoned.

		"delivered" indicates that the message was successfully delivered to
					the recipient address specified by the sender, which
					includes "delivery" to a mailing list exploder.  It does
					not indicate that the message has been read.  This is a
					terminal state and no further DSN for this recipient
					should be expected.


		"relayed"   indicates that the message has been relayed or gatewayed
					into an environment that does not accept responsibility
					for generating DSNs upon successful delivery.  This
					action-value SHOULD NOT be used unless the sender has
					requested notification of successful delivery for this
					recipient.

		"expanded"  indicates that the message has been successfully
					delivered to the recipient address as specified by the
					sender, and forwarded by the Reporting-MTA beyond that
					destination to multiple additional recipient addresses.
					An action-value of "expanded" differs from "delivered" in
					that "expanded" is not a terminal state.  Further
					"failed" and/or "delayed" notifications may be provided.

		Using the terms "mailing list" and "alias" as defined in [RFC3461],
		section 7.2.7: An action-value of "expanded" is only to be used when
		the message is delivered to a multiple-recipient "alias".  An
		action-value of "expanded" SHOULD NOT be used with a DSN issued on
		delivery of a message to a "mailing list".

		NOTE ON ACTION VS. STATUS CODES: Although the 'action' field
			might seem to be redundant with the 'status' field, this is not
			the case.  In particular, a "temporary failure" ("4") status code
			could be used with an action-value of either "delayed" or
			"failed".  For example, assume that an SMTP client repeatedly
			tries to relay a message to the mail exchanger for a recipient,
			but fails because a query to a domain name server timed out.

			After a few hours, it might issue a "delayed" DSN to inform the
			sender that the message had not yet been delivered.  After a few
			days, the MTA might abandon its attempt to deliver the message
			and return a "failed" DSN.  The status code (which would begin
			with a "4" to indicate "temporary failure") would be the same for
			both DSNs.

			Another example for which the action and status codes may appear
			contradictory: If an MTA or mail gateway cannot deliver a message
			because doing so would entail conversions resulting in an
			unacceptable loss of information, it would issue a DSN with the
			'action' field of "failure" and a status code of 'XXX'.  If the
			message had instead been relayed, but with some loss of
			information, it might generate a DSN with the same XXX status-
			code, but with an action field of "relayed".
	*/
	Action RecipientAction

	/*
		2.3.4 Status field

		The per-recipient Status field contains a transport-independent
		status code that indicates the delivery status of the message to that
		recipient.  This field MUST be present for each delivery attempt
		which is described by a DSN.

		The syntax of the status field is:

		status-field = "Status" ":" status-code

		status-code = DIGIT "." 1*3DIGIT "." 1*3DIGIT

			; White-space characters and comments are NOT allowed within
			; a status-code, though a comment enclosed in parentheses
			; MAY follow the last numeric sub-field of the status-code.
			; Each numeric sub-field within the status-code MUST be
			; expressed without leading zero digits.

		Status codes thus consist of three numerical fields separated by ".".
		The first sub-field indicates whether the delivery attempt was
		successful (2= success, 4 = persistent temporary failure, 5 =
		permanent failure).  The second sub-field indicates the probable
		source of any delivery anomalies, and the third sub-field denotes a
		precise error condition, if known.

		The initial set of status-codes is defined in [RFC3463].

	*/
	Status string

	/*
		2.3.5 Remote-MTA field

		The value associated with the Remote-MTA DSN field is a printable
		ASCII representation of the name of the "remote" MTA that reported
		delivery status to the "reporting" MTA.

			remote-mta-field = "Remote-MTA" ":" mta-name-type ";" mta-name

		NOTE: The Remote-MTA field preserves the "while talking to"
		information that was provided in some pre-existing nondelivery
		reports.

		This field is optional.  It MUST NOT be included if no remote MTA was
		involved in the attempted delivery of the message to that recipient.
	*/
	RemoteMTA TypeValueField

	/*
		2.3.6 Diagnostic-Code field

		For a "failed" or "delayed" recipient, the Diagnostic-Code DSN field
		contains the actual diagnostic code issued by the mail transport.
		Since such codes vary from one mail transport to another, the
		diagnostic-type sub-field is needed to specify which type of
		diagnostic code is represented.

		diagnostic-code-field =
				"Diagnostic-Code" ":" diagnostic-type ";" *text

		NOTE: The information in the Diagnostic-Code field may be somewhat
		redundant with that from the Status field.  The Status field is
		needed so that any DSN, regardless of origin, may be understood by
		any user agent or gateway that parses DSNs.  Since the Status code
		will sometimes be less precise than the actual transport diagnostic
		code, the Diagnostic-Code field is provided to retain the latter
		information.  Such information may be useful in a trouble ticket sent
		to the administrator of the Reporting MTA, or when tunneling foreign
		non-delivery reports through DSNs.

		If the Diagnostic Code was obtained from a Remote MTA during an
		attempt to relay the message to that MTA, the Remote-MTA field should
		be present.  When interpreting a DSN, the presence of a Remote-MTA
		field indicates that the Diagnostic Code was issued by the Remote
		MTA.  The absence of a Remote-MTA indicates that the Diagnostic Code
		was issued by the Reporting MTA.

		In addition to the Diagnostic-Code itself, additional textual
		description of the diagnostic, MAY appear in a comment enclosed in
		parentheses.

		This field is optional, because some mail systems supply no
		additional information beyond that which is returned in the 'action'
		and 'status' fields.  However, this field SHOULD be included if
		transport-specific diagnostic information is available.
	*/
	DiagnosticCode TypeValueField

	/*
		2.3.7 Last-Attempt-Date field

		The Last-Attempt-Date field gives the date and time of the last
		attempt to relay, gateway, or deliver the message (whether successful
		or unsuccessful) by the Reporting MTA.  This is not necessarily the
		same as the value of the Date field from the header of the message
		used to transmit this delivery status notification: In cases where
		the DSN was generated by a gateway, the Date field in the message
		header contains the time the DSN was sent by the gateway and the DSN
		Last-Attempt-Date field contains the time the last delivery attempt
		occurred.

			last-attempt-date-field = "Last-Attempt-Date" ":" date-time

		This field is optional.  It MUST NOT be included if the actual date
		and time of the last delivery attempt are not available (which might
		be the case if the DSN were being issued by a gateway).

		The date and time are expressed in RFC 822 'date-time' format, as
		modified by [HOSTREQ].  Numeric timezones ([+/-]HHMM format) MUST be
		used.
	*/
	LastAttemptDate string

	/*
		2.3.8 final-log-id field

		The "final-log-id" field gives the final-log-id of the message that
		was used by the final-mta.  This can be useful as an index to the
		final-mta's log entry for that delivery attempt.

			final-log-id-field = "Final-Log-ID" ":" *text

		This field is optional.
	*/
	FinalLogID string

	/*
		2.3.9 Will-Retry-Until field

		For DSNs of type "delayed", the Will-Retry-Until field gives the date
		after which the Reporting MTA expects to abandon all attempts to
		deliver the message to that recipient.  The Will-Retry-Until field is
		optional for "delay" DSNs, and MUST NOT appear in other DSNs.

			will-retry-until-field = "Will-Retry-Until" ":" date-time

		The date and time are expressed in RFC 822 'date-time' format, as
		modified by [RFC1123].  Numeric timezones ([+/-]HHMM format) MUST be
		used.
	*/
	WillRetryUntil string

	/*
		2.4 Extension fields

		Additional per-message or per-recipient DSN fields may be defined in
		the future by later revisions or extensions to this specification.
		Extension-field names beginning with "X-" will never be defined as
		standard fields;  such names are reserved for experimental use.  DSN
		field names NOT beginning with "X-" MUST be registered with the
		Internet Assigned Numbers Authority (IANA) and published in an RFC.

		Extension DSN fields may be defined for the following reasons:

		(a) To allow additional information from foreign delivery status
			reports to be tunneled through Internet DSNs.  The names of such
			DSN fields should begin with an indication of the foreign
			environment name (e.g., X400-Physical-Forwarding-Address).

		(b) To allow the transmission of diagnostic information which is
			specific to a particular mail transport protocol.  The names of
			such DSN fields should begin with an indication of the mail
			transport being used (e.g., SMTP-Remote-Recipient-Address).  Such
			fields should be used for diagnostic purposes only and not by
			user agents or mail gateways.

		(c) To allow transmission of diagnostic information which is specific
			to a particular message transfer agent (MTA).  The names of such
			DSN fields should begin with an indication of the MTA
			implementation that produced the DSN. (e.g., Foomail-Queue-ID).

		MTA implementers are encouraged to provide adequate information, via
		extension fields if necessary, to allow an MTA maintainer to
		understand the nature of correctable delivery failures and how to fix
		them.  For example, if message delivery attempts are logged, the DSN
		might include information that allows the MTA maintainer to easily
		find the log entry for a failed delivery attempt.

		If an MTA developer does not wish to register the meanings of such
		extension fields, "X-" fields may be used for this purpose.  To avoid
		name collisions, the name of the MTA implementation should follow the
		"X-", (e.g., "X-Foomail-Log-ID").
	*/
	Extensions Extensions
}

func (record *RecipientRecord) fillFromHeader(hdr textproto.MIMEHeader) {
	record.Extensions = make(Extensions)

	var (
		keyOriginalRecipient = textproto.CanonicalMIMEHeaderKey("Original-Recipient")
		keyFinalRecipient    = textproto.CanonicalMIMEHeaderKey("Final-Recipient")
		keyAction            = textproto.CanonicalMIMEHeaderKey("Action")
		keyStatus            = textproto.CanonicalMIMEHeaderKey("Status")
		keyRemoteMTA         = textproto.CanonicalMIMEHeaderKey("Remote-MTA")
		keyDiagnosticCode    = textproto.CanonicalMIMEHeaderKey("Diagnostic-Code")
		keyLastAttemptDate   = textproto.CanonicalMIMEHeaderKey("Last-Attempt-Date")
		keyFinalLogID        = textproto.CanonicalMIMEHeaderKey("Final-Log-ID")
		keyWillRetryUntil    = textproto.CanonicalMIMEHeaderKey("Will-Retry-Until")
	)

	for k, v := range hdr {
		val := strings.Join(v, "\n")

		switch k {
		case keyOriginalRecipient:
			record.OriginalRecipient = ParseTypeValueField(val)
		case keyFinalRecipient:
			record.FinalRecipient = ParseTypeValueField(val)
		case keyAction:
			record.Action = RecipientAction(val)
		case keyStatus:
			record.Status = val
		case keyRemoteMTA:
			record.RemoteMTA = ParseTypeValueField(val)
		case keyDiagnosticCode:
			record.DiagnosticCode = ParseTypeValueField(val)
		case keyLastAttemptDate:
			record.LastAttemptDate = val
		case keyFinalLogID:
			record.FinalLogID = val
		case keyWillRetryUntil:
			record.WillRetryUntil = val

		default:
			record.Extensions.Set(k, val)
		}
	}
}
