package main

import (
	"fmt"
	"os"
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

	arr := []LRC{}

	txt := strings.Split(string(okbyte), "\n\r")
	for _, str := range txt {
		obj := coverTime(str)
		if obj.time == "" || obj.subtitle == "" {
			continue
		}
		arr = append(arr, obj)
	}

	lrc := []string{"[00:00.000] lrc : laof"}

	if i18n {
		lrc = append([]string{"[ml:1.0]"}, lrc...)
	}

	for i, current := range arr {

		lrc = append(lrc, fmt.Sprintf("%s %s", current.time, current.subtitle))

		if !i18n {
			continue
		}

		time.Sleep(1 * time.Second)
		zh := translate.Translator(current.subtitle)
		status := fmt.Sprintf("(%v/%v) [%v] => [%v]", i+1, len(arr), current.subtitle, zh)
		fmt.Println(status)
		nextTime := ""
		if i != len(arr)-1 {
			next := arr[i+1]
			nextTime = next.time
		}

		if nextTime == "" {
			nextTime = current.time
		}

		lrc = append(lrc, fmt.Sprintf("%s %s", nextTime, zh))
	}

	text := strings.Join(lrc, "\r")
	fs := filename + ".lrc"
	os.WriteFile(fs, []byte(text), os.ModePerm)
	fmt.Println(fs + " done!")
}

type LRC struct {
	time     string
	subtitle string
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

			obj.time = fmt.Sprintf("[%s:%s.%s]", minutes, seconds, detail[1])
		} else {
			obj.time = fmt.Sprintf("[%s]", timeline)
		}

		obj.subtitle = strings.TrimLeft(info[2], "\n\r")
		return obj
	}

	return obj
}
