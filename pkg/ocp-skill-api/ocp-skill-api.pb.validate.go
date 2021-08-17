// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-skill-api/ocp-skill-api.proto

package ocp_skill_api

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

// Validate checks the field values on CreateSkillRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSkillRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for Name

	return nil
}

// CreateSkillRequestV1ValidationError is the validation error returned by
// CreateSkillRequestV1.Validate if the designated constraints aren't met.
type CreateSkillRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSkillRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSkillRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSkillRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSkillRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSkillRequestV1ValidationError) ErrorName() string {
	return "CreateSkillRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSkillRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSkillRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSkillRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSkillRequestV1ValidationError{}

// Validate checks the field values on CreateSkillResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSkillResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSkill()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSkillResponseV1ValidationError{
				field:  "Skill",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateSkillResponseV1ValidationError is the validation error returned by
// CreateSkillResponseV1.Validate if the designated constraints aren't met.
type CreateSkillResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSkillResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSkillResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSkillResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSkillResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSkillResponseV1ValidationError) ErrorName() string {
	return "CreateSkillResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSkillResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSkillResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSkillResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSkillResponseV1ValidationError{}

// Validate checks the field values on DescribeSkillRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeSkillRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DescribeSkillRequestV1ValidationError is the validation error returned by
// DescribeSkillRequestV1.Validate if the designated constraints aren't met.
type DescribeSkillRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeSkillRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeSkillRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeSkillRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeSkillRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeSkillRequestV1ValidationError) ErrorName() string {
	return "DescribeSkillRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeSkillRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeSkillRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeSkillRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeSkillRequestV1ValidationError{}

// Validate checks the field values on DescribeSkillResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeSkillResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSkill()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeSkillResponseV1ValidationError{
				field:  "Skill",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeSkillResponseV1ValidationError is the validation error returned by
// DescribeSkillResponseV1.Validate if the designated constraints aren't met.
type DescribeSkillResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeSkillResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeSkillResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeSkillResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeSkillResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeSkillResponseV1ValidationError) ErrorName() string {
	return "DescribeSkillResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeSkillResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeSkillResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeSkillResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeSkillResponseV1ValidationError{}

// Validate checks the field values on ListSkillsRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListSkillsRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	return nil
}

// ListSkillsRequestV1ValidationError is the validation error returned by
// ListSkillsRequestV1.Validate if the designated constraints aren't met.
type ListSkillsRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSkillsRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSkillsRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSkillsRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSkillsRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSkillsRequestV1ValidationError) ErrorName() string {
	return "ListSkillsRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e ListSkillsRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSkillsRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSkillsRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSkillsRequestV1ValidationError{}

// Validate checks the field values on ListSkillsResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListSkillsResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSkills() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSkillsResponseV1ValidationError{
					field:  fmt.Sprintf("Skills[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListSkillsResponseV1ValidationError is the validation error returned by
// ListSkillsResponseV1.Validate if the designated constraints aren't met.
type ListSkillsResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSkillsResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSkillsResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSkillsResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSkillsResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSkillsResponseV1ValidationError) ErrorName() string {
	return "ListSkillsResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e ListSkillsResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSkillsResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSkillsResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSkillsResponseV1ValidationError{}

// Validate checks the field values on RemoveSkillRequestV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveSkillRequestV1) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// RemoveSkillRequestV1ValidationError is the validation error returned by
// RemoveSkillRequestV1.Validate if the designated constraints aren't met.
type RemoveSkillRequestV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSkillRequestV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSkillRequestV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSkillRequestV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSkillRequestV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSkillRequestV1ValidationError) ErrorName() string {
	return "RemoveSkillRequestV1ValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSkillRequestV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSkillRequestV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSkillRequestV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSkillRequestV1ValidationError{}

// Validate checks the field values on RemoveSkillResponseV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveSkillResponseV1) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// RemoveSkillResponseV1ValidationError is the validation error returned by
// RemoveSkillResponseV1.Validate if the designated constraints aren't met.
type RemoveSkillResponseV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSkillResponseV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSkillResponseV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSkillResponseV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSkillResponseV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSkillResponseV1ValidationError) ErrorName() string {
	return "RemoveSkillResponseV1ValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSkillResponseV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSkillResponseV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSkillResponseV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSkillResponseV1ValidationError{}

// Validate checks the field values on Skill with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Skill) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for Name

	return nil
}

// SkillValidationError is the validation error returned by Skill.Validate if
// the designated constraints aren't met.
type SkillValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SkillValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SkillValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SkillValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SkillValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SkillValidationError) ErrorName() string { return "SkillValidationError" }

// Error satisfies the builtin error interface
func (e SkillValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSkill.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SkillValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SkillValidationError{}