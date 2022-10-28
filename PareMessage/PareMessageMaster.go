package PareMessage

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"telegrame-kitemoon-bot/DownModle"
	"telegrame-kitemoon-bot/MessageModle"
)

func MessageType(context *gin.Context, tgToken string) {
	all, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		fmt.Println("解析出错了")
		return
	}
	fmt.Printf("%s\n", string(all))
	FirstMessageUpdateIDData := new(FirstMessageUpdateID)
	err = json.Unmarshal(all, &FirstMessageUpdateIDData)
	if err != nil {
		fmt.Println("解析json失败")
		fmt.Println(err)
		return
	}
	if FirstMessageUpdateIDData.Update_id == 0 || FirstMessageUpdateIDData.Message.MessageId == 0 && FirstMessageUpdateIDData.Message.From.Id == 0 {
		fmt.Println("不可识别的消息,字段不匹配")
		return
	}
	fmt.Println("发信人ID:", FirstMessageUpdateIDData.Message.From.Id)
	fmt.Println("发信人ID:", FirstMessageUpdateIDData.Message.Chat.Id)
	fmt.Println("发信人用户名:", FirstMessageUpdateIDData.Message.From.Username)
	fmt.Println("发信人使用语言:", FirstMessageUpdateIDData.Message.From.LanguageCode)
	fmt.Println("是否为会员:", FirstMessageUpdateIDData.Message.From.IsPremium)
	sendMessageStr := fmt.Sprintf("发信ID:<code>%d</code>\n发信用户名:%s\n发信人语言:%s\n发信人会员状态:%t", FirstMessageUpdateIDData.Message.From.Id,
		FirstMessageUpdateIDData.Message.From.Username, FirstMessageUpdateIDData.Message.From.LanguageCode,
		FirstMessageUpdateIDData.Message.From.IsPremium)
	MessageModle.SendMessage(FirstMessageUpdateIDData.Message.Chat.Id, sendMessageStr, tgToken)
}

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
