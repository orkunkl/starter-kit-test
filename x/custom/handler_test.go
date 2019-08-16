package custom

import (
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/weavetest/assert"
)

func TestCreateTimedState(t *testing.T) {
	meta := &weave.Metadata{Schema: 1}

	cases := map[string]struct {
		msg             weave.Msg
		expected        *TimedState
		wantCheckErrs   map[string]*errors.Error
		wantDeliverErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CreateTimedStateMsg{
				Metadata:       meta,
				InnerStateEnum: InnerStateEnum_CaseOne,
				Str:            "cstm_str",
				Byte:           []byte{0, 1},
			},
			expected: &TimedState{
				Metadata:       meta,
				InnerStateEnum: InnerStateEnum_CaseOne,
				Str:            "cstm_str",
				Byte:           []byte{0, 1},
			},
			wantCheckErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"InnerStateEnum": nil,
				"Str":            nil,
				"Byte":           nil,
				"DeletedAt":      nil,
			},
			wantDeliverErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"InnerStateEnum": nil,
				"Str":            nil,
				"Byte":           nil,
				"DeletedAt":      nil,
			},
		},
		"failure, empty message": {
			msg: &CreateTimedStateMsg{},
			wantCheckErrs: map[string]*errors.Error{
				"Metadata":       errors.ErrMetadata,
				"InnerStateEnum": errors.ErrState,
				"Str":            errors.ErrEmpty,
				"Byte":           errors.ErrEmpty,
				"DeletedAt":      nil,
			},
			wantDeliverErrs: map[string]*errors.Error{
				"Metadata":       errors.ErrMetadata,
				"InnerStateEnum": errors.ErrState,
				"Str":            errors.ErrEmpty,
				"Byte":           errors.ErrEmpty,
				"DeletedAt":      nil,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			auth := &weavetest.Auth{}

			h := NewTimedStateHandler(auth)
			kv := store.MemStore()
			bucket := NewTimedStateBucket()
			migration.MustInitPkg(kv, packageName)

			tx := &weavetest.Tx{Msg: tc.msg}

			if _, err := h.Check(nil, kv, tx); err != nil {
				for field, wantErr := range tc.wantCheckErrs {
					assert.FieldError(t, err, field, wantErr)
				}
			}

			res, err := h.Deliver(nil, kv, tx)
			for field, wantErr := range tc.wantDeliverErrs {
				assert.FieldError(t, err, field, wantErr)
			}

			if tc.expected != nil {
				stored, err := bucket.GetTimedState(kv, res.Data)

				assert.Nil(t, err)
				assert.Equal(t, tc.expected, stored)
			}
		})
	}
}

func TestCreateState(t *testing.T) {
	meta := &weave.Metadata{Schema: 1}
	now := weave.AsUnixTime(time.Now())
	address := weavetest.NewCondition().Address()

	cases := map[string]struct {
		msg             weave.Msg
		expected        *State
		wantCheckErrs   map[string]*errors.Error
		wantDeliverErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CreateStateMsg{
				Metadata:   meta,
				InnerState: &InnerState{St1: 1, St2: 2},
				Address:    address,
			},
			expected: &State{
				Metadata:   meta,
				InnerState: &InnerState{St1: 1, St2: 2},
				Address:    address,
				CreatedAt:  now,
			},
			wantCheckErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"InnerState": nil,
				"Address":    nil,
				"CreatedAt":  nil,
			},
			wantDeliverErrs: map[string]*errors.Error{
				"Metadata":   nil,
				"InnerState": nil,
				"Address":    nil,
				"CreatedAt":  nil,
			},
		},
		"failure empty message": {
			msg: &CreateStateMsg{},
			wantCheckErrs: map[string]*errors.Error{
				"Metadata":   errors.ErrMetadata,
				"InnerState": errors.ErrEmpty,
				"Address":    errors.ErrEmpty,
				"CreatedAt":  nil,
			},
			wantDeliverErrs: map[string]*errors.Error{
				"Metadata":   errors.ErrMetadata,
				"InnerState": errors.ErrEmpty,
				"Address":    errors.ErrEmpty,
				"CreatedAt":  nil,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			auth := &weavetest.Auth{}

			h := NewStateHandler(auth)
			kv := store.MemStore()
			bucket := NewStateBucket()
			migration.MustInitPkg(kv, packageName)

			tx := &weavetest.Tx{Msg: tc.msg}

			if _, err := h.Check(nil, kv, tx); err != nil {
				for field, wantErr := range tc.wantCheckErrs {
					assert.FieldError(t, err, field, wantErr)
				}
			}

			res, err := h.Deliver(nil, kv, tx)
			for field, wantErr := range tc.wantDeliverErrs {
				assert.FieldError(t, err, field, wantErr)
			}

			if tc.expected != nil {
				err := bucket.Has(kv, res.Data)
				assert.Nil(t, err)
			}
		})
	}
}