package models

import (
	"time"
)

type Station struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	From         *time.Time `json:"from_time"`
	To           *time.Time `json:"to_time"`
	Price        float64    `json:"price"`
	Rating       int        `json:"rating"`
	Description  string     `json:"description"`
	Image        string     `json:"image"`
	Stadion_type string     `json:"stadion_type"`
	Location     string     `json:"location"`
}

type CreateStation struct {
	Name         string     `json:"name"`
	From         *time.Time `json:"from_time"`
	To           *time.Time `json:"to_time"`
	Price        float64    `json:"price"`
	Description  string     `json:"description"`
	Image        string     `json:"image"`
	Stadion_type string     `json:"stadion_type"`
	Location     *Location     `json:"location"`
}

type Location struct {
	Long string `json:"long"`
	Lat string `json:"lat"`
	From_Home string `json:"from_home"`
}

type StationUpdate struct {
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
	Stadion_type string  `json:"stadion_type"`
}
