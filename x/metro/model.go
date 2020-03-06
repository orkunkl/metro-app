package metro

import (
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
)

var _ orm.SerialModel = (*Station)(nil)

// SetPrimaryKey is a minimal implementation, useful when the PrimaryKey is a separate protobuf field
func (m *Station) SetPrimaryKey(pk []byte) error {
	m.PrimaryKey = pk
	return nil
}

// Validate validates user's fields
func (m *Station) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "PrimaryKey", orm.ValidateSequence(m.PrimaryKey))

	// validate data
	return errs
}

var _ orm.SerialModel = (*Train)(nil)

// SetPrimaryKey is a minimal implementation, useful when the PrimaryKey is a separate protobuf field
func (m *Train) SetPrimaryKey(pk []byte) error {
	m.PrimaryKey = pk
	return nil
}

// Validate validates user's fields
func (m *Train) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "PrimaryKey", orm.ValidateSequence(m.PrimaryKey))
	errs = errors.AppendField(errs, "Address", m.Address.Validate())

	// validate data
	return errs
}

var _ orm.SerialModel = (*Passenger)(nil)

// SetPrimaryKey is a minimal implementation, useful when the PrimaryKey is a separate protobuf field
func (m *Passenger) SetPrimaryKey(pk []byte) error {
	m.PrimaryKey = pk
	return nil
}

// Validate validates user's fields
func (m *Passenger) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "PrimaryKey", orm.ValidateSequence(m.PrimaryKey))
	errs = errors.AppendField(errs, "Address", m.Address.Validate())

	// validate data
	return errs
}
