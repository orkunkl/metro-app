package metro

import (
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
)

var _ orm.SerialModel = (*TrainArriveStationEvent)(nil)

// SetPrimaryKey is a minimal implementation, useful when the PrimaryKey is a separate protobuf field
func (e *TrainArriveStationEvent) SetPrimaryKey(pk []byte) error {
	e.PrimaryKey = pk
	return nil
}

// Validate validates user's fields
func (m *TrainArriveStationEvent) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "PrimaryKey", orm.ValidateSequence(m.PrimaryKey))

	// validate data
	return errs
}

