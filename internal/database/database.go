package database

import "sync"

type Repository struct {
	Id            string
	SessionId     string
	IsInitialized bool
}

type Database interface {
	Connect() error
	Close() error
	CreateRepository(string) (string, error)
	GetRepositoryBySessionId(string) (*Repository, error)
	SetRepositoryIsInitialized(string, bool) error
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
