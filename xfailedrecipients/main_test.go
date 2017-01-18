package xfailedrecipients

import (
	"net/mail"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	value := `Delivered-To: from-user@example.com
Received: by 10.182.29.99 with SMTP id j3csp1775010obh;
        Mon, 5 Dec 2016 09:08:13 -0800 (PST)
X-Received: by 10.46.72.18 with SMTP id v18mr20566600lja.12.1480957693436;
        Mon, 05 Dec 2016 09:08:13 -0800 (PST)
Return-Path: <>
Received: from mail02.example.com (mail02.example.com. [8.8.8.8])
        by mx.google.com with ESMTP id 86si7519844lfs.419.2016.12.05.09.08.13
        for <from-user@example.com>;
        Mon, 05 Dec 2016 09:08:13 -0800 (PST)
Received-SPF: pass (google.com: domain of postmaster@mail02.example.com designates 8.8.8.8 as permitted sender) client-ip=8.8.8.8;
Authentication-Results: mx.google.com;
       spf=pass (google.com: domain of postmaster@mail02.example.com designates 8.8.8.8 as permitted sender) smtp.helo=mail02.example.com
Received: from example.com (mail.example.local [172.22.0.26]) by mail02.example.com (Postfix) with ESMTPS id 1CCB91600AC for <from-user@example.com>; Mon,
  5 Dec 2016 20:08:13 +0300 (MSK)
Received: from Debian-exim by example.com with local id 1cDwkO-0008LO-Th for from-user@example.com; Mon, 05 Dec 2016 20:08:12 +0300
X-Failed-Recipients: adajdasdadsadas@example.com,
  asdadsadsasdasdasdasda@example.com
Auto-Submitted: auto-replied
From: Mail Delivery System <Mailer-Daemon@mail01>
To: from-user@example.com
Subject: Mail delivery failed: returning message to sender
Message-Id: <E1cDwkO-0008LO-Th@example.com>
Date: Mon, 05 Dec 2016 20:08:12 +0300

This message was created automatically by mail delivery software.

A message that you sent could not be delivered to one or more of its
recipients. This is a permanent error. The following address(es) failed:

  adajdasdadsadas@example.com
    Unrouteable address
  asdadsadsasdasdasdasda@example.com
    Unrouteable address

------ This is a copy of the message, including all the headers. ------

Return-path: <from-user@example.com>
Received: from localhost ([::1] helo=local)
	by example.com with smtp (envelope-from <from-user@example.com>)
	id 1cDwkI-0008LG-PH; Mon, 05 Dec 2016 20:08:12 +0300
Subject: test
Message-Id: <E1cDwkI-0008LG-PH@example.com>
From: from-user@example.com
Date: Mon, 05 Dec 2016 20:08:12 +0300

a test `

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	recipients, err := Parse(msg)

	assert.NoError(t, err)

	assert.Len(t, recipients, 2, "Recipients len")
	assert.EqualValues(t, []string{"adajdasdadsadas@example.com", "asdadsadsasdasdasdasda@example.com"}, recipients)
}

func Test_ParseNilMessage(t *testing.T) {
	_, err := Parse(nil)

	assert.EqualError(t, err, ErrorNilMessage.Error(), "Nil message")
}

func Test_ParseInvalid(t *testing.T) {
	value := `Delivered-To: from-user@example.com
Received: by 10.182.29.99 with SMTP id j3csp1775010obh;
        Mon, 5 Dec 2016 09:08:13 -0800 (PST)
X-Received: by 10.46.72.18 with SMTP id v18mr20566600lja.12.1480957693436;
        Mon, 05 Dec 2016 09:08:13 -0800 (PST)
Return-Path: <>
Received: from mail02.example.com (mail02.example.com. [8.8.8.8])
        by mx.google.com with ESMTP id 86si7519844lfs.419.2016.12.05.09.08.13
        for <from-user@example.com>;
        Mon, 05 Dec 2016 09:08:13 -0800 (PST)
Received-SPF: pass (google.com: domain of postmaster@mail02.example.com designates 8.8.8.8 as permitted sender) client-ip=8.8.8.8;
Authentication-Results: mx.google.com;
       spf=pass (google.com: domain of postmaster@mail02.example.com designates 8.8.8.8 as permitted sender) smtp.helo=mail02.example.com
Received: from example.com (mail.example.local [172.22.0.26]) by mail02.example.com (Postfix) with ESMTPS id 1CCB91600AC for <from-user@example.com>; Mon,
  5 Dec 2016 20:08:13 +0300 (MSK)
Received: from Debian-exim by example.com with local id 1cDwkO-0008LO-Th for from-user@example.com; Mon, 05 Dec 2016 20:08:12 +0300
Auto-Submitted: auto-replied
From: Mail Delivery System <Mailer-Daemon@mail01>
To: from-user@example.com
Subject: Mail delivery failed: returning message to sender
Message-Id: <E1cDwkO-0008LO-Th@example.com>
Date: Mon, 05 Dec 2016 20:08:12 +0300

a test `

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	_, err := Parse(msg)

	assert.EqualError(t, err, ErrorDSNNotFound.Error(), "DSN not found")
}

