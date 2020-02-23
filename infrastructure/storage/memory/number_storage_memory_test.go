package memory

import (
	"math/rand"
	"testing"
)

// @todo add negative cases

func TestNumberStorage_AddNumber_Basics(t *testing.T) {

	storage := NewNumberStorage()

	var numberToInsert uint64 = 123456789

	status, err := storage.AddNumber(numberToInsert)

	if err != nil {
		t.Errorf("Unexpected error happened: %e.", err)
	}

	if status != true {
		t.Errorf("Expected status: %t, got: %t", true, status)
	}

	storageLen := len(storage.storage)

	if storageLen != 1 {
		t.Errorf("Expected storage length: %d, got: %d", 1, storageLen)
	} else {
		isInserted := storage.storage[numberToInsert]

		if !isInserted {
			t.Errorf("Expected number %d to be in the map", numberToInsert)
		}
	}
}

func TestNumberStorage_AddNumber_Load(t *testing.T) {

	storage := NewNumberStorage()

	for i := 0; i < 10*1000*1000; i++ {
		number := rand.Uint64()

		status, err := storage.AddNumber(number)

		if err != nil {
			t.Errorf("Unexpected error happened: %e.", err)
		}

		if status != true {
			t.Errorf("Expected status: %t, got: %t", true, status)
		}
	}
}

func TestNumberStorage_IsNumberExists_Basic_Positive(t *testing.T) {

	storage := NewNumberStorage()

	var numberToInsert uint64 = 123456789

	storage.storage[numberToInsert] = true

	isExists := storage.IsNumberExists(numberToInsert)

	if !isExists {
		t.Errorf("Expected number %d to be in the map", numberToInsert)
	}
}

func TestNumberStorage_IsNumberExists_Basic_Negative(t *testing.T) {

	storage := NewNumberStorage()

	var numberToInsert uint64 = 123456789

	storage.storage[numberToInsert] = true

	isExists := storage.IsNumberExists(numberToInsert + 1)

	if isExists {
		t.Errorf("Expected number %d not to be in the map", numberToInsert)
	}
}

func TestNumberStorage_IsNumberExists_Load(t *testing.T) {

	storage := NewNumberStorage()

	for i := 0; i < 10*1000*1000; i++ {
		number := rand.Uint64()
		storage.storage[number] = true
	}

	for i := 0; i < 10*1000*1000; i++ {
		number := rand.Uint64()

		storage.IsNumberExists(number)
	}
}

func TestNumberStorage_GetLength(t *testing.T) {
	storage := NewNumberStorage()

	countOfMessages := 10 * 1000 * 1000
	var expectedAmountOfUniqNumbers uint64 = 0

	for i := 0; i < countOfMessages; i++ {
		number := rand.Uint64()
		status, _ := storage.AddNumber(number)

		if status {
			expectedAmountOfUniqNumbers++
		}
	}

	if storage.GetLength() != expectedAmountOfUniqNumbers {
		t.Errorf("Expected storage length: %d, got: %d", expectedAmountOfUniqNumbers, storage.GetLength())
	}
}
