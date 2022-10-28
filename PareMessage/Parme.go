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

type FirstMessageUpdateID struct {
	Update_id int64 `bind:"update_id" binding:"require"`
	Message   struct {
		MessageId int `json:"message_id" binding:"require"`
		From      struct {
			Id           int    `json:"id" binding:"require"d`
			IsBot        bool   `json:"is_bot" binding:"require"`
			FirstName    string `json:"first_name" binding:"require"`
			LastName     string `json:"last_name" binding:"require"`
			Username     string `json:"username" binding:"require"`
			LanguageCode string `json:"language_code" binding:"require"`
			IsPremium    bool   `json:"is_premium" binding:"require"`
		} `json:"from" bind:"message"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
	}
}
