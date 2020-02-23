package memory

import (
	"errors"
)

type NumberStorage struct {
	storage       map[uint64]bool
	elementsCount uint64
}

func (s *NumberStorage) IsNumberExists(number uint64) bool {
	if val, ok := s.storage[number]; ok {
		return val
	}

	return false
}

func (s *NumberStorage) AddNumber(number uint64) (bool, error) {

	if s.IsNumberExists(number) {
		return false, errors.New("number exists: " + string(number))
	}

	s.storage[number] = true
	s.elementsCount++

	return true, nil
}

func (s *NumberStorage) GetLength() uint64 {
	return s.elementsCount
}

func NewNumberStorage() *NumberStorage {
	storage := make(map[uint64]bool)

	return &NumberStorage{storage: storage, elementsCount: 0}
}
