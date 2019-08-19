package custom

import (
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/x"
)

const (
	packageName        = "custom"
	newStateCost int64 = 100
)

// RegisterQuery registers buckets for querying.
func RegisterQuery(qr weave.QueryRouter) {
	NewTimedStateBucket().Register("timedStates", qr)
	NewStateBucket().Register("states", qr)
}

// RegisterRoutes registers handlers for message processing.
func RegisterRoutes(r weave.Registry, auth x.Authenticator) {
	r = migration.SchemaMigratingRegistry(packageName, r)

	r.Handle(&CreateTimedStateMsg{}, NewTimedStateHandler(auth))
	r.Handle(&CreateStateMsg{}, NewStateHandler(auth))
}

// ------------------- TimedState HANDLERS -------------------

// TimedStateHandler will handle creating custom indexed state buckets
type TimedStateHandler struct {
	auth x.Authenticator
	b    *TimedStateBucket
}

var _ weave.Handler = TimedStateHandler{}

// NewTimedStateHandler creates a handler
func NewTimedStateHandler(auth x.Authenticator) weave.Handler {
	return TimedStateHandler{
		auth: auth,
		b:    NewTimedStateBucket(),
	}
}

// validate does all common pre-processing between Check and Deliver
func (h TimedStateHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*CreateTimedStateMsg, error) {
	var msg CreateTimedStateMsg

	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, errors.Wrap(err, "load msg")
	}

	if msg.DeleteAt != 0 && weave.InThePast(ctx, msg.DeleteAt.Time()) == true {
		return nil, errors.AppendField(nil, "DeleteAt", errors.ErrInput)
	}

	return &msg, nil
}

// Check just verifies it is properly formed and returns
// the cost of executing it.
func (h TimedStateHandler) Check(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	return &weave.CheckResult{GasAllocated: newStateCost}, nil
}

// Deliver creates an custom state and saves if all preconditions are met
func (h TimedStateHandler) Deliver(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, err := h.validate(ctx, store, tx)

	if err != nil {
		return nil, err
	}

	timedState := &TimedState{
		Metadata:       &weave.Metadata{},
		InnerStateEnum: msg.InnerStateEnum,
		Str:            msg.Str,
		Byte:           msg.Byte,
		DeleteAt:       msg.DeleteAt,
	}

	key, err := h.b.Put(store, nil, timedState)
	if err != nil {
		return nil, errors.Wrap(err, "cannot store indexed state")
	}

	return &weave.DeliverResult{Data: key}, nil
}

// DeleteTimedStateHandler will handle deleting timed state
type DeleteTimedStateHandler struct {
	auth x.Authenticator
	b    *TimedStateBucket
}

var _ weave.Handler = DeleteTimedStateHandler{}

// NewDeleteTimedStateHandler creates a handler
func NewDeleteTimedStateHandler(auth x.Authenticator) weave.Handler {
	return DeleteTimedStateHandler{
		auth: auth,
		b:    NewTimedStateBucket(),
	}
}

// validate does all common pre-processing between Check and Deliver
func (h DeleteTimedStateHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*DeleteTimedStateMsg, error) {
	var msg DeleteTimedStateMsg

	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, errors.Wrap(err, "load msg")
	}

	return &msg, nil
}

// Check just verifies it is properly formed and returns
// the cost of executing it.
func (h DeleteTimedStateHandler) Check(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	return &weave.CheckResult{GasAllocated: newStateCost}, nil
}

// Deliver delete state if all preconditions are met
func (h DeleteTimedStateHandler) Deliver(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, err := h.validate(ctx, store, tx)

	if err != nil {
		return nil, err
	}
	err = h.b.Delete(store, msg.TimedStateID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot store indexed state")
	}

	return &weave.DeliverResult{}, nil
}

// ------------------- State HANDLERS -------------------

// StateHandler will handle creating custom state buckets
type StateHandler struct {
	auth x.Authenticator
	b    *StateBucket
}

var _ weave.Handler = StateHandler{}

// NewStateHandler creates a handler
func NewStateHandler(auth x.Authenticator) weave.Handler {
	return StateHandler{
		auth: auth,
		b:    NewStateBucket(),
	}
}

// validate does all common pre-processing between Check and Deliver
func (h StateHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*CreateStateMsg, error) {
	var msg CreateStateMsg

	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, errors.Wrap(err, "load msg")
	}

	return &msg, nil
}

// Check just verifies it is properly formed and returns
// the cost of executing it.
func (h StateHandler) Check(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	return &weave.CheckResult{GasAllocated: newStateCost}, nil
}

// Deliver creates an custom state and saves if all preconditions are met
func (h StateHandler) Deliver(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, err := h.validate(ctx, store, tx)

	if err != nil {
		return nil, err
	}

	now := weave.AsUnixTime(time.Now())

	state := &State{
		Metadata:   &weave.Metadata{},
		InnerState: msg.InnerState,
		Address:    msg.Address,
		CreatedAt:  now,
	}

	res, err := h.b.Put(store, nil, state)

	if err != nil {
		return nil, err
	}

	return &weave.DeliverResult{Data: res}, err
}
