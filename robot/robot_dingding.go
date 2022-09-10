package robot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const dingding_webhook = "https://oapi.dingtalk.com/robot/send?access_token=3dda9daa1ba6936cd779aa9b63757b9b416a53ae25bb0a3f22e93cb78a490d7d"

type dingdingContent struct {
	Content string `json:"content"`
}

type dingdingData struct {
	Msgtype string          `json:"msgtype"`
	Text    dingdingContent `json:"text"`
}

func DispatchThresholdInfo(text string) {
	textSend := "CoinDog\n" + text

	client := &http.Client{}
	data := dingdingData{Msgtype: "text", Text: dingdingContent{Content: textSend}}
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", dingding_webhook, bytes.NewReader(bytesData))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
