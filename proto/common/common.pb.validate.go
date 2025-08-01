// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/common.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Paginate with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Paginate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Paginate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PaginateMultiError, or nil
// if none found.
func (m *Paginate) ValidateAll() error {
	return m.validate(true)
}

func (m *Paginate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.AfterId != nil {
		// no validation rules for AfterId
	}

	if m.BeforeId != nil {
		// no validation rules for BeforeId
	}

	if m.Page != nil {
		// no validation rules for Page
	}

	if m.PageSize != nil {
		// no validation rules for PageSize
	}

	if len(errors) > 0 {
		return PaginateMultiError(errors)
	}

	return nil
}

// PaginateMultiError is an error wrapping multiple validation errors returned
// by Paginate.ValidateAll() if the designated constraints aren't met.
type PaginateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginateMultiError) AllErrors() []error { return m }

// PaginateValidationError is the validation error returned by
// Paginate.Validate if the designated constraints aren't met.
type PaginateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginateValidationError) ErrorName() string { return "PaginateValidationError" }

// Error satisfies the builtin error interface
func (e PaginateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginateValidationError{}

// Validate checks the field values on PaginatedResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PaginatedResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PaginatedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PaginatedResponseMultiError, or nil if none found.
func (m *PaginatedResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PaginatedResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.TotalPages != nil {
		// no validation rules for TotalPages
	}

	if m.TotalCounts != nil {
		// no validation rules for TotalCounts
	}

	if m.TotalItems != nil {
		// no validation rules for TotalItems
	}

	if m.NextId != nil {
		// no validation rules for NextId
	}

	if m.PrevId != nil {
		// no validation rules for PrevId
	}

	if len(errors) > 0 {
		return PaginatedResponseMultiError(errors)
	}

	return nil
}

// PaginatedResponseMultiError is an error wrapping multiple validation errors
// returned by PaginatedResponse.ValidateAll() if the designated constraints
// aren't met.
type PaginatedResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginatedResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginatedResponseMultiError) AllErrors() []error { return m }

// PaginatedResponseValidationError is the validation error returned by
// PaginatedResponse.Validate if the designated constraints aren't met.
type PaginatedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginatedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginatedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginatedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginatedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginatedResponseValidationError) ErrorName() string {
	return "PaginatedResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PaginatedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginatedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginatedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginatedResponseValidationError{}

// Validate checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Filter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FilterMultiError, or nil if none found.
func (m *Filter) ValidateAll() error {
	return m.validate(true)
}

func (m *Filter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FilterValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FilterValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FilterMultiError(errors)
	}

	return nil
}

// FilterMultiError is an error wrapping multiple validation errors returned by
// Filter.ValidateAll() if the designated constraints aren't met.
type FilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FilterMultiError) AllErrors() []error { return m }

// FilterValidationError is the validation error returned by Filter.Validate if
// the designated constraints aren't met.
type FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterValidationError) ErrorName() string { return "FilterValidationError" }

// Error satisfies the builtin error interface
func (e FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterValidationError{}

// Validate checks the field values on Sort with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Sort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Sort with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SortMultiError, or nil if none found.
func (m *Sort) ValidateAll() error {
	return m.validate(true)
}

func (m *Sort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	// no validation rules for Order

	if len(errors) > 0 {
		return SortMultiError(errors)
	}

	return nil
}

// SortMultiError is an error wrapping multiple validation errors returned by
// Sort.ValidateAll() if the designated constraints aren't met.
type SortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SortMultiError) AllErrors() []error { return m }

// SortValidationError is the validation error returned by Sort.Validate if the
// designated constraints aren't met.
type SortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SortValidationError) ErrorName() string { return "SortValidationError" }

