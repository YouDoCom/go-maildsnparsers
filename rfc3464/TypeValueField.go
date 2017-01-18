package rfc3464

import (
	"fmt"
	"strings"
)

/*
TypeValueField represent field type which present
in DSN and RecipientRecord

Example:
    Field: type; value
*/
type TypeValueField struct {
	// Type of value
	Type string
	// Value
	Value string
}

// String returns string representation of TypeValueField
func (tvf TypeValueField) String() string {
	return fmt.Sprintf("%s; %s", tvf.Type, tvf.Value)
}

// ParseTypeValueField parses TypeValueField from string
func ParseTypeValueField(value string) TypeValueField {
	data := strings.SplitN(value, ";", 2)

	return TypeValueField{
		Type:  strings.TrimSpace(data[0]),
		Value: strings.TrimSpace(data[1]),
	}
}
