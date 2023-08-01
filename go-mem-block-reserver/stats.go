package main

import (
	"runtime"
)

type RuntimeInfo struct {
	NumCPU       int    `json:"num_cpu"`
	NumGoroutine int    `json:"num_goroutine"`
	GoMaxProcs   int    `json:"gomaxprocs"`
	Version      string `json:"version"`
}

type Stat struct {
	Runtime  RuntimeInfo      `json:"runtime"`
	MemStats runtime.MemStats `json:"memstats"`
}

func GetStats() Stat {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	runtimeInfo := RuntimeInfo{
		NumCPU:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(),
		GoMaxProcs:   runtime.GOMAXPROCS(0),
		Version:      runtime.Version(),
	}

	return Stat{
		Runtime:  runtimeInfo,
		MemStats: memStats,
	}
}
