package words

import "database/sql"

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

func (r *Repo) UpdateWordById(id int, newTitle, newTranslation string) error {
	_, err := r.db.Exec(`UPDATE ru_en SET title = $1, translation = $2 WHERE id = $3`, newTitle, newTranslation, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteWordById(id int) error {
	_, err := r.db.Exec(`DELETE FROM ru_en WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) SearchWords(query string) ([]Word, error) {
	rows, err := r.db.Query(`SELECT id, title, translation FROM ru_en WHERE $1 <% title ORDER BY title <-> $1 LIMIT 100 `, query)
	if err != nil {
		return nil, err
	}

	var words []Word
	for rows.Next() {
		var word Word
		err := rows.Scan(&word.Id, &word.Title, &word.Translation)
		if err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}
