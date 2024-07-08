package database

import "sync"

type Repository struct {
	Id            string
	IsInitialized bool
}

type Database interface {
	Connect() error
	Close() error
	CreateRepository() (string, error)
	GetRepositoryById(string) (*Repository, error)
	DeleteRepository(string) error
}

var (
	singletonInstance Database
	once              sync.Once
)

func GetSingletonInstance() Database {
	once.Do(func() {
		singletonInstance = &MongoDb{}
	})

	return singletonInstance
}
