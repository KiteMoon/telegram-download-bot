package PareMessage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"telegrame-kitemoon-bot/DownModle"
)

func PareMessageList(tgToken string) string {
	messageList, messageListErr := http.Get("https://api.telegram.org/bot" + tgToken + "/getUpdates")
	if messageListErr != nil {
		fmt.Println("请求消息列表失败")
		fmt.Println(messageListErr)
		return "ERROR"
	}
	defer messageList.Body.Close()
	messageListIO, messageListIOErr := ioutil.ReadAll(messageList.Body)
	if messageListIOErr != nil {
		fmt.Println("解析消息列表失败")
		fmt.Println(messageListErr)
		return "ERROR"
	}
	//fmt.Println(string(messageListIO))
	x := new(PareMessageStruct)
	err := json.Unmarshal(messageListIO, x)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for i := 0; i < len(x.Result); i++ {
		fmt.Println("------")
		fmt.Println("来源ID:", x.Result[i].Message.From.Id)
		fmt.Println("消息ID:", x.Result[i].Message.Chat.Id)
		fmt.Println("消息发送人:", x.Result[i].Message.From.FirstName)
		fmt.Println("上传时间:", x.Result[i].UpdateId)
		if x.Result[i].Message.Document.MimeType != "" {
			fmt.Println(x.Result[i].Message.Document.MimeType)
			fmt.Println(x.Result[i].Message.Document.FileName)
			fmt.Println(x.Result[i].Message.Document.FileId)
			fmt.Println(x.Result[i].Message.Document.FileSize)
			miniPareData := PareMessageMiniStruct{
				FileName:  x.Result[i].Message.Document.FileName,
				MimeType:  x.Result[i].Message.Document.MimeType,
				FileId:    x.Result[i].Message.Document.FileId,
				ChatID:    x.Result[i].Message.Chat.Id,
				FileSize:  x.Result[i].Message.Document.FileSize,
				MessageID: x.Result[i].Message.MessageId,
			}
			DownModle.PareTGDocImg(DownModle.PareMessageMiniStruct(miniPareData), tgToken)
		} else {
			fmt.Println("不存在字段")
		}

	}

	return string(messageListIO)
}
