package metro

import (
	"github.com/iov-one/weave/migration"

	"github.com/iov-one/weave/orm"
)

type StationBucket struct {
	orm.SerialModelBucket
}

// NewStationBucket returns a new station bucket
func NewStationBucket() orm.SerialModelBucket {
	b := &StationBucket{
		orm.NewSerialModelBucket("stati", &Station{}),
	}
	return migration.NewSerialModelBucket("station", &Station{}, b)
}

type TrainBucket struct {
	orm.SerialModelBucket
}

// NewTrainBucket returns a new train bucket
func NewTrainBucket() orm.SerialModelBucket {
	b := &TrainBucket{
		orm.NewSerialModelBucket("trai", &Train{}),
	}
	return migration.NewSerialModelBucket("train", &Train{}, b)
}

type PassengerBucket struct {
	orm.SerialModelBucket
}

// NewPassengerBucket returns a new passenger bucket
func NewPassengerBucket() orm.SerialModelBucket {
	b := &PassengerBucket{
		orm.NewSerialModelBucket("pass", &Passenger{}),
	}
	return migration.NewSerialModelBucket("passngr", &Passenger{}, b)
}

type TrainArriveStationEventBucket struct {
	orm.SerialModelBucket
}

// NewTrainArriveStationEvent returns a new train event bucket
func NewTrainArriveStationEventBucket() orm.SerialModelBucket {
	b := &TrainArriveStationEventBucket{
		orm.NewSerialModelBucket("traiarr", &TrainArriveStationEvent{}),
	}
	return migration.NewSerialModelBucket("trainarr", &TrainArriveStationEvent{}, b)
}
