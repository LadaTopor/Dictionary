package words

import (
	"database/sql"
	"errors"
	"strconv"
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

func (r *Repo) UpdateNewWord(id, word, translate string) error {
	idInt, _ := strconv.Atoi(id)
	result, err := r.db.Exec(`UPDATE ru_en SET title = $2, translation = $3 WHERE id = $1`, idInt, word, translate)
	if err != nil {
		return err
	}

	if aff, err := result.RowsAffected(); err != nil || aff != 1 {
		return errors.New("unexpected id")
	}

	return nil
}

func (r *Repo) DeleteWordById(id int) error {
	result, err := r.db.Exec(`DELETE FROM ru_en WHERE id = $1`, id)
	if err != nil {
		return err
	}

	if aff, err := result.RowsAffected(); err != nil || aff != 1 {
		return errors.New("no such id")
	}

	return nil
}
