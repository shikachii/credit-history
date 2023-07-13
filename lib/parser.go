package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/shikachii/credit-history/domain/model"
)

func str2time(str string) time.Time {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	parsedTime, _ := time.ParseInLocation("2006/01/02 15:04", str, loc)
	return parsedTime
}

func time2unix(t time.Time) int64 {
	return t.Unix()
}

func Parse(mail string) (*model.CreditHistory, error) {
	ch := model.CreditHistory{}
	mail = strings.ReplaceAll(mail, "\r", "")
	mail = strings.ReplaceAll(mail, "\n", "\n ")

	exp, err := regexp.Compile(`◇利用日：(.*)\n`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if date := exp.FindStringSubmatch(mail); len(date) != 0 {
		fmt.Println(date[1])
		ch.Date = date[1]
		unixtime := time2unix(str2time(ch.Date))
		ch.Timestamp = strconv.FormatInt(unixtime, 10)
	} else {
		return nil, fmt.Errorf("date not found")
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
	amount = strings.Split(amount, "円")[0]
	ch.Amount = strings.ReplaceAll(amount, ",", "")

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