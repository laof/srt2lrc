package translate

import (
	"encoding/json"
	"fmt"
	"srt2lrc/translate/utils"
	"srt2lrc/translate/utils/authv3"
)

func T1(txt string) string {

	// 您的应用ID
	appKey := "2ccac2276928012f"
	// 您的应用密钥
	appSecret := "tm0aC9BVe2qZq4DuHRR9p5KdEA7y6l1Y"

	return translator(txt, appKey, appSecret)
}

func T2(txt string) string {

	// 您的应用ID
	mkey := "127b8fcd5dc1eb9e"
	// 您的应用密钥
	appSecret := "Tzl7WkRG9Nzwdp9ew0GWCjIsqJ4XVRlv"

	return translator(txt, mkey, appSecret)
}

func translator(txt string, appKey string, appSecret string) string {
	// 添加请求参数
	paramsMap := createRequestParams(txt)
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv3.AddAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	res := utils.DoPost("https://openapi.youdao.com/api", header, paramsMap, "application/json")
	// 打印返回结果
	if res != nil {
		result := Result{}
		ok := json.Unmarshal(res, &result)
		if ok == nil && len(result.Translation) > 0 {
			return result.Translation[0]
		}

		fmt.Println(txt)
		fmt.Println(string(res))
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
