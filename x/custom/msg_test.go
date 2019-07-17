package custom

import (
	"testing"
	
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/weavetest"
)

func TestValidateCreateCustomStateMsg(t *testing.T) {
	cases := map[string]struct {
		msg weave.Msg
		wantErr *errors.Error
	}{
		"success": {
			msg: &CreateCustomStateMsg {
				Metadata: &weave.Metadata{Schema: 1},
				CustomInt: 1,
				CustomString: "str",
				CustomByte: []byte{0, 0, 0, 0},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErr: nil,
		},
		"missing metadata": {
			msg: &CreateCustomStateMsg {
				CustomInt: 1,
				CustomString: "str",
				CustomByte: []byte{0, 0, 0, 1},
				CustomAddress: weavetest.NewCondition().Address(),
			},
			wantErr: errors.ErrMetadata,
		},
		"bad address": {
			msg: &CreateCustomStateMsg {
				CustomInt: 1,
				CustomString: "str",
				CustomByte: []byte{0, 0, 0, 1},
				CustomAddress: []byte{0, 0, 0, 2},
			},
			wantErr: errors.ErrMetadata,
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
