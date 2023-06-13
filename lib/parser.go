package lib

import (
	"fmt"
	"regexp"
)

func Parse() {
	mail := `Ｏｌｉｖｅ会員　様

	いつも三井住友カードをご利用頂きありがとうございます。
	お客様のカードご利用内容をお知らせいたします。
	
	ご利用カード：Ｏｌｉｖｅ／クレジット
	
	◇利用日：2023/06/12 20:28
	◇利用先：MYBASKET HONKOMAGOME3CHOU
	◇利用取引：買物
	◇利用金額：1,281円
	
	ご利用内容について、万が一身に覚えのない場合は、以下URLにてお問い合わせ先のご案内をしております。
	▼身に覚えのない明細でお困りの方
	https://www.smbc-card.com/mem/info/meisai_inquiry.jsp
	
	また、ご自身でカードのご利用を一時的に制限することが可能なあんしん利用制限サービスをご用意しております。
	▼あんしん利用制限サービスについて
	http://vpass.jp/usage2m/
	
	※カードご利用の承認照会があった場合に通知されるサービスであり、カードのご利用 及び ご請求を確定するものではありません。
	※ご利用店舗は、当社に売上の情報が到着後、Vpassのご利用明細照会やWEB明細で確認していただけます。反映までにお日にちがかかる場合がございます。
	※あとからリボ、あとから分割はご利用の内容がVpassのご利用明細照会やWEB明細に反映後、お申込みいただけるようになります。
	※携帯電話や公共料金などの継続的なご利用（注） 及び ETCやPiTaPa等、一部の電子マネー利用については通知されません。
	（注）利用内容によっては通知される可能性がございます。
	
	▼Oliveフレキシブルペイをお持ちのお客様
	本通知はクレジットモードのご利用のお知らせです。明細の確認や利用制限等の各種照会・設定はアプリからお手続きください。
	
	▼Vpassのログインはこちら
	https://www.smbc-card.com/mem/index.jsp
	
	▼ご利用通知サービスの詳細、設定変更・解除はこちら
	https://www.smbc-card.com/mem/service/sec/selfcontrol/usage_notice.jsp
	
	※このメールアドレスは送信専用です。ご返信に回答できません。
	
	
	■発行者■
	三井住友カード株式会社
	https://www.smbc-card.com/
	〒135-0061 東京都江東区豊洲2丁目2番31号 SMBC豊洲ビル`

	exp, err := regexp.Compile(`◇利用日：(.*)`)
	
	if err != nil {
		fmt.Println(err.Error())
	}
	date := exp.FindStringSubmatch(mail)[1]
	fmt.Println(date)

	exp, err = regexp.Compile(`◇利用先：(.*)`)
	if err != nil {
		fmt.Println(err.Error())
	}
	shop := exp.FindStringSubmatch(mail)[1]
	fmt.Println(shop)

	exp, err = regexp.Compile(`◇利用金額：(.*)`)
	if err != nil {
		fmt.Println(err.Error())
	}
	amount := exp.FindStringSubmatch(mail)[1]
	fmt.Println(amount)

	exp, err = regexp.Compile(`◇利用取引：(.*)`)
	if err != nil {
		fmt.Println(err.Error())
	}
	transaction := exp.FindStringSubmatch(mail)[1]
	fmt.Println(transaction)

	exp, err = regexp.Compile(`ご利用カード：(.*)`)
	if err != nil {
		fmt.Println(err.Error())
	}
	card := exp.FindStringSubmatch(mail)[1]
	fmt.Println(card)

}