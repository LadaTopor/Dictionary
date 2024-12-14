package search

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) SearchTitle(title string) ([]Information, error) {
	var information []Information
	result, err := r.db.Query(`SELECT id, title, translation
									 FROM ru_en 
									 WHERE word_similarity(title, $1) > 0.0 
									 ORDER BY word_similarity(title, $1) DESC 
									 LIMIT 100`, title)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var info Information
		err = result.Scan(&info.Id, &info.Title, &info.Translation)
		if err != nil {
			return nil, err
		}
		information = append(information, info)
	}

	return information, nil
}
