package repository

import "github.com/shikachii/credit-history/domain/model"

type HistoryRepository interface {
	Insert(*model.CreditHistory) error
}