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

type Supabase struct {
	Pool *pgxpool.Pool
}

func (pg *Supabase) SaveURL(r models.URLRecord) error {
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
	query := `INSERT INTO urls (url, user_id, description, check_interval, last_checked_at, last_known_hash, last_known_content, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = pg.Pool.Exec(context.Background(), query,
		r.URL,
		r.User_id,
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

func (pg *Supabase) UpdateURLInfo(r models.URLRecord) error {
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

func (pg *Supabase) URL_GetAll() ([]models.URLRecord, error) {
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

func (pg *Supabase) URL_GetOne(urlID int) (models.URLRecord, error) {
	query1 := `SELECT EXISTS (SELECT 1 FROM urls WHERE id = $1)`
	query2 := `SELECT * from urls WHERE id = $1`

	var exists bool
	err := pg.Pool.QueryRow(context.Background(), query1, urlID).Scan(&exists)

	if err != nil || !exists {
		return models.URLRecord{}, fmt.Errorf("error checking for existence of queried url or url does not exist in database: %v", err)
	}

	var r models.URLRecord
	row := pg.Pool.QueryRow(context.Background(), query2, urlID)
	err = row.Scan(&r.ID, &r.URL, &r.Description, &r.CheckInterval, &r.LastCheckedAt, &r.LastKnownHash, &r.LastKnownContent, &r.Created_at)

	if err != nil {
		return models.URLRecord{}, fmt.Errorf("error scanning row: %v", err)
	}

	return r, nil

}

func (pg *Supabase) URL_Delete(urlID int) error {
	query := `DELETE FROM urls WHERE id = $1`
	res, err := pg.Pool.Exec(context.Background(), query, urlID)

	if err != nil {
		return fmt.Errorf("error deleting url from urls table: %v", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("id could not be deleted because it couldn't be find in table: %d", urlID)
	}

	return nil
}

func (pg *Supabase) LogURLChange(l models.ChangeRecord) error {
	//find id of desired url
	var urlID int
	query := `SELECT id from urls WHERE url = $1`
	err := pg.Pool.QueryRow(context.Background(), query, l.URL).Scan(&urlID)
	if err != nil {
		return fmt.Errorf("error querying for url id in url table: %v", err)
	}

	//create new log record in change logs table
	query = `INSERT INTO changelogs (url_id, url, timestamp, added, diff_summary) VALUES ($1, $2, $3, $4, $5)`
	_, err = pg.Pool.Exec(context.Background(), query,
		urlID,
		l.URL,
		time.Now(),
		l.Added,
		l.DiffSummary,
	)

	if err != nil {
		return fmt.Errorf("error inserting new change record into change logs table: %v", err)
	}

	return nil
}

func (pg *Supabase) ChangeLog_GetAll() ([]models.ChangeRecord, error) {
	query := `SELECT * FROM changelogs`
	rows, err := pg.Pool.Query(context.Background(), query)

	if err != nil {
		return []models.ChangeRecord{}, errors.New("error fetching all records from change logs table")
	}

	defer rows.Close()

	var logs []models.ChangeRecord
	for rows.Next() {
		var r models.ChangeRecord
		if err := rows.Scan(&r.ID, &r.URL_id, &r.URL, &r.Timestamp, &r.Added, &r.DiffSummary); err != nil {
			return []models.ChangeRecord{}, fmt.Errorf("error scanning row: %v", err)
		}
		logs = append(logs, r)
	}
	return logs, nil
}

func (pg *Supabase) ChangeLog_GetOneUrl(urlID int) ([]models.ChangeRecord, error) {
	query := `SELECT * FROM changelogs WHERE url_id = $1`
	rows, err := pg.Pool.Query(context.Background(), query, urlID)

	if err != nil {
		return nil, fmt.Errorf("error fetching change log records for specified url id: %d", urlID)
	}

	defer rows.Close()

	var records []models.ChangeRecord
	for rows.Next() {
		var r models.ChangeRecord
		if err := rows.Scan(&r.ID, &r.URL_id, &r.URL, &r.Timestamp, &r.Added, &r.DiffSummary); err != nil {
			return []models.ChangeRecord{}, fmt.Errorf("error scanning row: %v", err)
		}
		records = append(records, r)
	}

	return records, nil
}
