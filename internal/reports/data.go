package reports

import (
	"database/sql"
	"strconv"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// RGetReportById ищем репорт по id
func (r *Repo) RGetReportById(id int) (*Reports, error) {
	var reports Reports
	err := r.db.QueryRow(`SELECT id, Title, description, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&reports.Id, &reports.Title, &reports.Description, &reports.Created_at, &reports.Updated_at)
	if err != nil {
		return nil, err
	}

	return &reports, nil
}

func (r *Repo) CreateNewReports(title, description, time string) (*Reports, error) {
	_, err := r.db.Exec(`INSERT INTO reports (title, description, created_at, updated_at) VALUES ($1, $2, $3, $4)`, title, description, time)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (r *Repo) NewUpdateReportById(id, title, description, updateTime string) (*Reports, error) {
	idInt, err := strconv.Atoi(id)
	_, err = r.db.Exec(`UPDATE reports SET title = $2, description = $3, updated_at = $4  WHERE id = $1`, idInt, title, description, updateTime)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (r *Repo) DeleteReport(id int) (*Reports, error) {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return nil, err
}
