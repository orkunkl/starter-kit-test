package custom

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
)

type TimedStateBucket struct {
	orm.IDGenBucket
}

func NewTimedStateBucket() *TimedStateBucket {
	b := migration.NewBucket(packageName, "stateind", orm.NewSimpleObj(nil, &TimedState{}))
	return &TimedStateBucket{
		IDGenBucket: orm.WithSeqIDGenerator(b, "id"),
	}
}

// GetTimedState loads the TimedState for the given id. If it does not exist then ErrNotFound is returned.
func (b *TimedStateBucket) GetTimedState(db weave.KVStore, id []byte) (*TimedState, error) {
	obj, err := b.Get(db, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load indexed state")
	}
	return asTimedState(obj)
}

func asTimedState(obj orm.Object) (*TimedState, error) {
	if obj == nil || obj.Value() == nil {
		return nil, errors.Wrap(errors.ErrNotFound, "unknown id")
	}
	rev, ok := obj.Value().(*TimedState)
	if !ok {
		return nil, errors.Wrapf(errors.ErrModel, "invalid type: %T", obj.Value())
	}
	return rev, nil
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