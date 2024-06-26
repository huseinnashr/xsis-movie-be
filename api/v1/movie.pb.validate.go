// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/movie.proto

package v1

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

// Validate checks the field values on Movie with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Movie) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Movie with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MovieMultiError, or nil if none found.
func (m *Movie) ValidateAll() error {
	return m.validate(true)
}

func (m *Movie) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for Rating

	// no validation rules for Image

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return MovieMultiError(errors)
	}

	return nil
}

// MovieMultiError is an error wrapping multiple validation errors returned by
// Movie.ValidateAll() if the designated constraints aren't met.
type MovieMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MovieMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MovieMultiError) AllErrors() []error { return m }

// MovieValidationError is the validation error returned by Movie.Validate if
// the designated constraints aren't met.
type MovieValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MovieValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MovieValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MovieValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MovieValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MovieValidationError) ErrorName() string { return "MovieValidationError" }

// Error satisfies the builtin error interface
func (e MovieValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMovie.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MovieValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MovieValidationError{}

// Validate checks the field values on ListMovieRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListMovieRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMovieRequestMultiError, or nil if none found.
func (m *ListMovieRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMovieRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetPageSize(); val <= 0 || val > 10 {
		err := ListMovieRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 10]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PageToken

	if len(errors) > 0 {
		return ListMovieRequestMultiError(errors)
	}

	return nil
}

// ListMovieRequestMultiError is an error wrapping multiple validation errors
// returned by ListMovieRequest.ValidateAll() if the designated constraints
// aren't met.
type ListMovieRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMovieRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMovieRequestMultiError) AllErrors() []error { return m }

