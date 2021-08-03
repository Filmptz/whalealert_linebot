package svc

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/Filmptz/whalealert_bot/transactions/serializer"
	"github.com/Filmptz/whalealert_bot/line/model"
)

var prv_id string

func FormatResponse(Trans *serializer.TransactionResp) error {
	
	var flag string
	var id = Trans.Tranctions[0].ID
	var amountInt = int(Trans.Tranctions[0].Amount)
	var amountIntUsd = int(Trans.Tranctions[0].Amount_USD)

	for i:=0; i<amountIntUsd/10000000; i++ {
		flag = flag + string(emoji.PoliceCarLight)
	}
	if amountIntUsd/10000000 <= 0 {
		flag = string(emoji.Fire)
	}

	msg := fmt.Sprintf("%v \nAlert : %v \nAmount : %v %v /n%v usdt",flag,Trans.Tranctions[0].Symbol,amountInt,Trans.Tranctions[0].Symbol,amountIntUsd)
			
	if (amountInt != 0 && amountIntUsd != 0 && id != prv_id){
		prv_id  = id

		text := model.Text{
			Type : "text",
			Text : msg,
		}
				
		message := model.BroadcastMessage{
			Messages : []model.Text{
				text,
				},
			}
		BroadcastMessageLine(&message)
	}
	return nil
}