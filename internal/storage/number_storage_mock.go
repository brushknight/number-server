package storage

type NumberStorageMock struct {
	storage                       map[uint64]bool
	methodAddNumberCalls          map[uint64]uint64
	methodIsNumberExistsCalls     map[uint64]uint64
	methodAddNumberCallTimes      uint64
	methodIsNumberExistsCallTimes uint64
}

func (s *NumberStorageMock) MethodCalledTimes(method string) uint64 {
	switch method {
	case "IsNumberExists":
		return s.methodIsNumberExistsCallTimes
	case "AddNumber":
		return s.methodAddNumberCallTimes
	}
	return 0
}

func (s *NumberStorageMock) MethodCalledTimesWithValue(method string, number uint64) uint64 {
	switch method {
	case "IsNumberExists":
		if isNumberInTheStorageMap(s.methodIsNumberExistsCalls, number) {
			return s.methodIsNumberExistsCalls[number]
		}
		return 0
	case "AddNumber":
		if isNumberInTheStorageMap(s.methodAddNumberCalls, number) {
			return s.methodAddNumberCalls[number]
		}
		return 0
	}
	return 0
}

func (s *NumberStorageMock) IsNumberExists(number uint64) bool {
	s.methodIsNumberExistsCallTimes++
	if isNumberInTheStorageMap(s.methodIsNumberExistsCalls, number) {
		s.methodIsNumberExistsCalls[number] = s.methodIsNumberExistsCalls[number] + 1
	} else {
		s.methodIsNumberExistsCalls[number] = 1
	}

	if _, ok := s.storage[number]; ok {
		return true
	}

	return false
}

func (s *NumberStorageMock) AddNumber(number uint64) (bool, error) {
	s.methodAddNumberCallTimes++
	if isNumberInTheStorageMap(s.methodAddNumberCalls, number) {
		s.methodAddNumberCalls[number] = s.methodAddNumberCalls[number] + 1
	} else {
		s.methodAddNumberCalls[number] = 1
	}

	s.storage[number] = true

	return true, nil
}

func (s *NumberStorageMock) GetLength() uint64 {
	return uint64(len(s.storage))
}

func isNumberInTheStorageMap(checkingMap map[uint64]uint64, number uint64) bool {
	if _, ok := checkingMap[number]; ok {
		return true
	}

	return false
}

func NewMockStorage() *NumberStorageMock {
	storage := make(map[uint64]bool)
	methodIsNumberExistsCalls := make(map[uint64]uint64)
	methodAddNumberCalls := make(map[uint64]uint64)
	return &NumberStorageMock{storage: storage, methodIsNumberExistsCalls: methodIsNumberExistsCalls, methodAddNumberCalls: methodAddNumberCalls}
}
