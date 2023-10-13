package stores

import (
	"sync"
)

// InMemory implements DataStore interface
// using an in-memory storage based on tables/maps
type InMemory struct {
	sync.RWMutex
	datas   []interface{}
	dataIDs map[string]int
}

const (
	// DefaultCapacity is default's in-memory capacity
	DefaultCapacity int = 10
)

// NewInMemory returns an empty in-memory store for data Items
// with provided capacity
func NewInMemory() *InMemory {
	return &InMemory{
		datas:   make([]interface{}, 0, DefaultCapacity),
		dataIDs: make(map[string]int, DefaultCapacity),
	}
}

// GetCapacity returns current  maximum capacity for in-memory store
func (store *InMemory) GetCapacity() int {
	store.RLock()
	defer store.RUnlock()
	return cap(store.datas)
}

// SetCapacity allows to define a newer maximum capacity for in-memory store
// must be greater than prevous one
func (store *InMemory) SetCapacity(newCapacity int) *StoreError {
	store.Lock()
	defer store.Unlock()
	currCapacity := cap(store.datas)
	if newCapacity <= currCapacity {
		return &StoreError{"invalid_capacity", "new capacity must be greater than current one"}
	}
	newDatas := make([]interface{}, 0, newCapacity)
	copy(newDatas, store.datas[:currCapacity])
	store.datas = newDatas
	return nil
}

// GetDatas retrieves data items from in-memory data store
func (store *InMemory) GetDatas(offset, limit int) ([]interface{}, int, *StoreError) {
	store.RLock()
	defer store.RUnlock()
	size := len(store.datas)
	if offset >= size {
		return nil, size, nil
	}
	max := offset + limit
	if max > size {
		max = size
	}
	return store.datas[offset:max], size, nil
}

// AddData adds a data item to in-memory data store
func (store *InMemory) AddData(id string, itemToAdd interface{}) (err *StoreError) {
	store.Lock()
	defer store.Unlock()
	if len(store.datas) >= cap(store.datas) {
		err = &StoreError{"capacity_exceeded", "cannot perform operation for store capacity has been exceeded"}
		return
	}
	store.datas = append(store.datas, itemToAdd)
	store.dataIDs[id] = len(store.datas) - 1
	return
}

// GetData returns a specific item in the store
func (store *InMemory) GetData(id string) (item interface{}, err *StoreError) {
	store.RLock()
	defer store.RUnlock()
	itemIndex, found := store.dataIDs[id]
	if found {
		item = store.datas[itemIndex]
	}
	return
}

// UpdateData updates a specific item in the store
func (store *InMemory) UpdateData(id string, itemToUpdate interface{}) (err *StoreError) {
	store.RLock()
	defer store.RUnlock()
	itemIndex, found := store.dataIDs[id]
	if found {
		store.datas[itemIndex] = itemToUpdate
	}
	return
}

type emptyData interface{}

// DeleteData returns a specific item in the store
func (store *InMemory) DeleteData(id string) (err *StoreError) {
	store.RLock()
	defer store.RUnlock()
	itemIndex, found := store.dataIDs[id]
	if found {
		lastIndex := len(store.datas) - 1
		store.datas[itemIndex] = store.datas[lastIndex]
		store.datas[lastIndex] = nil
		store.datas = store.datas[:lastIndex]
		delete(store.dataIDs, id)
	}
	return nil
}
