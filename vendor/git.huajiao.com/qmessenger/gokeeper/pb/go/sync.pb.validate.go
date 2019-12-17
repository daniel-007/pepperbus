// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sync.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ConfigEvent with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConfigEvent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EventType

	// no validation rules for Data

	return nil
}

// ConfigEventValidationError is the validation error returned by
// ConfigEvent.Validate if the designated constraints aren't met.
type ConfigEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigEventValidationError) ErrorName() string { return "ConfigEventValidationError" }

// Error satisfies the builtin error interface
func (e ConfigEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigEventValidationError{}
