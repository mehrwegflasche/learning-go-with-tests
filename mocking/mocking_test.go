package main

import (
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

	spySleepPrinter := &SpySleeperPrinter{}
	Countdown(spySleepPrinter, spySleepPrinter)

	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
		t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
	}

}
