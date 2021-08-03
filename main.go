package main

import (
	"log"
	"github.com/Filmptz/whalealert_linebot/transactions/svc"
)

func main() {
	if err := svc.TrickerFetchAPI();err != nil {
		log.Fatalln(err)
	}
}