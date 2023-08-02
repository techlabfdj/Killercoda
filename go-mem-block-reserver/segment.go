package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Segment struct {
	ID        uuid.UUID `json:"id"`
	Size      int       `json:"size"`
	Unit      string    `json:"unit"`
	Data      []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func getByteMultiplier(unit string) int {
	var multiplier int
	switch strings.ToUpper(unit) {
	case "BYTES":
		multiplier = 1
	case "KILOBYTES":
		multiplier = 1024
	case "MEGABYTES":
		multiplier = (1024 * 1024)
	case "GIGABYTES":
		multiplier = (1024 * 1024 * 1024)
	}
	return multiplier
}

func maxValue(unit string) int {
	return math.MaxInt64 / getByteMultiplier(unit)
}

func NewSegment(size int, unit string) (*Segment, error) {
	// Add error checking for size and unit here
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}

	unit = strings.ToUpper(unit)
	if unit != "BYTES" && unit != "KILOBYTES" && unit != "MEGABYTES" && unit != "GIGABYTES" {
		return nil, errors.New("invalid unit")
	}

	if size > maxValue(unit) {
		return nil, errors.New("size exceeds maximum allowed capacity")
	}

	// Use unit to compute real size
	realSize := size * getByteMultiplier(unit)

	data := make([]byte, realSize)
	fmt.Printf("%d bytes allocated !", realSize)

	return &Segment{
		ID:        uuid.New(),
		Size:      size,
		Unit:      unit,
		Data:      data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
