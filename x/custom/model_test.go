package custom

import (
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/weavetest/assert"

	"github.com/iov-one/tutorial/morm"
)

func TestValidateCustomStateIndexed(t *testing.T) {
	now := weave.AsUnixTime(time.Now())

	cases := map[string]struct {
		model    morm.Model
		wantErrs map[string]*errors.Error
	}{
		"success, with id": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				InnerStateEnum: InnerStateEnum_CaseOne,
				CustomString:   "cstm_string",
				CustomByte:     []byte{0, 1},
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
		},
		"success, no id": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				CustomString:   "cstm_string",
				CustomByte:     []byte{0, 1},
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
		},
		"failure, missing metadata": {
			model: &CustomStateIndexed{
				ID:             weavetest.SequenceID(1),
				CustomString:   "cstm_string",
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       errors.ErrMetadata,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
		},
		"failure, missing custom string": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   errors.ErrEmpty,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
		},
		"failure, custom string does not begin with 'cstm'": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				CustomString:   "string",
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   errors.ErrInput,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
		},
		"failure, missing inner state enum": {
			model: &CustomStateIndexed{
				Metadata:     &weave.Metadata{Schema: 1},
				ID:           weavetest.SequenceID(1),
				CustomString: "cstm_string",
				CustomByte:   []byte{0, 1},
				DeletedAt:    now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": errors.ErrState,
				"CustomString":   nil,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
		},
		"failure, missing custom byte": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				CustomString:   "cstm_string",
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     errors.ErrEmpty,
				"DeletedAt":      nil,
			},
		},
		"failure, missing deleted at": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				CustomString:   "cstm_string",
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
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

func TestValidateCustomState(t *testing.T) {
	now := weave.AsUnixTime(time.Now())

	cases := map[string]struct {
		model    orm.Model
		wantErrs map[string]*errors.Error
	}{
		"success": {
			model: &CustomState{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
				CreatedAt:     now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      nil,
				"InnerState":    nil,
				"CustomAddress": nil,
				"CreatedAt":     nil,
			},
		},
		"failure, missing metadata": {
			model: &CustomState{
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
				CreatedAt:     now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      errors.ErrMetadata,
				"InnerState":    nil,
				"CustomAddress": nil,
				"CreatedAt":     nil,
			},
		},
		"failure, missing inner state": {
			model: &CustomState{
				Metadata:      &weave.Metadata{Schema: 1},
				CustomAddress: weavetest.NewCondition().Address(),
				CreatedAt:     now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      nil,
				"InnerState":    errors.ErrEmpty,
				"CustomAddress": nil,
				"CreatedAt":     nil,
			},
		},
		"failure, missing custom address": {
			model: &CustomState{
				Metadata:   &weave.Metadata{Schema: 1},
				InnerState: &InnerState{St1: 1, St2: 2},
				CreatedAt:  now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      nil,
				"InnerState":    nil, 
				"CustomAddress": errors.ErrEmpty,
				"CreatedAt":     nil,
			},
		},
		"failure, invalid address lenght": {
			model: &CustomState{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: []byte{0, 1},
				CreatedAt:     now,
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      nil,
				"InnerState":    nil, 
				"CustomAddress": errors.ErrInput,
				"CreatedAt":     nil,
			},
		},
		"failure, missing created at": {
			model: &CustomState{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      nil,
				"InnerState":    nil, 
				"CustomAddress": nil,
				"CreatedAt":     errors.ErrEmpty,
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
