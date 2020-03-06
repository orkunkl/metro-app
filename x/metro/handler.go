package metro

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/x"
)

const (
	packageName               = "metro"
)

// RegisterQuery registers buckets for querying.
func RegisterQuery(qr weave.QueryRouter) {
	NewStationBucket().Register("stations", qr)
}

// RegisterRoutes registers handlers for message processing.
func RegisterRoutes(r weave.Registry, auth x.Authenticator) {
	//r = migration.SchemaMigratingRegistry(packageName, r)
	r.Handle(&RegisterPassengerMsg{}, NewRegisterPassengerHandler(auth))
	r.Handle(&TrainArriveStationEventMsg{}, NewTrainArriveStationEventHandler(auth))
}

// ------------------- RegisterPassengerHandler -------------------

// RegisterPassengerHandler will handle RegisterPassengerMSg
type RegisterPassengerHandler  struct {
	auth x.Authenticator
	b    orm.SerialModelBucket
}

var _ weave.Handler = RegisterPassengerHandler{}

// NewRegisterPassengerHandler creates a passenger message handler
func NewRegisterPassengerHandler(auth x.Authenticator) weave.Handler {
	return RegisterPassengerHandler{
		auth: auth,
		b:    NewPassengerBucket(),
	}
}

// validate does all common pre-processing between Check and Deliver
func (h RegisterPassengerHandler) validate(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*RegisterPassengerMsg, *Passenger, error) {
	var msg RegisterPassengerMsg

	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}

	blockTime, err := weave.BlockTime(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "no block time in header")
	}
	now := weave.AsUnixTime(blockTime)

	p := &Passenger{
		Metadata:      &weave.Metadata{Schema: 1},
		Address:      x.AnySigner(ctx, h.auth).Address(),
		RegisteredAt: now,
	}

	return &msg, p, nil
}

// Check just verifies it is properly formed and returns
// the cost of executing it.
func (h RegisterPassengerHandler) Check(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, _, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	return &weave.CheckResult{}, nil
}

// Deliver creates an custom state and saves if all preconditions are met
func (h RegisterPassengerHandler) Deliver(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	_, passenger, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	err = h.b.Save(store, passenger)
	if err != nil {
		return nil, errors.Wrap(err, "cannot store passenger")
	}

	// Returns generated user PrimaryKey as response
	return &weave.DeliverResult{Data: passenger.PrimaryKey}, nil
}

// ------------------- TrainArriveStationEventHandler -------------------

// TrainArriveStationEventHandler will handle TrainArriveStationEventMsg
type TrainArriveStationEventHandler struct {
	auth x.Authenticator
	b    orm.SerialModelBucket
}

var _ weave.Handler = TrainArriveStationEventHandler{}

// NewTrainArriveStationEventHandler creates a event message handler
func NewTrainArriveStationEventHandler(auth x.Authenticator) weave.Handler {
	return TrainArriveStationEventHandler{
		auth: auth,
		b:    NewTrainArriveStationEventBucket(),
	}
}

// validate does all common pre-processing between Check and Deliver
func (h TrainArriveStationEventHandler) validate(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*TrainArriveStationEventMsg, *TrainArriveStationEvent, error) {
	var msg TrainArriveStationEventMsg

	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}

	blockTime, err := weave.BlockTime(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "no block time in header")
	}
	now := weave.AsUnixTime(blockTime)

	tae := &TrainArriveStationEvent{
		Metadata:   &weave.Metadata{Schema: 1},
		StationKey: msg.StationKey,
		TrainKey:   msg.TrainKey,
		ArrivedAt:  now,
	}

	return &msg, tae, nil
}

// Check just verifies it is properly formed and returns
// the cost of executing it.
func (h TrainArriveStationEventHandler) Check(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, _, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	return &weave.CheckResult{}, nil
}

// Deliver creates an custom state and saves if all preconditions are met
func (h TrainArriveStationEventHandler) Deliver(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	_, tae, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	err = h.b.Save(store, tae)
	if err != nil {
		return nil, errors.Wrap(err, "cannot store passenger")
	}

	// Returns generated user PrimaryKey as response
	return &weave.DeliverResult{Data: tae.PrimaryKey}, nil
}

