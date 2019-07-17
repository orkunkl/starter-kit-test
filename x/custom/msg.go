package custom

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
)

func init() {
	migration.MustRegister(1, &CreateCustomStateMsg{}, migration.NoModification)
}

const (
	pathCreateCustomStateMsg = "custom/create_custom_state"
)

var _ weave.Msg = (*CreateCustomStateMsg)(nil)

func (CreateCustomStateMsg) Path() string {
	return "custom/create_custom_state"
}

func (m CreateCustomStateMsg) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "trader", m.CustomAddress.Validate())
	// TODO add custom validation for your state fields
	return errs 
}
