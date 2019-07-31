package custom

import (
	"github.com/iov-one/tutorial/morm"
	"github.com/iov-one/weave/orm"
)

type CustomStateIndexedBucket struct {
	morm.ModelBucket
}

func newCustomStateIndexedBucket() *CustomStateIndexedBucket {
	b := morm.NewModelBucket("customStateIndexed", &CustomStateIndexed{})
	return &CustomStateIndexedBucket{
		ModelBucket: b,
	}
}

type CustomStateBucket struct {
	orm.ModelBucket
}

func newCustomStateBucket() *CustomStateBucket {
	b := orm.NewModelBucket("customState", &CustomState{})
	return &CustomStateBucket{
		ModelBucket: b,
	}
}
