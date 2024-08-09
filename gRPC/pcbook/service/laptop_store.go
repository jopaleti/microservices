package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jopaleti/pcbook/pb"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

// Laptop is an interface to store laptop
type LaptopStore interface {
	// Save laptop to the store
	Save (laptop *pb.Laptop) error 
	Find (id string) (*pb.Laptop, error)
	Search (ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error)  error
}
// type Store interface {
// 	Find (id string) (*pb.Laptop, error)
// }

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
	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

// Find: finds a laptop by ID
func (store *InMemoryLaptopStore) Find (id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// Deep copy
	return deepCopy(laptop)
}

func (store *InMemoryLaptopStore) Search (
	ctx context.Context,
	filter *pb.Filter,
	found func(laptop *pb.Laptop) error,
	) error {
		store.mutex.RLock()
		defer store.mutex.RUnlock()

		for _, laptop := range store.data {
			time.Sleep(time.Second)
			log.Print("Checking laptop id: ", laptop.GetId())

			if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
				log.Print("context is canceled")
				return errors.New("context is canceled")
			}

			if isQualified(filter, laptop) {
				// Deep copy
				other, err := deepCopy(laptop)
				if err != nil {
					return err
				}

				err = found(other)
				if err != nil {
					return err
				}
			}
		}

		return nil
}

func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	// Deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}

	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}

	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}

	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}

	return true
}

func toBit(memory *pb.Memory) uint64 {
	value := memory.GetValue()

	switch memory.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3 // Bit operator
	case pb.Memory_KILOBYTE:
		return value << 13  // 1024 * 8 = 2^10 * 2^3 = 2^13
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}