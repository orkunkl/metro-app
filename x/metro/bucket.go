package metro

import (
	"github.com/iov-one/weave/orm"
)

type StationBucket struct {
	orm.SerialModelBucket
}

// NewStationBucket returns a new station bucket
func NewStationBucket() orm.SerialModelBucket {
	b := &StationBucket{
		orm.NewSerialModelBucket("station", &Station{}),
	}
	return b
}

type TrainBucket struct {
	orm.SerialModelBucket
}

// NewTrainBucket returns a new train bucket
func NewTrainBucket() orm.SerialModelBucket {
	b := &TrainBucket{
		orm.NewSerialModelBucket("train", &Train{}),
	}
	return b
}

type PassengerBucket struct {
	orm.SerialModelBucket
}

// NewPassengerBucket returns a new passenger bucket
func NewPassengerBucket() orm.SerialModelBucket {
	b := &PassengerBucket{
		orm.NewSerialModelBucket("pass", &Passenger{}),
	}
	return b
}

type TrainArriveStationEventBucket struct {
	orm.SerialModelBucket
}

// NewTrainArriveStationEvent returns a new train event bucket
func NewTrainArriveStationEventBucket() orm.SerialModelBucket {
	b := &TrainArriveStationEventBucket{
		orm.NewSerialModelBucket("traiarr", &TrainArriveStationEvent{}),
	}
	return b
}
