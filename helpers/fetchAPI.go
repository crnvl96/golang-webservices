package helpers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Quotation struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type USDToBRL struct {
	USDBRL Quotation `json:"usdbrl"`
}

func FetchAPI(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		"https://economia.awesomeapi.com.br/json/last/USD-BRL",
		nil,
	)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	parsed, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var data USDToBRL
	err = json.Unmarshal(parsed, &data)
	if err != nil {
		return "", err
	}

	return data.USDBRL.Bid, nil
}
