package reports

import (
	"database/sql"
	"github.com/lib/pq"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// RGetReportById ищем репорт по id
func (r *Repo) RGetReportById(ids []int) ([]Reports, error) {
	var reports []Reports
	var args []any

	query := `SELECT id, title, description, created_at, updated_at FROM reports `
	where := `WHERE id = ANY($1)`
	if len(ids) != 0 {
		query += where
		args = append(args, pq.Array(ids))
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var report Reports
		err := rows.Scan(&report.Id, &report.Title, &report.Description, &report.Created_at, &report.Updated_at)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func (r *Repo) CreateNewReport(title, description, time string) error {
	_, err := r.db.Exec(`INSERT INTO reports (title, description, created_at, updated_at) VALUES ($1, $2, $3, $4)`, title, description, time)
	if err != nil {
		return nil
	}

	return nil
}

func (r *Repo) NewUpdateReportById(id int, title, description, updateTime string) (*Reports, error) {
	_, err := r.db.Exec(`UPDATE reports SET title = $2, description = $3, updated_at = $4  WHERE id = $1`, id, title, description, updateTime)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (r *Repo) DeleteReport(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	if err != nil {
		return nil
	}
	return nil
}
