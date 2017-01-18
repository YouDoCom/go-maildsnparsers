package rfc3464

import (
	"net/textproto"
)

// Extensions represent map of extension fileds for DSN and RecipientRecord
type Extensions map[string]string

// Set sets the entry associated with key to
// the value. It replaces any existing
// values associated with key.
func (e Extensions) Set(key, value string) {
	e[textproto.CanonicalMIMEHeaderKey(key)] = value
}

// Del deletes the value associated with key.
func (e Extensions) Del(key string) {
	delete(e, textproto.CanonicalMIMEHeaderKey(key))
}

// Get gets value associated with the given key.
// If there are no values associated with the key, Get returns "".
// Get is a convenience method. For more complex queries,
// access the map directly.
func (e Extensions) Get(key string) string {
	if e == nil {
		return ""
	}
	return e[textproto.CanonicalMIMEHeaderKey(key)]
}
