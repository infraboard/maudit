// Code generated by github.com/infraboard/mcube/v2
// DO NOT EDIT

package event

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseLevelFromString Parse Level from string
func ParseLevelFromString(str string) (Level, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Level_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Level: %s", str)
	}

	return Level(v), nil
}

// Equal type compare
func (t Level) Equal(target Level) bool {
	return t == target
}

// IsIn todo
func (t Level) IsIn(targets ...Level) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t Level) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Level) UnmarshalJSON(b []byte) error {
	ins, err := ParseLevelFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseTypeFromString Parse Type from string
func ParseTypeFromString(str string) (Type, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := Type_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Type: %s", str)
	}

	return Type(v), nil
}

// Equal type compare
func (t Type) Equal(target Type) bool {
	return t == target
}

// IsIn todo
func (t Type) IsIn(targets ...Type) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t Type) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *Type) UnmarshalJSON(b []byte) error {
	ins, err := ParseTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseContentTypeFromString Parse ContentType from string
func ParseContentTypeFromString(str string) (ContentType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ContentType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ContentType: %s", str)
	}

	return ContentType(v), nil
}

// Equal type compare
func (t ContentType) Equal(target ContentType) bool {
	return t == target
}

// IsIn todo
func (t ContentType) IsIn(targets ...ContentType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ContentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ContentType) UnmarshalJSON(b []byte) error {
	ins, err := ParseContentTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
