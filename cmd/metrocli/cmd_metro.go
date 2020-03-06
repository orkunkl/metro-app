package main

import (
	"flag"
	"fmt"
	"io"

	app "github.com/orkunkl/metro-app/cmd/metro/app"
	"github.com/orkunkl/metro-app/x/metro"
	"github.com/iov-one/weave"
)

func cmdRegisterPassenger(input io.Reader, output io.Writer, args []string) error {
	fl := flag.NewFlagSet("", flag.ExitOnError)
	fl.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Register a passenger.
		`)
		fl.PrintDefaults()
	}

	fl.Parse(args)

	msg := metro.RegisterPassengerMsg{
		Metadata: &weave.Metadata{Schema: 1},
	}

	if err := msg.Validate(); err != nil {
		return fmt.Errorf("given data produce an invalid message: %s", err)
	}

	tx := &app.Tx{
		Sum: &app.Tx_MetroRegisterPassengerMsg{
			MetroRegisterPassengerMsg: &msg,
		},
	}
	_, err := writeTx(output, tx)
	return err
}

func cmdTrainArriveStation(input io.Reader, output io.Writer, args []string) error {
	fl := flag.NewFlagSet("", flag.ExitOnError)
	fl.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Train arrives to station
		`)
		fl.PrintDefaults()
	}
	var (
		stationFl = flSeq(fl, "station_key", "", "Primary key of a station")
		trainFl   = flSeq(fl, "train_key", "", "Primary key of a station")
	)
	fl.Parse(args)

	msg := metro.TrainArriveStationEventMsg{
		Metadata:   &weave.Metadata{Schema: 1},
		StationKey: *stationFl,
		TrainKey:   *trainFl,
	}
	if err := msg.Validate(); err != nil {
		return fmt.Errorf("given data produce an invalid message: %s", err)
	}

	tx := &app.Tx{
		Sum: &app.Tx_MetroTrainArriveStationEventMsg{
			MetroTrainArriveStationEventMsg: &msg,
		},
	}
	_, err := writeTx(output, tx)
	return err
}
