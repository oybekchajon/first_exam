package main

import (
	"database/sql"
	"time"
)

type Car struct{
	ID int64
	Brand string
	Year int32
	Color string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Images struct {
	ID int64
	ImageUrl string
	SequenceNumber int32
}

type DBManager struct{
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager{
	return &DBManager{db}
}

func (m *DBManager) CreateCar(car *Car) (*Car, error){
	time := time.Now()
	query := `
		INSERT INTO car (brand, year, color, created_at)
		VALUES ($1, $2, $3, $4)
	`

	row := m.db.QueryRow(query,
		car.Brand,
		car.Year,
		car.Color,
		time,
	)

	var result Car 

	err := row.Scan(
		&result.ID,
		&result.Brand,
		&result.Year,
		&result.Color,
		&result.CreatedAt,
	)

	

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *DBManager) GetCar (id int) (*Car, error){
	query := `
		SELECT 
			id,
			brand,
			year,
			color
		FROM car
		WHERE id=$1
	`

	row := m.db.QueryRow(query,id)

	var result Car

	err := row.Scan(
		&result.ID,
		&result.Brand,
		&result.Year,
		&result.Color,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil	
}

func (m *DBManager) UpdateCar(car *Car) (*Car, error){
	time := time.Now()

	query := `
		UPDATE car SET
			brand=$1,
			year=$2,
			color=$3,
			updated_at=$4
			WHERE id = $5
			RETURNING year, brand, color, created_at, updated_at
	`

	row := m.db.QueryRow(query,
		car.Brand,
		car.Year,
		car.Color,
		time,
		car.ID,
	)

	var result Car

	err := row.Scan(
		&result.Year,
		&result.Brand,
		&result.Color,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *DBManager) DeleteCar(id int32) (*Car, error){
	time := time.Now()

	query := `
		UPDATE car SET 
			deleted_at=$1
		WHERE id=$2
		RETURNING id, brand, year, color, created_at, deleted_at
	`

	row := m.db.QueryRow(query,
		time,
		id,
	)

	var result Car 

	err := row.Scan(
		&result.ID,
		&result.Brand,
		&result.Year,
		&result.Color,
		&result.CreatedAt,
		&result.DeletedAt,
	)	

	if err != nil {
		return nil, err
	}

	return &result, nil
}