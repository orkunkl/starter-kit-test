package custom

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
)

type TimedStateBucket struct {
	orm.ModelBucket
}

func NewTimedStateBucket() *TimedStateBucket {
	b := orm.NewModelBucket("timedstate", &TimedState{})
	return &TimedStateBucket{
		ModelBucket: migration.NewModelBucket(packageName, b),
	}
}

// GetTimedState loads the TimedState for the given id. If it does not exist then ErrNotFound is returned.
func (b *TimedStateBucket) GetTimedState(db weave.KVStore, id []byte) (*TimedState, error) {
	var timedState TimedState
	err := b.One(db, id, &timedState)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load timed state")
	}
	return &timedState, nil
}

type StateBucket struct {
	orm.ModelBucket
}

func NewStateBucket() *StateBucket {
	b := orm.NewModelBucket("state", &State{}, orm.WithIDSequence(stateSeq))
	return &StateBucket{
		ModelBucket: migration.NewModelBucket(packageName, b),
	}
}

var stateSeq = orm.NewSequence("state", "id")
