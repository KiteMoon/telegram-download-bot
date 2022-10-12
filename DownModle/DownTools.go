package DownModle

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"telegrame-kitemoon-bot/MessageModle"
)

func PareTGDocImg(miniPareData PareMessageMiniStruct, tgToken string) {
	// 预下载URL(这里可能会挂，取决于图片下载速度。
	pareDownUrlApi := "https://api.telegram.org/bot" + tgToken + "/getFile"
	// 拉起请求
	pareDownUrlApiReq, _ := http.NewRequest("GET", pareDownUrlApi, nil)
	// 添加file_id下载参数
	pareDownUrlApiReqQuery := pareDownUrlApiReq.URL.Query()
	pareDownUrlApiReqQuery.Add("file_id", miniPareData.FileId)
	pareDownUrlApiReq.URL.RawQuery = pareDownUrlApiReqQuery.Encode()
	// 发起
	res, _ := http.DefaultClient.Do(pareDownUrlApiReq)
	defer res.Body.Close()
	// 解析读取值
	body, _ := ioutil.ReadAll(res.Body)
	reqBody := new(PareTGDocImgReqStuct)
	pareReqJsonErr := json.Unmarshal(body, reqBody)
	fmt.Println(string(body))
	if pareReqJsonErr != nil || reqBody.Ok == false {
		fmt.Println("解析失败，未知BUG")
		fmt.Println(pareReqJsonErr)
		return
	}
	fmt.Println(reqBody.Result.FileSize)
	fmt.Println(reqBody.Result.FilePath)
	fmt.Println(reqBody)
	DownTGImg(miniPareData, reqBody.Result.FilePath, tgToken)
}

func DownTGImg(miniPareData PareMessageMiniStruct, filePath string, tgToken string) {
	if miniPareData.FileId == "" {
		fmt.Println("检测到空文件传入，忽略")
		return
	}
	downUrl := "https://api.telegram.org/file/bot" + tgToken + "/" + filePath
	fmt.Println(downUrl)
	downrequests := http.Client{}
	tgRequest, _ := http.NewRequest("GET", downUrl, nil)
	response, err := downrequests.Do(tgRequest)
	if err != nil {
		fmt.Println("下载失败")
		fmt.Println(err.Error())
		return
	}
	saveFilePath := fmt.Sprintf("download/%d/%s", miniPareData.MessageID, miniPareData.FileName)
	fmt.Println("下载文件名", miniPareData.MessageID)
	fmt.Println("下载文件名", miniPareData.FileName)
	fmt.Println("下载文件名", saveFilePath)
	if response.StatusCode == 200 {
		mkDirerr := os.MkdirAll(fmt.Sprintf("download/%d/", miniPareData.MessageID), os.ModePerm)
		if mkDirerr != nil {
			fmt.Println("发送错误，创建文件夹失败")
			fmt.Println(mkDirerr)
			return
		}
		file, err := os.Create(saveFilePath)
		if err != nil {
			fmt.Println("创建数据失败")
			fmt.Println(err)
			return

		}
		defer file.Close()
		n, err := io.Copy(file, response.Body)
		if err != nil {
			fmt.Println("复制数据失败")
			fmt.Println(err)
			return
		}
		sendMessageStr := fmt.Sprintf("---远程下载机器人---\n已经完成下载任务\n下载文件ID:%s\n下载文件名:%s\n下载文件类型:%s\n下载发起人ID:%d\n理论下载文件大小:%dKB\n时间下载文件大小:%dKB\n下载文件消息ID:%d", miniPareData.FileId, miniPareData.FileName, miniPareData.MimeType, miniPareData.ChatID, miniPareData.FileSize/1024, n/1024, miniPareData.MessageID)
		MessageModle.SendMessage(719971136, sendMessageStr, tgToken)

	}
}
