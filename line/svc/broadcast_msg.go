package svc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Filmptz/whalealert_linebot/line/model"
	"github.com/Filmptz/whalealert_linebot/line/types"
	"github.com/Filmptz/whalealert_linebot/config/svc"
)

func BroadcastMessageLine(Message *model.BroadcastMessage) error {
	value, _ := json.Marshal(Message)

	url := types.LineBroadcastAPI

	chanelToken , err := svc.DotEnv("CHANEL_ACCESS_TOKEN")
	if err != nil {
		return err
	}

	var jsonStr = []byte(value)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ chanelToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return err
}