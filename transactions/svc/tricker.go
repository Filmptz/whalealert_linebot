package svc

import (

	"io/ioutil"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"fmt"

	"github.com/Filmptz/whalealert_linebot/config/svc"
	 lineSvc "github.com/Filmptz/whalealert_linebot/line/svc"
	"github.com/Filmptz/whalealert_linebot/transactions/serializer"
	"strings"
)

const (
	duration = 6*time.Second
)

func TrickerFetchAPI() error {
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			apiKey , err := svc.DotEnv("API_KEY")
			if err != nil {
				fmt.Println("error:", err) 
			}

			minValue := 1000000
			start := time.Now().Unix() - 500

			url := fmt.Sprintf("https://api.whale-alert.io/v1/transactions?api_key=%v&min_value=%v&start=%v&limit=1",apiKey,minValue,start)

			req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(``))
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			log.Println("response Status:", resp.Status)
			body, _ := ioutil.ReadAll(resp.Body)
			

			var TransactionReturn serializer.TransactionResp
			if err := json.Unmarshal(body, &TransactionReturn); err != nil {
				return err
			}

			if err := lineSvc.FormatResponse(&TransactionReturn); err !=nil {
				return err
			}
			ticker.Reset(duration)
		}
	}
}