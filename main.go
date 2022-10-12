package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"telegrame-kitemoon-bot/PareMessage"
)

func main() {
	// API BOT token
	tgToken := "xxx"
	getBotStatus, getBotStatusErr := http.Get("https://api.telegram.org/bot" + tgToken + "/getMe")
	if getBotStatusErr != nil || getBotStatus.StatusCode != 200 {
		fmt.Println("发生错误，无法连接到TG服务器")
		panic("ERROR")
	}
	defer getBotStatus.Body.Close()
	getBotStatusBody, getBotStatusBodyErr := ioutil.ReadAll(getBotStatus.Body)
	if getBotStatusBodyErr != nil {
		fmt.Println("解析实体失败")
	}
	fmt.Println(string(getBotStatusBody))
	getBotStatusBodyJson := new(GetBotStatusStruct)
	parsingErr := json.Unmarshal(getBotStatusBody, getBotStatusBodyJson)
	if parsingErr != nil {
		fmt.Println("解析json失败")
		fmt.Println(parsingErr)
		return
	}
	fmt.Println("机器人ID:", getBotStatusBodyJson.Result.Id)
	fmt.Println("机器人昵称:", getBotStatusBodyJson.Result.First_name)
	fmt.Println("机器人用户名:", getBotStatusBodyJson.Result.Username)
	fmt.Println("初始化成功")
	PareMessage.PareMessageList(tgToken)
}