// Error satisfies the builtin error interface
func (e SortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SortValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Avatar

	// no validation rules for Email

	// no validation rules for Role

	// no validation rules for RoleId

	// no validation rules for Realname

	// no validation rules for Gender

	// no validation rules for Location

	// no validation rules for LoginCount

	// no validation rules for Mobile

	// no validation rules for State

	// no validation rules for Summary

	// no validation rules for WechatAccount

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.PasswordHash != nil {
		// no validation rules for PasswordHash
	}

	if m.SelectEnterpriseId != nil {
		// no validation rules for SelectEnterpriseId
	}

	if m.SelectEnterprise != nil {

		if all {
			switch v := interface{}(m.GetSelectEnterprise()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserValidationError{
						field:  "SelectEnterprise",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserValidationError{
						field:  "SelectEnterprise",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSelectEnterprise()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserValidationError{
					field:  "SelectEnterprise",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on Member with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Member) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Member with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MemberMultiError, or nil if none found.
func (m *Member) ValidateAll() error {
	return m.validate(true)
}

func (m *Member) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for EnterpriseId

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MemberValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MemberValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MemberValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MemberValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MemberValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MemberValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.User != nil {

		if all {
			switch v := interface{}(m.GetUser()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MemberValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MemberValidationError{
						field:  "User",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MemberValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Enterprise != nil {

		if all {
			switch v := interface{}(m.GetEnterprise()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MemberValidationError{
						field:  "Enterprise",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MemberValidationError{
						field:  "Enterprise",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEnterprise()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MemberValidationError{
					field:  "Enterprise",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MemberMultiError(errors)
	}

	return nil
}

// MemberMultiError is an error wrapping multiple validation errors returned by
// Member.ValidateAll() if the designated constraints aren't met.
type MemberMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemberMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemberMultiError) AllErrors() []error { return m }

// MemberValidationError is the validation error returned by Member.Validate if
// the designated constraints aren't met.
type MemberValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemberValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemberValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemberValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemberValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemberValidationError) ErrorName() string { return "MemberValidationError" }

// Error satisfies the builtin error interface
func (e MemberValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMember.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemberValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemberValidationError{}

// Validate checks the field values on Enterprise with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Enterprise) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Enterprise with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EnterpriseMultiError, or
// nil if none found.
func (m *Enterprise) ValidateAll() error {
	return m.validate(true)
}

func (m *Enterprise) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Logo

	// no validation rules for Website

	// no validation rules for Grade

	// no validation rules for Depth

	// no validation rules for OnwerId

	// no validation rules for ChildrenCount

	for idx, item := range m.GetMembers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  fmt.Sprintf("Members[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  fmt.Sprintf("Members[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnterpriseValidationError{
					field:  fmt.Sprintf("Members[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnterpriseValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnterpriseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnterpriseValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnterpriseValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnterpriseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnterpriseValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnterpriseValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Owner != nil {

		if all {
			switch v := interface{}(m.GetOwner()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  "Owner",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  "Owner",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetOwner()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnterpriseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Parent != nil {

		if all {
			switch v := interface{}(m.GetParent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  "Parent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  "Parent",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetParent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnterpriseValidationError{
					field:  "Parent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ParentId != nil {
		// no validation rules for ParentId
	}

	if len(errors) > 0 {
		return EnterpriseMultiError(errors)
	}

	return nil
}

// EnterpriseMultiError is an error wrapping multiple validation errors
// returned by Enterprise.ValidateAll() if the designated constraints aren't met.
type EnterpriseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnterpriseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnterpriseMultiError) AllErrors() []error { return m }

// EnterpriseValidationError is the validation error returned by
// Enterprise.Validate if the designated constraints aren't met.
type EnterpriseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnterpriseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnterpriseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnterpriseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnterpriseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnterpriseValidationError) ErrorName() string { return "EnterpriseValidationError" }

// Error satisfies the builtin error interface
func (e EnterpriseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnterprise.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnterpriseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnterpriseValidationError{}

// Validate checks the field values on UserInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserInfoMultiError, or nil
// if none found.
func (m *UserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ExternalUserid

	// no validation rules for Username

	// no validation rules for Type

	// no validation rules for Avatar

	// no validation rules for Gender

	if len(errors) > 0 {
		return UserInfoMultiError(errors)
	}

	return nil
}

// UserInfoMultiError is an error wrapping multiple validation errors returned
// by UserInfo.ValidateAll() if the designated constraints aren't met.
type UserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserInfoMultiError) AllErrors() []error { return m }

// UserInfoValidationError is the validation error returned by
// UserInfo.Validate if the designated constraints aren't met.
type UserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoValidationError) ErrorName() string { return "UserInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoValidationError{}
