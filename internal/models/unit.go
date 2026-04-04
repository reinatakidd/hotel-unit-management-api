package models

import "time"

type Unit struct {
	ID int 	 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Status string `json:"status"`
	LastUpdated time.Time `json:"last_updated"`
}