package PareMessage

type PareMessageStruct struct {
	Ok     bool `json:"ok,omitempty"`
	Result []struct {
		UpdateId int64 `json:"update_id" :"update_id"`
		Message  struct {
			From struct {
				Id           int64  `json:"id" :"id"`
				IsBot        bool   `json:"is_bot" :"is_bot"`
				FirstName    string `json:"first_name" :"first_name"`
				LanguageCode string `json:"language_code" :"language_code"`
			} `json:"from" :"from"`
			Chat struct {
				Id        int    `json:"id"`
				FirstName string `json:"first_name"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date      int64
			Text      string
			MessageId int `json:"message_id"`

			photo    interface{}
			Document struct {
				MimeType     string `json:"mime_type"`
				FileName     string `json:"file_name"`
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				FileSize     int    `json:"file_size"`
			}
		} `:"message"`
	} `json:"result"`
}

type PareMessageMiniStruct struct {
	FileName  string
	MimeType  string
	FileId    string
	ChatID    int
	FileSize  int
	MessageID int
}
