package feishu_webhook

import (
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"strings"
)

func SendMessge() {
	url := `https://open.feishu.cn/open-apis/bot/v2/hook/71f4effc-4c76-484a-9e3a-9b6286ac4a79`
	sendMsg(url, "fasfasf")
}

func sendMsg(apiUrl, msg string) {
	// json
	contentType := "application/json"
	// data
	card := fmt.Sprintf(`{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "template": "%s",
    "title": {
      "content": "%s",
      "tag": "plain_text"
    }
  },
  "i18n_elements": {
    "zh_cn": [
      {
        "tag": "hr"
      },
      {
        "tag": "markdown",
        "content": "**命名空间：**%s\n**实例ID：**%s\n**实例名称：**%s\n**实例状态：**%s\n"
      },
      {
        "tag": "action",
        "actions": [
          {
            "tag": "button",
            "text": {
              "tag": "plain_text",
              "content": "查看详情"
            },
            "type": "primary",
            "url": "%s"
          }
        ]
      }
    ]
  }
}`, "red", "标题", "标题", "标题", "标题", "标题", "标题")
	fmt.Println("ooooooooooooooooooooo", card)
	sendData := `{
			"card": ` + card + `,
			"msg_type": "interactive"
	}`
	// request
	result, err := http.Post(apiUrl, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	fmt.Println("hhhhhhhh", result.StatusCode)
	all, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println("22", err)
	}

	type resp struct {
		Code    int    `json:"StatusCode"`
		Message string `json:"StatusMessage"`
	}

	responseResult := new(resp)
	err = json.Unmarshal(all, &responseResult)
	if err != nil {
		fmt.Println("000000", err)
	}
	fmt.Println("---------------", string((all)))
	fmt.Println("---------------", fmt.Sprintf("%+v", responseResult))
	defer result.Body.Close()

}
