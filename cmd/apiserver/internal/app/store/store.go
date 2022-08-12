package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Start...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

//New...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open...
func (s *Store) Open() error {
	db, err := sql.Open("Postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

//Close...
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository

}
