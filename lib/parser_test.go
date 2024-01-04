package lib

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/shikachii/credit-history/domain/model"
)


func TestParseHtml(t *testing.T) {
	test_data := [...]string{
		"../testdata/domino_mail.html",
		"../testdata/jr_mail.html",
	}

	mailPaths := []io.Reader{}

	for _, htmlPath := range test_data {
		f, err := os.Open(htmlPath)
		if err != nil {
			t.Errorf("failed to open file: %v", err)
		}
		defer f.Close()
	
		b, err := io.ReadAll(f)
		if err != nil {
			t.Errorf("failed to read file: %v", err)
		}
	
		htmlReader := strings.NewReader(string(b))
		mailPaths = append(mailPaths, htmlReader)
	}
	
	type args struct {
		mail io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *model.CreditHistory
		wantErr bool
	}{
		{
			name: "Domino Pizza Test",
			args: args{
				mail: mailPaths[0],
			},
			want: &model.CreditHistory{
				Date:        "2024/01/02 10:21",
				Shop:        "Dominos Pizza Japan Inc",
				Amount:      "5747",
				Timestamp:   "1704158460",
				Card:        "Ｏｌｉｖｅ／クレジット",
				Transaction: "買物",
			},
			wantErr: false,
		},
		{
			name: "JR Test",
			args: args{
				mail: mailPaths[1],
			},
			want: &model.CreditHistory{
				Date:        "2023/12/27 23:21",
				Shop:        "JR EAST VIEW PLAZA",
				Amount:      "6780",
				Timestamp:   "1703686860",
				Card:        "Ｏｌｉｖｅ／クレジット",
				Transaction: "買物",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch, err := ParseHtml(tt.args.mail)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Printf("%+v\n", ch)

			if ch == nil {
				t.Errorf("Parse() ch = %v, want %v", ch, tt.want)
				return
			}

			if ch.Date != tt.want.Date {
				t.Errorf("Parse() Date = %v, want %v", ch.Date, tt.want.Date)
			}
			if ch.Timestamp != tt.want.Timestamp {
				t.Errorf("Parse() Timestamp = %v, want %v", ch.Timestamp, tt.want.Timestamp)
			}
			if ch.Shop != tt.want.Shop {
				t.Errorf("Parse() Shop = %v, want %v", ch.Shop, tt.want.Shop)
			}
			if ch.Amount != tt.want.Amount {
				t.Errorf("Parse() Amount = %v, want %v", ch.Amount, tt.want.Amount)
			}
			if ch.Card != tt.want.Card {
				t.Errorf("Parse() Card = %v, want %v", ch.Card, tt.want.Card)
			}
			if ch.Transaction != tt.want.Transaction {
				t.Errorf("Parse() Transaction = %v, want %v", ch.Transaction, tt.want.Transaction)
			}
		})
	}
	
}