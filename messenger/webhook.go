package messenger

type Webhook struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID        string      `json:"id"`
	Time      int         `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
	Sender    `json:"sender"`
	Recipient `json:"recipient"`
	Timestamp int `json:"timestamp"`
	*Message  `json:"message,omitempty"`
	*Postback `json:"postback,omitempty"`
}

type Sender struct {
	ID string `json:"id"`
}

type Message struct {
	Mid           string `json:"mid,omitempty"`
	Text          string `json:"text,omitempty"`
	*QuickReply   `json:"quick_reply,omitempty"`
	*Attachment   `json:"attachment,omitempty"`
	*QuickReplies `json:"quick_replies,omitempty"`
}

type QuickReplies []QuickReply

type QuickReply struct {
	ContentType string `json:"content_type,omitempty"`
	Title       string `json:"title,omitempty"`
	Payload     string `json:"payload,omitempty"`
}

type Postback struct {
	Payload string `json:"payload"`
}

type Referral struct {
	Ref    string `json:"ref"`
	Source string `json:"source"`
	Type   string `json:"type"`
}
