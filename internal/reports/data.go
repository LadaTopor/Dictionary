package reports

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// GetReport ищем репорт по id
func (r *Repo) GetReport(id int) (*Report, error) {
	var report Report
	err := r.db.QueryRow(`SELECT id, title, description, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&report.Id, &report.Title, &report.Description, &report.CreatedAt, &report.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// CreateNewReport добавляет новый репорт в базу даных
func (r *Repo) CreateNewReport(title, description string) error {
	_, err := r.db.Exec(`INSERT INTO reports (title, description, created_at, updated_at) VALUES ($1, $2, $3, $3)`, title, description, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// UpdateReport обновляем существующий репорт в базе данных
func (r *Repo) UpdateReport(id int, title, description string) error {
	_, err := r.db.Exec(`UPDATE reports SET title = $1, description = $2, updated_at = $3 WHERE id = $4`, title, description, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteReport удаляем репорт из базы данных
func (r *Repo) DeleteReport(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
