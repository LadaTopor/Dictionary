package words

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// RGetWordById ищем слово по id
func (r *Repo) RGetWordById(id int) (*Word, error) {
	var word Word
	err := r.db.QueryRow(`SELECT id, title, translation FROM ru_en WHERE id = $1`, id).
		Scan(&word.Id, &word.Title, &word.Translation)
	if err != nil {
		return nil, err
	}

	return &word, nil
}

// CreateNewWords добавляет новые переводы в базу даных
func (r *Repo) CreateNewWords(word, translate string) error {
	_, err := r.db.Exec(`INSERT INTO ru_en (title, translation) VALUES ($1, $2)`, word, translate)
	if err != nil {
		return err
	}

	return nil
}

// UpdateWord обновляет существующее слово в базе данных
func (r *Repo) UpdateWord(id int, word, translate string) error {
	_, err := r.db.Exec(`UPDATE ru_en SET title = $1, translation = $2 WHERE id = $3`, word, translate, id)

	if err != nil {
		return err
	}

	return nil
}

// DeleteWord удаляет существующее слово в базе данных
func (r *Repo) DeleteWord(id int) error {
	_, err := r.db.Exec(`DELETE FROM ru_en WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

// SearchWords выдаёт ранжируемый список совпадений с переданным title
func (r *Repo) SearchWords(title string) ([]Word, error) {
	rows, err := r.db.Query(`
        SELECT title, translation
        FROM ru_en
        ORDER BY similarity(title, $1) DESC
        LIMIT 100
    `, title)
	if err != nil {
		return nil, err
	}

	var words []Word
	for rows.Next() {
		var word Word
		err = rows.Scan(&word.Title, &word.Translation)
		if err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}
