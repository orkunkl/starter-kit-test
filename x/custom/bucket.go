package custom

import (
	"github.com/iov-one/tutorial/morm"
	"github.com/iov-one/weave/orm"
)

type StateIndexedBucket struct {
	morm.ModelBucket
}

func newCustomStateIndexedBucket() *StateIndexedBucket {
	b := morm.NewModelBucket("stateIndexed", &StateIndexed{})
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
