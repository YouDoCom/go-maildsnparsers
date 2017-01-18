package rfc3464

import "testing"
import "github.com/stretchr/testify/assert"

func Test_ParseTypeValueField(t *testing.T) {
	type fixture struct {
		value    string
		expected TypeValueField
	}

	fixtures := []fixture{
		fixture{
			value: `rfc822;boing7071@e1.sample`,
			expected: TypeValueField{
				Type:  "rfc822",
				Value: "boing7071@e1.sample",
			},
		},
		fixture{
			value: `rfc822; vision_foto@mail.sample`,
			expected: TypeValueField{
				Type:  "rfc822",
				Value: "vision_foto@mail.sample",
			},
		},
		fixture{
			value: `dns; mxs.mail.sample`,
			expected: TypeValueField{
				Type:  "dns",
				Value: "mxs.mail.sample",
			},
		},
		fixture{
			value: "smtp; 550 Message was not accepted -- invalid mailbox.  Local\nmailbox tankist.zao@sample.sample is unavailable: user is terminated",
			expected: TypeValueField{
				Type:  "smtp",
				Value: "550 Message was not accepted -- invalid mailbox.  Local\nmailbox tankist.zao@sample.sample is unavailable: user is terminated",
			},
		},
	}

	for _, f := range fixtures {
		got := ParseTypeValueField(f.value)
		assert.Equal(t, f.expected, got)
	}
}

func Test_TypeValue_String(t *testing.T) {
	type fixture struct {
		expected string
		value    TypeValueField
	}

	fixtures := []fixture{
		fixture{
			expected: `rfc822; boing7071@e1.sample`,
			value: TypeValueField{
				Type:  "rfc822",
				Value: "boing7071@e1.sample",
			},
		},
		fixture{
			expected: `rfc822; vision_foto@mail.sample`,
			value: TypeValueField{
				Type:  "rfc822",
				Value: "vision_foto@mail.sample",
			},
		},
		fixture{
			expected: `dns; mxs.mail.sample`,
			value: TypeValueField{
				Type:  "dns",
				Value: "mxs.mail.sample",
			},
		},
		fixture{
			expected: "smtp; 550 Message was not accepted -- invalid mailbox.  Local\nmailbox tankist.zao@sample.sample is unavailable: user is terminated",
			value: TypeValueField{
				Type:  "smtp",
				Value: "550 Message was not accepted -- invalid mailbox.  Local\nmailbox tankist.zao@sample.sample is unavailable: user is terminated",
			},
		},
	}

	for _, f := range fixtures {
		got := f.value.String()
		assert.Equal(t, f.expected, got)
	}
}
