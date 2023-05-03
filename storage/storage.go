package storage

import "football/models"

type StoragaI interface{
	Football() FootballRepoI
}

type FootballRepoI interface {
	Create(entity models.CreateStation)(err error) 
	GetAll(query models.Query)(resp []models.Station,err error)
	GetById(ID string)(resp models.Station, err error)
	Update(entity models.StationUpdate)(row int64, err error)
	Delete(ID string)(rows int, err error)
}
