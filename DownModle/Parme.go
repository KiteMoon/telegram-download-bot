package DownModle

type PareTGDocImgReqStuct struct {
	Ok     bool `json:"ok"`
	Result struct {
		FileId       string `json:"file_id"`
		FileUniqueId string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		FilePath     string `json:"file_path"`
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
