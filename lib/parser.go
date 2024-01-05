package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

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

func ParseHtml(htmlMail string) (*model.CreditHistory, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlMail))
	if err != nil {
		return nil, err
	}

	ch := model.CreditHistory{}

	query := "html body table tr td table tr td table tbody tr td table tbody tr td table tbody tr td table tbody tr td"
	doc.Find(query).Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		trimmedText := strings.TrimSpace(text)
		replacedText := strings.ReplaceAll(trimmedText, "\n", "")

		// Date, Timestamp
		if i == 0 {
			ch.Date = strings.Replace(replacedText, "ご利用日時：", "", 1)
			unixtime := time2unix(str2time(ch.Date))
			ch.Timestamp = strconv.FormatInt(unixtime, 10)
		}

		// Shop, Transaction
		if i == 1 {
			ch.Shop = strings.Split(replacedText, "（")[0]
			ch.Transaction = strings.Replace(strings.Split(replacedText, "（")[1], "）", "", 1)
		}

		// Amount
		if i == 2 {
			ch.Amount = strings.ReplaceAll(strings.Split(replacedText, "円")[0], ",", "")
		}
	})

	cardQuery := "html body table tr td table tr td"
	doc.Find(cardQuery).Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		trimmedText := strings.TrimSpace(text)
		replacedText := strings.ReplaceAll(trimmedText, "\n", "")

		// Card
		if i == 6 {
			exp, err := regexp.Compile(`ございます。(.*)について`)
			if err != nil {
				fmt.Println(err.Error())
			}
			card := exp.FindStringSubmatch(replacedText)[1]
			ch.Card = card
		}
	})
	/**
	0: ご利用日時：2024/01/02 10:21,
	1: Dominos Pizza Japan Inc（買物）,
	2: 5,747円,
	3: ,
	4: Vpassアプリ,
	5: 生体認証で素早く安全にログイン,
	6: LINE公式アカウント,
	7: 簡単にお支払い・ポイント確認,
	*/

	/*
	html>body>table>tr>td>table>tr>td>table>tbody>tr>td>table>tbody>tr>td>table>tbody>tr>td>table>tbody>tr>td
	*/

	fmt.Println(doc)
	
	// ch.Date = "2024/01/02 10:21"
	// ch.Shop = "Dominos Pizza Japan Inc"
	// ch.Amount = "5747"
	// ch.Timestamp = "1704158460"
	// ch.Card = "Ｏｌｉｖｅ／クレジット"
	// ch.Transaction = "買物"

	return &ch, nil
}