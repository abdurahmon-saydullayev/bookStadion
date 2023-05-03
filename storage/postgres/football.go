package postgres

import (
	"database/sql"
	"encoding/json"
	"football/models"
	"football/storage"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
)

type footballRepo struct {
	db *sqlx.DB
}

func NewFootballRepo(db *sqlx.DB) storage.FootballRepoI {
	return footballRepo{
		db: db,
	}
}

func (r footballRepo) Create(entity models.CreateStation) (err error) {
	data, err := json.Marshal(entity.Location)
	if err != nil {
		panic(err)
	}

	loc := types.JSONText(string(data))

	locat, err := loc.Value()
	if err != nil {
		panic(err)
	}
	insertQuery := `INSERT INTO football(
		name,
		from_time,
		to_time,
		price,
		description,
		image,
		stadion_type,
		location
	) values (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	)`

	_, err = r.db.Exec(insertQuery,
		entity.Name,
		entity.From,
		entity.To,
		entity.Price,
		entity.Description,
		entity.Image,
		entity.Stadion_type,
		locat,
	)

	return err
}

func (r footballRepo) GetAll(query models.Query) (resp []models.Station, err error) {
	var rows *sql.Rows
	rows, err = r.db.Query(
		`Select name, price, description, image, stadion_type,location from football
		offset $1 limit $2`,
		query.Offset,
		query.Limit,
	)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next(){
		var a models.Station
		err = rows.Scan(
			&a.Name, &a.Price, &a.Description, &a.Image, &a.Stadion_type,&a.Location,
		)
		resp =append(resp, a)
		if err !=nil{
			return resp, err
		}
	}
	return resp, nil
}

func (r footballRepo) GetById(ID string) (resp models.Station, err error) {
	return resp, nil
}

func (r footballRepo) Update(entity models.StationUpdate) (rows int64, err error) {
	return rows, nil
}

func (r footballRepo) Delete(ID string) (rows int, err error) {
	return rows, nil
}
