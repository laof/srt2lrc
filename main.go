package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path"
	"srt2lrc/translate"
	"strconv"
	"strings"
	"time"
)

var i18n = false

func main() {

	args := len(os.Args)

	if args == 3 {
		i18n = os.Args[2] == "y"
	}

	if args >= 2 {
		create(os.Args[1])
	} else {
		fmt.Println("please enter file info")
	}
}

func create(filename string) {
	okbyte, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	list := []LRC{}

	txt := strings.Split(string(okbyte), "\n\r")
	for _, str := range txt {
		obj := coverTime(str)
		if obj.Time == "" || obj.Subtitle == "" {
			continue
		}
		list = append(list, obj)
	}

	appLength := len(translate.Apps)
	dataLength := len(list)

	app, _ := strconv.ParseFloat(strconv.Itoa(appLength), 64)
	data, _ := strconv.ParseFloat(strconv.Itoa(dataLength), 64)

	total := int(math.Ceil(data / app))

	res := []LRC{}
	for i := 0; i < total; i++ {

		if i18n {
			time.Sleep(1 * time.Second)
		}

		start := i * appLength
		end := start + appLength

		var arr []LRC

		if end > dataLength {
			arr = list[start:]
		} else {
			arr = list[start:end]
		}

		for zhi, obj := range arr {

			if i18n {
				obj.Translation = translate.Translator(obj.Subtitle, zhi)
				fmt.Printf("(%v / %v) [%v] => [%v]\n", start+zhi+1, dataLength, obj.Subtitle, obj.Translation)
			}

			res = append(res, obj)

		}

	}

	txtbyte, _ := json.Marshal(res)

	fs := strings.TrimSuffix(filename, path.Ext(filename)) + ".json"
	os.WriteFile(fs, txtbyte, os.ModePerm)
	fmt.Println(fs + " done!")
}

type LRC struct {
	Translation string `json:"i18n"`
	Subtitle    string `json:"s"`
	Time        string `json:"t"`
}

/**
1
00:00:16,766 --> 00:00:18,066
reporting work
*/

func coverTime(session string) LRC {

	obj := LRC{}

	info := strings.Split(session, "\r")
	if len(info) < 3 {
		return obj
	}

	// 00:00:16,766 --> 00:00:18,066
	start := strings.Split(info[1], "-->")

	if len(start) == 2 {
		timeline := strings.TrimSpace(start[0])
		// 00:00:16,766
		detail := strings.Split(timeline, ",")
		// 00:00:16
		arr := strings.Split(detail[0], ":")

		if len(arr) == 3 {

			h, _ := strconv.Atoi(arr[0])
			m, _ := strconv.Atoi(arr[1])
			seconds := arr[2]

			m = h*60 + m

			minutes := strconv.Itoa(m)

			if m < 10 {
				minutes = "0" + strconv.Itoa(m)
			}

			obj.Time = fmt.Sprintf("%s:%s.%s", minutes, seconds, detail[1])
		} else {
			obj.Time = timeline
		}

		obj.Subtitle = strings.TrimLeft(info[2], "\n\r")
		return obj
	}

	return obj
}
