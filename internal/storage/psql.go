package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
type URLStore interface {
	SaveURL(r models.URLRecord) error
	UpdateURLInfo(r models.URLRecord) error
	URL_GetAll() ([]models.URLRecord, error)
}
*/

type Postgres struct {
	Pool *pgxpool.Pool
}

func (pg *Postgres) SaveURL(r models.URLRecord) error {
	//check to see if url already exists in db
	var exists bool
	err := pg.Pool.QueryRow(context.Background(), "SELECT EXISTS (SELECT 1 FROM urls WHERE url = $1)", r.URL).Scan(&exists)
	if err != nil {
		return errors.New("error checking for url existence in database")
	}

	if exists {
		return errors.New("url already exists in database, perhaps try updating it instead")
	}

	//insert new url record into url table
	query := `INSERT INTO urls (url, description, check_interval, last_checked_at, last_known_hash, last_known_content, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = pg.Pool.Exec(context.Background(), query,
		r.URL,
		r.Description,
		r.CheckInterval,
		time.Now(), // last_checked_at
		"",         // last_known_hash
		[]string{}, // last_known_content
		time.Now(), // created_at
	)

	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Postgres) UpdateURLInfo(r models.URLRecord) error {
	query := `UPDATE urls SET last_checked_at = $1, last_known_hash =  $2, last_known_content = $3 WHERE url = $4`

	_, err := pg.Pool.Exec(context.Background(), query,
		r.LastCheckedAt,
		r.LastKnownHash,
		r.LastKnownContent,
		r.URL,
	)

	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) URL_GetAll() ([]models.URLRecord, error) {
	query := `SELECT * from urls`
	rows, err := pg.Pool.Query(context.Background(), query)
	if err != nil {
		return []models.URLRecord{}, errors.New("error fetching urls from database")
	}
	defer rows.Close()

	var records []models.URLRecord
	for rows.Next() {
		var r models.URLRecord
		if err := rows.Scan(&r.ID, &r.URL, &r.Description, &r.CheckInterval, &r.LastCheckedAt, &r.LastKnownHash, &r.LastKnownContent, &r.Created_at); err != nil {
			return []models.URLRecord{}, fmt.Errorf("error scanning row: %v", err)
		}
		records = append(records, r)
	}
	return records, nil
}

func (pg *Postgres) LogURLChange(l models.ChangeRecord) error {
	//query := `SELECT `
	return nil
}

func (p *Postgres) ChangeLog_GetAll() ([][]models.ChangeRecord, error) {
	return [][]models.ChangeRecord{}, nil
}
