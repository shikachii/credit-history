package lib

import (
	"fmt"
	"regexp"
)

type CreditHistory struct {
	Date string
	Shop string
	Amount string
	Transaction string
	Card string
}

func Parse(mail string) (*CreditHistory, error) {
	ch := CreditHistory{}

	exp, err := regexp.Compile(`◇利用日：(.*)\n`)
	
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if date := exp.FindStringSubmatch(mail); len(date) != 0 {
		fmt.Println(date[1])
		ch.Date = date[1]
	} else {
		ch.Date = mail
		return &ch, fmt.Errorf("date not found")
	}

	exp, err = regexp.Compile(`◇利用先：(.*)`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	shop := exp.FindStringSubmatch(mail)[1]
	fmt.Println(shop)
	ch.Shop = shop

	exp, err = regexp.Compile(`◇利用金額：(.*)`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	amount := exp.FindStringSubmatch(mail)[1]
	fmt.Println(amount)
	ch.Amount = amount

	exp, err = regexp.Compile(`◇利用取引：(.*)`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	transaction := exp.FindStringSubmatch(mail)[1]
	fmt.Println(transaction)
	ch.Transaction = transaction

	exp, err = regexp.Compile(`ご利用カード：(.*)`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	card := exp.FindStringSubmatch(mail)[1]
	fmt.Println(card)
	ch.Card = card

	return &ch, nil
}