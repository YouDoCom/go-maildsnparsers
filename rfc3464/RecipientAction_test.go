package rfc3464

import "testing"
import "github.com/stretchr/testify/assert"

type TestRecipientActionFixture struct {
	value    string
	expected bool
}

func Test_RecipientAction_IsFailed(t *testing.T) {
	fixtures := []TestRecipientActionFixture{
		TestRecipientActionFixture{
			value:    "failed",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "Failed",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "",
			expected: false,
		},
	}

	for _, f := range fixtures {
		data := RecipientAction(f.value)
		got := data.IsFailed()

		assert.Equal(t, f.expected, got)
	}
}

func Test_RecipientAction_IsDelayed(t *testing.T) {
	fixtures := []TestRecipientActionFixture{
		TestRecipientActionFixture{
			value:    "delayed",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "Delayed",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "",
			expected: false,
		},
	}

	for _, f := range fixtures {
		data := RecipientAction(f.value)
		got := data.IsDelayed()

		assert.Equal(t, f.expected, got)
	}
}

func Test_RecipientAction_IsDelivered(t *testing.T) {
	fixtures := []TestRecipientActionFixture{
		TestRecipientActionFixture{
			value:    "delivered",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "Delivered",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "",
			expected: false,
		},
	}

	for _, f := range fixtures {
		data := RecipientAction(f.value)
		got := data.IsDelivered()

		assert.Equal(t, f.expected, got)
	}
}

func Test_RecipientAction_IsRelayed(t *testing.T) {
	fixtures := []TestRecipientActionFixture{
		TestRecipientActionFixture{
			value:    "relayed",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "Relayed",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "",
			expected: false,
		},
	}

	for _, f := range fixtures {
		data := RecipientAction(f.value)
		got := data.IsRelayed()

		assert.Equal(t, f.expected, got)
	}
}

func Test_RecipientAction_IsExpanded(t *testing.T) {
	fixtures := []TestRecipientActionFixture{
		TestRecipientActionFixture{
			value:    "expanded",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "Expanded",
			expected: true,
		},
		TestRecipientActionFixture{
			value:    "",
			expected: false,
		},
	}

	for _, f := range fixtures {
		data := RecipientAction(f.value)
		got := data.IsExpanded()

		assert.Equal(t, f.expected, got)
	}
}
