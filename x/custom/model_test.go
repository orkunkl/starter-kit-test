package custom

import (
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/weavetest"

	"github.com/iov-one/tutorial/morm"
)

func TestValidateCustomStateIndexed(t *testing.T) {
	now := weave.AsUnixTime(time.Now())

	cases := map[string]struct {
		model   morm.Model
		wantErr *errors.Error
	}{
		"success, with id": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				CustomString:   "cstm_string",
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErr: nil,
		},
		"success, no id": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				CustomString:   "cstm_string",
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErr: nil,
		},
		"failure, missing metadata": {
			model: &CustomStateIndexed{
				ID:             weavetest.SequenceID(1),
				CustomString:   "cstm_string",
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErr: errors.ErrMetadata,
		},
		"failure, missing custom string": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErr: errors.ErrEmpty,
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
			wantErr: errors.ErrInput,
		},
		"failure, missing inner state enum": {
			model: &CustomStateIndexed{
				Metadata:     &weave.Metadata{Schema: 1},
				ID:           weavetest.SequenceID(1),
				CustomString: "cstm_string",
				CustomByte:   []byte{0, 1},
				DeletedAt:    now,
			},
			wantErr: errors.ErrState,
		},
		"failure, missing custom byte": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				CustomString:   "cstm_string",
				InnerStateEnum: InnerStateEnum_CaseOne,
				DeletedAt:      now,
			},
			wantErr: errors.ErrEmpty,
		},
		"failure, missing deleted at": {
			model: &CustomStateIndexed{
				Metadata:       &weave.Metadata{Schema: 1},
				ID:             weavetest.SequenceID(1),
				CustomString:   "cstm_string",
				CustomByte:     []byte{0, 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
			},
			wantErr: errors.ErrEmpty,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if err := tc.model.Validate(); !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
		})
	}
}

func TestValidateCustomState(t *testing.T) {
	now := weave.AsUnixTime(time.Now())

	cases := map[string]struct {
		model   orm.Model
		wantErr *errors.Error
	}{
		"success": {
			model: &CustomState{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
				CreatedAt:     now,
			},
			wantErr: nil,
		},
		"failure, missing metadata": {
			model: &CustomState{
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
				CreatedAt:     now,
			},
			wantErr: errors.ErrMetadata,
		},
		"failure, missing inner state": {
			model: &CustomState{
				Metadata:      &weave.Metadata{Schema: 1},
				CustomAddress: weavetest.NewCondition().Address(),
				CreatedAt:     now,
			},
			wantErr: errors.ErrEmpty,
		},
		"failure, missing custom address": {
			model: &CustomState{
				Metadata:   &weave.Metadata{Schema: 1},
				InnerState: &InnerState{St1: 1, St2: 2},
				CreatedAt:  now,
			},
			wantErr: errors.ErrEmpty,
		},
		"failure, invalid address lenght": {
			model: &CustomState{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: []byte{0, 1},
				CreatedAt:     now,
			},
			wantErr: errors.ErrInput,
		},
		"failure, missing created at": {
			model: &CustomState{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErr: errors.ErrEmpty,
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if err := tc.model.Validate(); !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
		})
	}
}
