package storage

type NumberStorageInterface interface {
	IsNumberExists(number uint64) bool
	AddNumber(number uint64) (bool, error)
	GetLength() uint64
}
