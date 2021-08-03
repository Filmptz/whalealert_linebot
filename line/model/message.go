package model

type (
	LineMessage struct {
		Destination string `json:"destination"`
		Events      []struct {
			ReplyToken string `json:"replyToken"`
			Type       string `json:"type"`
			Timestamp  int64  `json:"timestamp"`
			Source     struct {
				Type   string `json:"type"`
				UserID string `json:"userId"`
			} `json:"source"`
			Message struct {
				ID   string `json:"id"`
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"events"`
	}

	ReplyMessage struct {
		ReplyToken string `json:"replyToken"`
		Messages   []Text `json:"messages"`
	}

	BroadcastMessage struct {
		Messages []Text `json:"messages"`
	}

	Text struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}
)
