package postgres

import (
	"football/storage"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sqlx.DB
	football storage.FootballRepoI
	
}

func NewPostgres(psqlConnString string) storage.StoragaI {
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Panic(err)
	}

	return &Store{
		db: db,
	}
}

func (s *Store) Football() storage.FootballRepoI {
	if s.football == nil {
		s.football = &footballRepo{db: s.db}
	}
	return s.football
}


