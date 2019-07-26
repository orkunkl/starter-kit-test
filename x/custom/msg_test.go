package custom

import (
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/weavetest"
)

func TestValidateCreateCustomStateIndexedMsg(t *testing.T) {
	cases := map[string]struct {
		msg     weave.Msg
		wantErr *errors.Error
	}{
		"success": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				CustomInt:      1,
				CustomString:   "cstm:str",
				CustomByte:     []byte{0, 1},
			},
			wantErr: nil,
		},
		"missing metadata": {
			msg: &CreateCustomStateIndexedMsg{
				InnerStateEnum: InnerStateEnum_CaseTwo,
				CustomInt:      1,
				CustomString:   "cstm:str",
				CustomByte:     []byte{0, 1},
			},
			wantErr: errors.ErrMetadata,
		},
		"missing inner state enum": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:     &weave.Metadata{Schema: 1},
				CustomInt:    1,
				CustomString: "cstm:str",
				CustomByte:   []byte{0, 1},
			},
			wantErr: errors.ErrState,
		},
		"missing custom string": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseTwo,
				CustomInt:      1,
				CustomByte:     []byte{0, 1},
			},
			wantErr: errors.ErrEmpty,
		},
		"custom string does not begin with cstm": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseTwo,
				CustomInt:      1,
				CustomString:   "str",
				CustomByte:     []byte{0, 1},
			},
			wantErr: errors.ErrInput,
		},
		"missing custom byte": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				CustomInt:      1,
				CustomString:   "cstm:str",
			},
			wantErr: errors.ErrEmpty,
		}}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if err := tc.msg.Validate(); !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
		})
	}
}

func TestValidateCreateCustomStateMsg(t *testing.T) {
	cases := map[string]struct {
		msg     weave.Msg
		wantErr *errors.Error
	}{
		"success": {
			msg: &CreateCustomStateMsg{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErr: nil,
		},
		"missing metadata": {
			msg: &CreateCustomStateMsg{
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErr: errors.ErrMetadata,
		},
		"missing inner state": {
			msg: &CreateCustomStateMsg{
				Metadata:      &weave.Metadata{Schema: 1},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErr: errors.ErrEmpty,
		},
		"bad address": {
			msg: &CreateCustomStateMsg{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: []byte{0, 1},
			},
			wantErr: errors.ErrInput,
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if err := tc.msg.Validate(); !tc.wantErr.Is(err) {
				t.Fatalf("unexpected error: %+v", err)
			}
		})
	}
}
