package custom

import (
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/weavetest/assert"
)

func TestValidateCreateCustomStateIndexedMsg(t *testing.T) {
	cases := map[string]struct {
		msg      weave.Msg
		wantErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				CustomString:   "cstm:str",
				CustomByte:     []byte{0, 1},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
			},
		},
		"missing metadata": {
			msg: &CreateCustomStateIndexedMsg{
				InnerStateEnum: InnerStateEnum_CaseTwo,
				CustomString:   "cstm:str",
				CustomByte:     []byte{0, 1},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       errors.ErrMetadata,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
			},
		},
		"missing inner state enum": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:     &weave.Metadata{Schema: 1},
				CustomString: "cstm:str",
				CustomByte:   []byte{0, 1},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"InnerStateEnum": errors.ErrState,
				"CustomString":   nil,
				"CustomByte":     nil,
			},
		},
		"missing custom string": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseTwo,
				CustomByte:     []byte{0, 1},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"InnerStateEnum": nil,
				"CustomString":   errors.ErrEmpty,
				"CustomByte":     nil,
			},
		},
		"custom string does not begin with cstm": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseTwo,
				CustomString:   "str",
				CustomByte:     []byte{0, 1},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"InnerStateEnum": nil,
				"CustomString":   errors.ErrInput,
				"CustomByte":     nil,
			},
		},
		"missing custom byte": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata:       &weave.Metadata{Schema: 1},
				InnerStateEnum: InnerStateEnum_CaseOne,
				CustomString:   "cstm:str",
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     errors.ErrEmpty,
			},
		}}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.msg.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}

func TestValidateCreateCustomStateMsg(t *testing.T) {
	cases := map[string]struct {
		msg      weave.Msg
		wantErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CreateCustomStateMsg{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      nil,
				"InnerState":    nil,
				"CustomAddress": nil,
			},
		},
		"missing metadata": {
			msg: &CreateCustomStateMsg{
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      errors.ErrMetadata,
				"InnerState":    nil,
				"CustomAddress": nil,
			},
		},
		"missing inner state": {
			msg: &CreateCustomStateMsg{
				Metadata:      &weave.Metadata{Schema: 1},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      nil,
				"InnerState":    errors.ErrEmpty,
				"CustomAddress": nil,
			},
		},
		"bad address": {
			msg: &CreateCustomStateMsg{
				Metadata:      &weave.Metadata{Schema: 1},
				InnerState:    &InnerState{St1: 1, St2: 2},
				CustomAddress: []byte{0, 1},
			},
			wantErrs: map[string]*errors.Error{
				"Metadata":      nil,
				"InnerState":    nil,
				"CustomAddress": errors.ErrInput,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := tc.msg.Validate()
			for field, wantErr := range tc.wantErrs {
				assert.FieldError(t, err, field, wantErr)
			}
		})
	}
}
