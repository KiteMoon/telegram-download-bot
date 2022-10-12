package main

type GetBotStatusStruct struct {
	Ok     bool `json:"ok,omitempty"`
	Result struct {
		Id         int    `json:"id,omitempty"`
		Is_bot     bool   `json:"is_bot,omitempty"`
		First_name string `json:"first_name,omitempty"`
		Username   string `json:"username,omitempty"`
	} `json:"result"`
}