func Test_IsDSNValid(t *testing.T) {
	value := `Delivered-To: from-user@example.com
Received: by 10.182.29.99 with SMTP id j3csp1775010obh;
        Mon, 5 Dec 2016 09:08:13 -0800 (PST)
X-Received: by 10.46.72.18 with SMTP id v18mr20566600lja.12.1480957693436;
        Mon, 05 Dec 2016 09:08:13 -0800 (PST)
Return-Path: <>
Received: from mail02.example.com (mail02.example.com. [8.8.8.8])
        by mx.google.com with ESMTP id 86si7519844lfs.419.2016.12.05.09.08.13
        for <from-user@example.com>;
        Mon, 05 Dec 2016 09:08:13 -0800 (PST)
Received-SPF: pass (google.com: domain of postmaster@mail02.example.com designates 8.8.8.8 as permitted sender) client-ip=8.8.8.8;
Authentication-Results: mx.google.com;
       spf=pass (google.com: domain of postmaster@mail02.example.com designates 8.8.8.8 as permitted sender) smtp.helo=mail02.example.com
Received: from example.com (mail.example.local [172.22.0.26]) by mail02.example.com (Postfix) with ESMTPS id 1CCB91600AC for <from-user@example.com>; Mon,
  5 Dec 2016 20:08:13 +0300 (MSK)
Received: from Debian-exim by example.com with local id 1cDwkO-0008LO-Th for from-user@example.com; Mon, 05 Dec 2016 20:08:12 +0300
X-Failed-Recipients: adajdasdadsadas@example.com,
  asdadsadsasdasdasdasda@example.com
Auto-Submitted: auto-replied
From: Mail Delivery System <Mailer-Daemon@mail01>
To: from-user@example.com
Subject: Mail delivery failed: returning message to sender
Message-Id: <E1cDwkO-0008LO-Th@example.com>
Date: Mon, 05 Dec 2016 20:08:12 +0300

This message was created automatically by mail delivery software.

A message that you sent could not be delivered to one or more of its
recipients. This is a permanent error. The following address(es) failed:

  adajdasdadsadas@example.com
    Unrouteable address
  asdadsadsasdasdasdasda@example.com
    Unrouteable address

------ This is a copy of the message, including all the headers. ------

Return-path: <from-user@example.com>
Received: from localhost ([::1] helo=local)
	by example.com with smtp (envelope-from <from-user@example.com>)
	id 1cDwkI-0008LG-PH; Mon, 05 Dec 2016 20:08:12 +0300
Subject: test
Message-Id: <E1cDwkI-0008LG-PH@example.com>
From: from-user@example.com
Date: Mon, 05 Dec 2016 20:08:12 +0300

a test `

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	assert.True(t, IsDSN(msg), "DSN Valid")
}

func Test_IsDSNInValid(t *testing.T) {
	value := `Delivered-To: from-user@example.com
Received: by 10.182.29.99 with SMTP id j3csp1775010obh;
        Mon, 5 Dec 2016 09:08:13 -0800 (PST)
X-Received: by 10.46.72.18 with SMTP id v18mr20566600lja.12.1480957693436;
        Mon, 05 Dec 2016 09:08:13 -0800 (PST)
Return-Path: <>
Received: from mail02.example.com (mail02.example.com. [8.8.8.8])
        by mx.google.com with ESMTP id 86si7519844lfs.419.2016.12.05.09.08.13
        for <from-user@example.com>;
        Mon, 05 Dec 2016 09:08:13 -0800 (PST)
Received-SPF: pass (google.com: domain of postmaster@mail02.example.com designates 8.8.8.8 as permitted sender) client-ip=8.8.8.8;
Authentication-Results: mx.google.com;
       spf=pass (google.com: domain of postmaster@mail02.example.com designates 8.8.8.8 as permitted sender) smtp.helo=mail02.example.com
Received: from example.com (mail.example.local [172.22.0.26]) by mail02.example.com (Postfix) with ESMTPS id 1CCB91600AC for <from-user@example.com>; Mon,
  5 Dec 2016 20:08:13 +0300 (MSK)
Received: from Debian-exim by example.com with local id 1cDwkO-0008LO-Th for from-user@example.com; Mon, 05 Dec 2016 20:08:12 +0300
Auto-Submitted: auto-replied
From: Mail Delivery System <Mailer-Daemon@mail01>
To: from-user@example.com
Subject: Mail delivery failed: returning message to sender
Message-Id: <E1cDwkO-0008LO-Th@example.com>
Date: Mon, 05 Dec 2016 20:08:12 +0300

a test `

	msg, _ := mail.ReadMessage(strings.NewReader(value))

	assert.False(t, IsDSN(msg), "DSN Invalid")
}

func Test_IsDSNNilMessage(t *testing.T) {
	assert.False(t, IsDSN(nil), "DSN nil Message")
}
