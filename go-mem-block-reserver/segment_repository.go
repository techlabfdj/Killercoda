package main

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type SegmentRepository struct {
	segments map[uuid.UUID]Segment
	mu       sync.RWMutex
}

func NewSegmentRepository() *SegmentRepository {
	return &SegmentRepository{
		segments: make(map[uuid.UUID]Segment),
	}
}

func (r *SegmentRepository) Create(s Segment) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.segments[s.ID] = s
}

func (r *SegmentRepository) Get(id uuid.UUID) (Segment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	segment, ok := r.segments[id]
	if !ok {
		return Segment{}, errors.New("Segment not found")
	}

	return segment, nil
}

func (r *SegmentRepository) Update(s Segment) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.segments[s.ID]
	if !ok {
		return errors.New("Segment not found")
	}

	r.segments[s.ID] = s
	return nil
}

func (r *SegmentRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.segments[id]
	if !ok {
		return errors.New("Segment not found")
	}

	delete(r.segments, id)
	return nil
}

func (r *SegmentRepository) List() []Segment {
	r.mu.RLock()
	defer r.mu.RUnlock()

	segments := make([]Segment, 0, len(r.segments))
	for _, segment := range r.segments {
		segments = append(segments, segment)
	}

	return segments
}
