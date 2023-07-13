package model

type CreditHistory struct {
	Date string
	Shop string
	Amount string
	Transaction string
	Card string
	Timestamp string
}

// func newCreditHistory(
// 	date, shop, amount, transaction, card string,
// ) *CreditHistory {
// 	return &CreditHistory{
// 		Date: date,
// 		Shop: shop,
// 		Amount: amount,
// 		Transaction: transaction,
// 		Card: card,
// 	}
// }