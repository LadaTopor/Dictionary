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

// EditWord изменяет перевод в базе данных или добавляет новый по id если такого не существует
func (r *Repo) EditWord(word, translate string, id int) error {
	//вариант без создания нового элемента при обращении к несуществующему айди:
	//	_, err := r.db.Exec(`UPDATE ru_en SET title = $2, translation = $3 WHERE id = $1;`, id, word, translate)
	_, err := r.db.Exec(`INSERT INTO ru_en (id, title, translation) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET title = $2, translation = $3;`, id, word, translate)
	if err != nil {
		return err
	}

	return nil
}

// DeleteWord удаляет переводы из базы данных по id
func (r *Repo) DeleteWord(id int) error {
	_, err := r.db.Exec(`DELETE FROM ru_en WHERE id = $1;`, id)
	if err != nil {
		return err
	}
	return nil
}

//SearchWordsByKeyword ищет 100 самых похожих слов на keyword
func (r *Repo) SearchWordsByKeyword(keyword string) ([]Word, error) {
	var wordSlice []Word
	rows, err := r.db.Query(`SELECT id, title, translation
FROM ru_en
ORDER BY similarity(title, $1) DESC
LIMIT 100;`, keyword)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var word Word
		err = rows.Scan(&word.Id, &word.Title, &word.Translation)
		if err != nil {
			return nil, err
		}
		wordSlice = append(wordSlice, word)
	}
	return wordSlice, nil
}
