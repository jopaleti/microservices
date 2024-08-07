package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/jopaleti/pcbook/pb"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

// Laptop is an interface to store laptop
type LaptopStore interface {
	// Save laptop to the store
	Save (laptop *pb.Laptop) error 
}

// InMemoryLaptopStore stores laptop in memory (RAM)
type InMemoryLaptopStore struct {
	// We use mutex to handle concurrency of multiple requests
	// to save laptop on memory
	mutex sync.RWMutex
	data map[string]*pb.Laptop
}

// Function to return a new In-Memory-Laptop-Store
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save function saves laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// Deep copy
	other := &pb.Laptop{}
	// Copy the laptop object parameter to other 
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %v", err)
	}

	store.data[other.Id] = other
	return nil
}