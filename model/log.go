package model

import (
	"time"
)

type AccessLog struct {
	TimeStamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}