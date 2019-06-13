package fperf

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestDeferAnonymous(t *testing.T) {
	f := func() {
		defer MeasureRunTime(report)()
	}
	f()
}

func TestDeferNamed(t *testing.T) {
	namedWithDefer()
}

func TestDeferMethod(t *testing.T) {
	var i integer
	i.methodWithDefer()
}

func TestEmbedAnonymous(t *testing.T) {
	f := func() {
	}
	ef := EmbedRunTimeMeasurement(f, report).(func())
	ef()
}

func TestEmbedNamed(t *testing.T) {
	ef := EmbedRunTimeMeasurement(named, report).(func())
	ef()
}

func TestEmbedMethod(t *testing.T) {
	var i integer
	ef := EmbedRunTimeMeasurement(i.method, report).(func())
	ef()
}

func namedWithDefer() {
	defer MeasureRunTime(report)()
}

func named() {
}

type integer int

func (i integer) methodWithDefer() {
	defer MeasureRunTime(report)()
}

func (i integer) method(ii int) {
	panic(ii)
}

func report(name string, duration time.Duration) {
	fmt.Fprintf(os.Stderr, "%q func took %v\n", name, duration)
}
