package custom

import (
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/migration"
)

const (
	packageName = "custom"
)
type StateIndexedBucket struct {
	orm.ModelBucket
}

func newCustomStateIndexedBucket() *StateIndexedBucket {
	b := orm.NewModelBucket("stateIndexed", &StateIndexed{})
	return &StateIndexedBucket{
		ModelBucket: migration.NewModelBucket("mStateIndexed", b),
	}
}

type StateBucket struct {
	orm.ModelBucket
}

func newCustomStateBucket() *StateBucket {
	b := orm.NewModelBucket("state", &State{})
	return &StateBucket{
		ModelBucket: migration.NewModelBucket("mState", b),
	}
}
