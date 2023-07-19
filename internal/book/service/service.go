package service

import (
	"go-rest-api/internal/book/domain"
)

func CreateBook(
	uow domain.UnitOfWorkInterface,
	book *domain.Book,
) error {
	tx, err := uow.Begin()

	if err != nil {
		return err
	}

	err = uow.GetBookRepo().CreateBook(tx, book)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetBookById(uow domain.UnitOfWorkInterface, id string) (*domain.Book, error) {
	tx, err := uow.Begin()

	if err != nil {
		return nil, err
	}

	return uow.GetBookRepo().GetById(tx, id)
}
