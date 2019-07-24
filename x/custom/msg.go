package custom

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
)

func init() {
	migration.MustRegister(1, &CreateCustomStateWIDMsg{}, migration.NoModification)
}

var _ weave.Msg = (*CreateCustomStateWIDMsg)(nil)

func (CreateCustomStateWIDMsg) Path() string {
	return "custom/create_custom_state"
}

func (m CreateCustomStateWIDMsg) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "trader", m.CustomAddress.Validate())
	// TODO add custom validation for your state fields
	return errs 
}
