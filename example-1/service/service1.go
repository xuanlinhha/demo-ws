package service

import (
	"demo-ws/example-1/repo"

	"github.com/pkg/errors"
)

type Service1 interface {
	GetRowsNo() (int, error)
}
type service1 struct {
	Repo1 repo.Repo1
}

func NewService1(repo1 repo.Repo1) Service1 {
	return &service1{Repo1: repo1}
}

func (r *service1) GetRowsNo() (int, error) {
	count, err := r.Repo1.CountRows()
	if err != nil {
		return -1, errors.Wrap(err, "CountRows")
	} else {
		return count, nil
	}
}
