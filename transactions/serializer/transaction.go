package serializer

import ()

type TransactionResp struct {
	Result     string `json:"result"`
	Cursor     string `json:"-"`
	Count      int    `json:"count"`
	Tranctions []struct {
		Blockchain       string `json:"blockchain"`
		Symbol           string `json:"symbol"`
		ID               string `json:"id"`
		Transaction_Type string `json:"-"`
		Hash             string `json:"-"`
		From             struct {
			Address    string `json:"-"`
			Owner_Type string `json:"owner_type"`
		} `json:"from"`
		To struct {
			Address    string `json:"-"`
			Owner_Type string `json:"owner_type"`
		} `json:"to"`
		Timestamp        int     `json:"timestamp"`
		Amount           float64 `json:"amount"`
		Amount_USD       float64 `json:"amount_usd"`
		TransactionCount int     `json:"-"`
	} `json:"transactions"`
}
