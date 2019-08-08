package custom

import (
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/weavetest/assert"
)

func TestValidateStateIndexed(t *testing.T) {
	now := weave.AsUnixTime(time.Now())

	cases := map[string]struct {
		model    orm.Model
		wantErrs map[string]*errors.Error
	}{
		"success, with id": {
			model: &StateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				InnerStateEnum: InnerStateEnum_CaseOne,
				Str:            "cstm_string",
				Byte:           []byte{0, 1},
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":            nil,
				"Byte":           nil,
				"DeletedAt":      nil,
			},
		},
		"success, no id": {
			model: &StateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				Str:            "cstm_string",
				Byte:           []byte{0, 1},
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":            nil,
				"Byte":           nil,
				"DeletedAt":      nil,
			},
		},
		"failure, missing metadata": {
			model: &StateIndexed{
				ID:             weavetest.SequenceID(1),
				Str:            "cstm_string",
				Byte:           []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       errors.ErrMetadata,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":            nil,
				"Byte":           nil,
				"DeletedAt":      nil,
			},
		},
		"failure, missing str": {
			model: &StateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				Byte:           []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":            errors.ErrEmpty,
				"Byte":           nil,
				"DeletedAt":      nil,
			},
		},
		"failure, str does not begin with 'cstm'": {
			model: &StateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				Str:            "string",
				Byte:           []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":            errors.ErrInput,
				"Byte":           nil,
				"DeletedAt":      nil,
			},
		},
		"failure, missing inner state enum": {
			model: &StateIndexed{
				Metadata:  &weave.Metadata{Schema: 1},
				ID:        weavetest.SequenceID(1),
				Str:       "cstm_string",
				Byte:      []byte{0, 1},
				DeletedAt: now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": errors.ErrState,
				"String":         nil,
				"Byte":           nil,
				"DeletedAt":      nil,
			},
		},
		"failure, missing custom byte": {
			model: &StateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				Str:            "cstm_string",
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":            nil,
				"Byte":           errors.ErrEmpty,
				"DeletedAt":      nil,
			},
		},
		"failure, missing deleted at": {
			model: &StateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				Str:            "cstm_string",
				Byte:           []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":            nil,
				"Byte":           nil,
				"DeletedAt":      errors.ErrEmpty,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.model.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}

func TestValidateState(t *testing.T) {
	now := weave.AsUnixTime(time.Now())

	cases := map[string]struct {
		model    orm.Model
		wantErrs map[string]*errors.Error
	}{
		"success": {
			model: &State{
				Metadata:   &weave.Metadata{Schema: 1},
				InnerState: &InnerState{St1: 1, St2: 2},
				Address:    weavetest.NewCondition().Address(),
				CreatedAt:  now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"InnerState": nil,
				"Address":    nil,
				"CreatedAt":  nil,
			},
		},
		"failure, missing metadata": {
			model: &State{
				InnerState: &InnerState{St1: 1, St2: 2},
				Address:    weavetest.NewCondition().Address(),
				CreatedAt:  now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   errors.ErrMetadata,
				"InnerState": nil,
				"Address":    nil,
				"CreatedAt":  nil,
			},
		},
		"failure, missing inner state": {
			model: &State{
				Metadata:  &weave.Metadata{Schema: 1},
				Address:   weavetest.NewCondition().Address(),
				CreatedAt: now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"InnerState": errors.ErrEmpty,
				"Address":    nil,
				"CreatedAt":  nil,
			},
		},
		"failure, missing custom address": {
			model: &State{
				Metadata:   &weave.Metadata{Schema: 1},
				InnerState: &InnerState{St1: 1, St2: 2},
				CreatedAt:  now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"InnerState": nil,
				"Address":    errors.ErrEmpty,
				"CreatedAt":  nil,
			},
		},
		"failure, invalid address lenght": {
			model: &State{
				Metadata:   &weave.Metadata{Schema: 1},
				InnerState: &InnerState{St1: 1, St2: 2},
				Address:    []byte{0, 1},
				CreatedAt:  now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"InnerState": nil,
				"Address":    errors.ErrInput,
				"CreatedAt":  nil,
			},
		},
		"failure, missing created at": {
			model: &State{
				Metadata:   &weave.Metadata{Schema: 1},
				InnerState: &InnerState{St1: 1, St2: 2},
				Address:    weavetest.NewCondition().Address(),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"InnerState": nil,
				"Address":    nil,
				"CreatedAt":  errors.ErrEmpty,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.model.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}
