package metro

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/migration"
)

func init() {
	migration.MustRegister(1, &RegisterPassengerMsg{}, migration.NoModification)
	migration.MustRegister(1, &TrainArriveStationEventMsg{}, migration.NoModification)
}

var _ weave.Msg = (*RegisterPassengerMsg)(nil)

// Path returns the routing path for this message.
func (RegisterPassengerMsg) Path() string {
	return "metro/register_passenger"
}

// Validate ensures the RegisterPassengerMsg is valid
func (m RegisterPassengerMsg) Validate() error {
	// data to validate
	return nil
}

var _ weave.Msg = (*TrainArriveStationEventMsg)(nil)

// Path returns the routing path for this message.
func (TrainArriveStationEventMsg) Path() string {
	return "metro/train_arrive_station"
}

// Validate ensures the TrainArriveStationEventMsg is valid
func (m TrainArriveStationEventMsg) Validate() error {
	// data to validate
	return nil
}
