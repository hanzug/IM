package wkserver

import "time"

type Options struct {
	Addr            string
	RequestPoolSize int
	ConnPath        string
	ClosePath       string
	RequestTimeout  time.Duration
}

func NewOptions() *Options {

	return &Options{
		Addr:            "tcp://0.0.0.0:12000",
		RequestPoolSize: 1000,
		ConnPath:        "/conn",
		ClosePath:       "/close",
		RequestTimeout:  10 * time.Second,
	}
}