// ListMovieRequestValidationError is the validation error returned by
// ListMovieRequest.Validate if the designated constraints aren't met.
type ListMovieRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMovieRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMovieRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMovieRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMovieRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMovieRequestValidationError) ErrorName() string { return "ListMovieRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListMovieRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMovieRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMovieRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMovieRequestValidationError{}

// Validate checks the field values on ListMovieResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListMovieResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMovieResponseMultiError, or nil if none found.
func (m *ListMovieResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMovieResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMovies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListMovieResponseValidationError{
						field:  fmt.Sprintf("Movies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListMovieResponseValidationError{
						field:  fmt.Sprintf("Movies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListMovieResponseValidationError{
					field:  fmt.Sprintf("Movies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListMovieResponseMultiError(errors)
	}

	return nil
}

// ListMovieResponseMultiError is an error wrapping multiple validation errors
// returned by ListMovieResponse.ValidateAll() if the designated constraints
// aren't met.
type ListMovieResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMovieResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMovieResponseMultiError) AllErrors() []error { return m }

// ListMovieResponseValidationError is the validation error returned by
// ListMovieResponse.Validate if the designated constraints aren't met.
type ListMovieResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMovieResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMovieResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMovieResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMovieResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMovieResponseValidationError) ErrorName() string {
	return "ListMovieResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListMovieResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMovieResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMovieResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMovieResponseValidationError{}

// Validate checks the field values on CreateMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateMovieRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMovieRequestMultiError, or nil if none found.
func (m *CreateMovieRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMovieRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Description

	if val := m.GetRating(); val < 0 || val > 10 {
		err := CreateMovieRequestValidationError{
			field:  "Rating",
			reason: "value must be inside range [0, 10]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Image

	if len(errors) > 0 {
		return CreateMovieRequestMultiError(errors)
	}

	return nil
}

// CreateMovieRequestMultiError is an error wrapping multiple validation errors
// returned by CreateMovieRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateMovieRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMovieRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMovieRequestMultiError) AllErrors() []error { return m }

// CreateMovieRequestValidationError is the validation error returned by
// CreateMovieRequest.Validate if the designated constraints aren't met.
type CreateMovieRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMovieRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMovieRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMovieRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMovieRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMovieRequestValidationError) ErrorName() string {
	return "CreateMovieRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMovieRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMovieRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMovieRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMovieRequestValidationError{}

// Validate checks the field values on CreateMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateMovieResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMovieResponseMultiError, or nil if none found.
func (m *CreateMovieResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMovieResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateMovieResponseMultiError(errors)
	}

	return nil
}

// CreateMovieResponseMultiError is an error wrapping multiple validation
// errors returned by CreateMovieResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateMovieResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMovieResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMovieResponseMultiError) AllErrors() []error { return m }

// CreateMovieResponseValidationError is the validation error returned by
// CreateMovieResponse.Validate if the designated constraints aren't met.
type CreateMovieResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMovieResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMovieResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMovieResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMovieResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMovieResponseValidationError) ErrorName() string {
	return "CreateMovieResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMovieResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMovieResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMovieResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMovieResponseValidationError{}

// Validate checks the field values on GetMovieRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetMovieRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMovieRequestMultiError, or nil if none found.
func (m *GetMovieRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMovieRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetMovieRequestMultiError(errors)
	}

	return nil
}

// GetMovieRequestMultiError is an error wrapping multiple validation errors
// returned by GetMovieRequest.ValidateAll() if the designated constraints
// aren't met.
type GetMovieRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMovieRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMovieRequestMultiError) AllErrors() []error { return m }

// GetMovieRequestValidationError is the validation error returned by
// GetMovieRequest.Validate if the designated constraints aren't met.
type GetMovieRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMovieRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMovieRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMovieRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMovieRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMovieRequestValidationError) ErrorName() string { return "GetMovieRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetMovieRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMovieRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMovieRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMovieRequestValidationError{}

// Validate checks the field values on UpdateMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateMovieRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateMovieRequestMultiError, or nil if none found.
func (m *UpdateMovieRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateMovieRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Title != nil {
		// no validation rules for Title
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if m.Rating != nil {

		if val := m.GetRating(); val < 0 || val > 10 {
			err := UpdateMovieRequestValidationError{
				field:  "Rating",
				reason: "value must be inside range [0, 10]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.Image != nil {
		// no validation rules for Image
	}

	if len(errors) > 0 {
		return UpdateMovieRequestMultiError(errors)
	}

	return nil
}

// UpdateMovieRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateMovieRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateMovieRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateMovieRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateMovieRequestMultiError) AllErrors() []error { return m }

// UpdateMovieRequestValidationError is the validation error returned by
// UpdateMovieRequest.Validate if the designated constraints aren't met.
type UpdateMovieRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMovieRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMovieRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMovieRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMovieRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMovieRequestValidationError) ErrorName() string {
	return "UpdateMovieRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMovieRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMovieRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMovieRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMovieRequestValidationError{}

// Validate checks the field values on UpdateMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateMovieResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateMovieResponseMultiError, or nil if none found.
func (m *UpdateMovieResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateMovieResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return UpdateMovieResponseMultiError(errors)
	}

	return nil
}

// UpdateMovieResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateMovieResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateMovieResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateMovieResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateMovieResponseMultiError) AllErrors() []error { return m }

// UpdateMovieResponseValidationError is the validation error returned by
// UpdateMovieResponse.Validate if the designated constraints aren't met.
type UpdateMovieResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMovieResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMovieResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMovieResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMovieResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMovieResponseValidationError) ErrorName() string {
	return "UpdateMovieResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMovieResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMovieResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMovieResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMovieResponseValidationError{}

// Validate checks the field values on DeleteMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteMovieRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteMovieRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteMovieRequestMultiError, or nil if none found.
func (m *DeleteMovieRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteMovieRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteMovieRequestMultiError(errors)
	}

	return nil
}

// DeleteMovieRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteMovieRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteMovieRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteMovieRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteMovieRequestMultiError) AllErrors() []error { return m }

// DeleteMovieRequestValidationError is the validation error returned by
// DeleteMovieRequest.Validate if the designated constraints aren't met.
type DeleteMovieRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMovieRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMovieRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMovieRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMovieRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMovieRequestValidationError) ErrorName() string {
	return "DeleteMovieRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMovieRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMovieRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMovieRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMovieRequestValidationError{}

// Validate checks the field values on DeleteMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteMovieResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteMovieResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteMovieResponseMultiError, or nil if none found.
func (m *DeleteMovieResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteMovieResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return DeleteMovieResponseMultiError(errors)
	}

	return nil
}

// DeleteMovieResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteMovieResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteMovieResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteMovieResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteMovieResponseMultiError) AllErrors() []error { return m }

// DeleteMovieResponseValidationError is the validation error returned by
// DeleteMovieResponse.Validate if the designated constraints aren't met.
type DeleteMovieResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMovieResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMovieResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMovieResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMovieResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMovieResponseValidationError) ErrorName() string {
	return "DeleteMovieResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMovieResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMovieResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMovieResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMovieResponseValidationError{}
