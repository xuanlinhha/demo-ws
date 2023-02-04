package repo

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type Repo1 interface {
	CountRows() (int, error)
}

type repo1 struct {
	DB *sql.DB
}

func NewRepo1(db *sql.DB) Repo1 {
	return &repo1{DB: db}
}
func (r *repo1) CountRows() (int, error) {
	countStmt, err := r.DB.PrepareContext(context.TODO(), CountTable1)
	if err != nil {
		return -1, errors.Wrap(err, "PrepareContext")
	}
	defer countStmt.Close()
	var count int
	if err = countStmt.QueryRow().Scan(&count); err != nil {
		return -1, errors.Wrap(err, "QueryRow")
	} else {
		return count, nil
	}
}
