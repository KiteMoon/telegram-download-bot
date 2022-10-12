package MessageModle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func SendMessage(chatId int, text, token string) {
	sendMessageUrl := "https://api.telegram.org/bot" + token + "/sendMessage"
	SendMessageReq, _ := http.NewRequest("GET", sendMessageUrl, nil)
	SendMessageReq.Header.Add("content-type", "multipart/form-data; boundary=---011000010111000001101001")
	SendMessageReqQuery := SendMessageReq.URL.Query()
	SendMessageReqQuery.Add("text", text)
	SendMessageReqQuery.Add("chat_id", strconv.Itoa(chatId))
	SendMessageReq.URL.RawQuery = SendMessageReqQuery.Encode()
	SendMessageRes, _ := http.DefaultClient.Do(SendMessageReq)

	defer SendMessageRes.Body.Close()
	body, _ := ioutil.ReadAll(SendMessageRes.Body)

	fmt.Println(SendMessageRes)
	fmt.Println(string(body))
}
