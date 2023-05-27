package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
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

	srt := strings.Split(string(okbyte), "\n\r")

	lrc := []string{}

	lrc = append(lrc, "[00:00.000] 作词 : laof")
	lrc = append(lrc, "[00:00.000] 作曲 : laof")

	for _, part := range srt {

		info := strings.Split(part, "\r")

		if len(info) >= 3 {

			/**
			1
			00:00:16,766 --> 00:00:18,066
			reporting work
			*/

			timeArr := strings.Split(info[1], "-->")

			// 00:00:16,766 --> 00:00:18,066

			if len(timeArr) == 2 {
				timeline := strings.TrimSpace(timeArr[0])
				// 00:00:16,766
				detail := strings.Split(timeline, ",")
				// 00:00:16
				time := strings.Split(detail[0], ":")

				if len(time) == 3 {

					h, _ := strconv.Atoi(time[0])
					m, _ := strconv.Atoi(time[1])
					seconds := time[2]

					m = h*60 + m

					minutes := strconv.Itoa(m)

					if m < 10 {
						minutes = "0" + strconv.Itoa(m)
					}

					timeline = fmt.Sprintf("[%s:%s.%s]", minutes, seconds, detail[1])
				} else {
					timeline = fmt.Sprintf("[%s]", timeline)
				}

				subtitle := info[2]
				subtitle = strings.TrimLeft(subtitle, "\n\r")
				lrc = append(lrc, fmt.Sprintf("%s %s", timeline, subtitle))
			}

		}

	}

	text := strings.Join(lrc, "\r")
	fs := filename + ".lrc"
	os.WriteFile(fs, []byte(text), os.ModePerm)
	fmt.Println(fs + " ok")
}
