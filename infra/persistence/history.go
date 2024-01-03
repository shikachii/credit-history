package persistence

import (
	"github.com/shikachii/credit-history/domain/model"
	"github.com/shikachii/credit-history/domain/repository"
)

type HistoryPersistence struct {}

func NewHistoryPersistence() repository.HistoryRepository {
	return &HistoryPersistence{}
}

func (hp *HistoryPersistence) Insert(*model.CreditHistory) error {
	return nil
}