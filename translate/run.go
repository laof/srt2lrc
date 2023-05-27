package translate

import (
	"encoding/json"
	"fmt"
	"srt2lrc/translate/utils"
	"srt2lrc/translate/utils/authv3"
)

// 您的应用ID
var appKey = "2ccac2276928012f"

// 您的应用密钥
var appSecret = "tm0aC9BVe2qZq4DuHRR9p5KdEA7y6l1Y"

func Translator(txt string) string {
	// 添加请求参数
	paramsMap := createRequestParams(txt)
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv3.AddAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	result := utils.DoPost("https://openapi.youdao.com/api", header, paramsMap, "application/json")
	// 打印返回结果
	if result != nil {
		data := ResultData{}
		ok := json.Unmarshal(result, &data)
		if ok == nil && len(data.Translation) > 0 {
			return data.Translation[0]
		}

		fmt.Println(txt)
		fmt.Println(string(result))
	}
	return ""
}

func createRequestParams(txt string) map[string][]string {

	/*
		note: 将下列变量替换为需要请求的参数
		取值参考文档: https://ai.youdao.com/DOCSIRMA/html/%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E7%BF%BB%E8%AF%91/API%E6%96%87%E6%A1%A3/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1-API%E6%96%87%E6%A1%A3.html
	*/
	q := txt
	from := "en"
	to := "zh-CHS"
	vocabId := ""

	return map[string][]string{
		"q":       {q},
		"from":    {from},
		"to":      {to},
		"vocabId": {vocabId},
	}
}
