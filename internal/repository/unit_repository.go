package repository

import (
	"database/sql"
	"unit-api/internal/database"
	"unit-api/internal/models"
)

func GetUnits(status string) ([]models.Unit, error) {

	query := "SELECT id, name, type, status, lastUpdated FROM units"

	if status != "" {
		query += " WHERE status = ?"
	}

	var rows *sql.Rows
	var err error

	if status != "" {
		rows, err = database.DB.Query(query, status)
	} else {
		rows, err = database.DB.Query(query)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var units []models.Unit

	for rows.Next() {
		var unit models.Unit

		err := rows.Scan(
			&unit.ID,
			&unit.Name,
			&unit.Type,
			&unit.Status,
			&unit.LastUpdated,
		)

		if err != nil {
			return nil, err
		}
		units = append(units, unit)
	}

	return units, nil
}

func GetUnitByID(id int) (*models.Unit, error) {
	query := "SELECT id, name, type, status, lastUpdated FROM units WHERE id = ?"

	row:= database.DB.QueryRow(query, id)

	var unit models.Unit

	err := row.Scan(
		&unit.ID,
		&unit.Name,
		&unit.Type,
		&unit.Status,
		&unit.LastUpdated,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &unit, nil
}

func CreateUnit(unit *models.Unit) error {
	query := `INSERT INTO units (name, type, status, lastUpdated) VALUES (?, ?, ?, NOW())`

	_, err := database.DB.Exec(query, unit.Name, unit.Type, unit.Status)
	return err
}

func UpdateUnitStatus(id int, status string) error {
	query := `UPDATE units SET status = ?, lastUpdated = NOW() WHERE id = ?`
	_, err := database.DB.Exec(query, status, id)
	return err
}