package route

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"

	"github.com/Filmptz/whalealert_bot/line/model"
	"github.com/Filmptz/whalealert_bot/line/svc"
)

func GetStatusOK() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(http.StatusOK, "ok")

	}
}

func UpdateMessages() gin.HandlerFunc{
	return func(c *gin.Context){
		Line := new(model.LineMessage)
		if err := c.Bind(Line); err != nil {
			log.Println(err)
			c.String(http.StatusOK, "error")
		}

		text := model.Text{
			Type : "text",
			Text : "ข้อความเข้ามา : " + Line.Events[0].Message.Text  + " ยินดีต้อนรับ : ",
		}
		
		message := model.ReplyMessage{
			ReplyToken : Line.Events[0].ReplyToken ,
			Messages : []model.Text{
				text,
			},
		}
		
		if err := svc.ReplyMessageLine(&message); err != nil {
			log.Println(err)
		}
		
		log.Println("%% message success")
		c.JSON(http.StatusOK, "ok")
	}	
}