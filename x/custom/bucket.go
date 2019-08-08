package custom

import (
	"github.com/iov-one/weave/orm"
)

type StateIndexedBucket struct {
	orm.ModelBucket
}

func newCustomStateIndexedBucket() *StateIndexedBucket {
	b := orm.NewModelBucket("stateIndexed", &StateIndexed{})
	return &StateIndexedBucket{
		ModelBucket: b,
	}
}

type StateBucket struct {
	orm.ModelBucket
}

func newCustomStateBucket() *StateBucket {
	b := orm.NewModelBucket("state", &State{})
	return &StateBucket{
		ModelBucket: b,
	}
}
