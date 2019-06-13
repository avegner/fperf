package fperf

import (
	"reflect"
	"runtime"
	"time"
)

type ReportCallback func(name string, duration time.Duration)

func EmbedRunTimeMeasurement(f interface{}, report ReportCallback) interface{} {
	rt := reflect.TypeOf(f)
	if rt.Kind() != reflect.Func {
		panic("f isn't a func")
	}

	tf := reflect.MakeFunc(rt, func(args []reflect.Value) (results []reflect.Value) {
		rv := reflect.ValueOf(f)
		name := runtime.FuncForPC(rv.Pointer()).Name()

		start := time.Now()
		results = rv.Call(args)
		dur := time.Since(start)

		report(name, dur)
		return results
	})

	return tf.Interface()
}

func MeasureRunTime(report ReportCallback) func() {
	name := "!unknown!"
	if pc, _, _, ok := runtime.Caller(1); ok {
		name = runtime.FuncForPC(pc).Name()
	}

	start := time.Now()
	return func() {
		report(name, time.Since(start))
	}
}
