package metro

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

// Initializer fulfils the Initializer interface to load data from the genesis
// file
type Initializer struct{}

var _ weave.Initializer = (*Initializer)(nil)

// FromGenesis will parse initial account info from genesis and save it to the
// database
func (*Initializer) FromGenesis(opts weave.Options, params weave.GenesisParams, kv weave.KVStore) error {
	var input struct {
		Station []struct {
			Station      string `json:"station"`
			Escalator    int64  `json:"escalator"`
			Elevator     int64  `json:"elevator"`
			IsPeronAda   bool   `json:"is_peron_ada"`
			TicketOffice int64  `json:"ticket_office"`
			TollGateEnt  int64  `json:"toll_gate_ent"`
			TollGateEx   int64  `json:"toll_gate_ex"`
			EntranceExit int64  `json:"entrance_exit"`
		}
		Train []struct {
			Address weave.Address `json:"address"`
		}
		Passenger []struct {
			Address      weave.Address `json:"address"`
			RegisteredAt int64         `json:"registered_at"`
		}
	}

	switch err := opts.ReadOptions("metro", &input); {
	case err == nil:
		// All good.
	case errors.ErrNotFound.Is(err):
		// No configuration defined.
		return nil
	default:
		return errors.Wrap(err, "cannot load station data")
	}

	stations := NewStationBucket()
	for _, d := range input.Station {
		station := Station{
			Metadata:     &weave.Metadata{Schema: 1},
			Station:      d.Station,
			Escalator:    d.Escalator,
			Elevator:     d.Elevator,
			IsPeronAda:   d.IsPeronAda,
			TicketOffice: d.TicketOffice,
			TollGateEnt:  d.TollGateEnt,
			TollGateEx:   d.TollGateEx,
			EntranceExit: d.EntranceExit,
		}
		if err := stations.Save(kv, &station); err != nil {
			return errors.Wrapf(err, "cannot store %d station", d)
		}
	}

	trains := NewTrainBucket()
	for _, d := range input.Train {
		train := Train{
			Metadata: &weave.Metadata{Schema: 1},
			Address:  d.Address,
		}
		if err := trains.Save(kv, &train); err != nil {
			return errors.Wrapf(err, "cannot store %d train", d)
		}
	}

	passengers := NewPassengerBucket()
	for _, d := range input.Passenger {
		passenger := Passenger{
			Metadata: &weave.Metadata{Schema: 1},
			Address:  d.Address,
		}
		if err := passengers.Save(kv, &passenger); err != nil {
			return errors.Wrapf(err, "cannot store %d passenger", d)
		}
	}

	return nil
}
