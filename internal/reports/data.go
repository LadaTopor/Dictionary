package reports

import "database/sql"

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// RGetReportById ищем репорт по id
func (r *Repo) RGetReportById(id int) (*Report, error) {
	var report Report
	err := r.db.QueryRow(`SELECT id, title, description, created_at, updated_at FROM reports WHERE id = $1`, id).
		Scan(&report.Id, &report.Title, &report.Description, &report.Created_at, &report.Updated_at)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// CreateNewReport добавляет новый репорт в базу даных
func (r *Repo) CreateNewReport(report, description string) error {
	_, err := r.db.Exec(`INSERT INTO reports (title, description) VALUES ($1, $2)`, report, description)
	if err != nil {
		return err
	}

	return nil
}

// EditReport изменяет репорт в базе данных или добавляет новый по id если такого не существует
func (r *Repo) EditReport(report, description string, id int) error {
	//вариант без создания нового элемента при обращении к несуществующему айди:
	//	_, err := r.db.Exec(`UPDATE reports SET title = $2, description = $3 WHERE id = $1;`, id, word, translate)
	_, err := r.db.Exec(`INSERT INTO reports (id, title, description) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET title = $2, description = $3;`, id, report, description)
	if err != nil {
		return err
	}

	return nil
}

// DeleteReport удаляет репорты из базы данных по id
func (r *Repo) DeleteReport(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1;`, id)
	if err != nil {
		return err
	}
	return nil
}
