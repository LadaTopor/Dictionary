package reports

import (
	"database/sql"
	"fmt"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetReportById(id int) (*Report, error) {
	var report Report
	err := r.db.QueryRow(`SELECT id, title, description, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&report.Id, &report.Title, &report.Description, &report.CreatedAt, &report.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *Repo) CreateReport(title, description string) error {
	_, err := r.db.Exec(`INSERT INTO reports (title, description, created_at, updated_at) VALUES ($1, $2, now(), now())`, title, description)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateReport(id int, title, description string) error {
	var returnedID int
	err := r.db.QueryRow(`UPDATE reports SET title = $1, description = $2, updated_at = now() WHERE id = $3 RETURNING id`, title, description, id).Scan(&returnedID)
	if err != nil {
		return fmt.Errorf("kto obzyvaetcya", id)
	}
	return nil
}

func (r *Repo) DeleteReport(id int) error {
	var returnedID int
	err := r.db.QueryRow(`DELETE FROM reports WHERE id = $1 RETURNING id`, id).Scan(&returnedID)
	if err != nil {
		return fmt.Errorf("tot cam tak nazyvaetcya", id)
	}
	return nil
}
