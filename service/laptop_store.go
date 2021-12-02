package service

import (
	"errors"
	"fmt"
	"grpc/pb"
	"sync"

	"github.com/jinzhu/copier"
)

type LaptopStore interface {
	Save(l *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

var ErrAlreadyExists = errors.New("record already exists")

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (s *InMemoryLaptopStore) Save(l *pb.Laptop) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.data[l.Id] != nil {
		return ErrAlreadyExists
	}

	other, err := deepCopy(l)
	if err != nil {
		return err
	}

	s.data[other.Id] = other
	return nil
}

func deepCopy(l *pb.Laptop) (*pb.Laptop, error) {
	other := &pb.Laptop{}

	err := copier.Copy(other, l)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}

func (s *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	laptop := s.data[id]
	if laptop == nil {
		return nil, nil
	}

	return deepCopy(laptop)
}
