package custom

import (
	"strings"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
)

func init() {
	migration.MustRegister(1, &CreateCustomStateMsg{}, migration.NoModification)
	migration.MustRegister(1, &CreateCustomStateIndexedMsg{}, migration.NoModification)
}

var _ weave.Msg = (*CreateCustomStateIndexedMsg)(nil)

func (CreateCustomStateIndexedMsg) Path() string {
	return "custom/create_custom_indexed_state"
}

func (m CreateCustomStateIndexedMsg) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "CustomString", customStringValidation(m.CustomString))
	if m.CustomByte == nil {
		errs = errors.Append(errs, errors.Field("CustomByte", errors.ErrEmpty, "missing custom byte"))
	}
	if m.InnerStateEnum != InnerStateEnum_CaseOne && m.InnerStateEnum != InnerStateEnum_CaseTwo {
		errs = errors.AppendField(errs, "InnerStateEnum", errors.ErrState)
	}
	return errs
}

var _ weave.Msg = (*CreateCustomStateMsg)(nil)

func (CreateCustomStateMsg) Path() string {
	return "custom/create_custom_state"
}

func (m CreateCustomStateMsg) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "CustomAddress", m.CustomAddress.Validate())
	if m.InnerState == nil {
		errs = errors.AppendField(errs, "InnerState", errors.ErrEmpty)
	}
	return errs
}

// validID returns an error if this is not an 8-byte ID
// as expected for orm.IDGenBucket
func validID(id []byte) error {
	if len(id) == 0 {
		return errors.Wrap(errors.ErrEmpty, "id missing")
	}
	if len(id) != 8 {
		return errors.Wrap(errors.ErrInput, "id is invalid length (expect 8 bytes)")
	}
	return nil
}

func customStringValidation(str string) error {
	if len(str) == 0 {
		return errors.Wrap(errors.ErrEmpty, "string missing")
	}
	if !strings.HasPrefix(str, "cstm") {
		return errors.Wrap(errors.ErrInput, "string does not have cstm prefix")
	}
	return nil
}
