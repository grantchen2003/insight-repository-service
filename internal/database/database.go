package database

import "sync"

type Database interface {
	Connect() error
	Close() error
	BatchSave(string, string, []interface{}) error
}

var (
	singletonInstance Database
	once              sync.Once
)

func GetInstance() Database {
	once.Do(func() {
		singletonInstance = &MongoDB{}
	})

	return singletonInstance
}
