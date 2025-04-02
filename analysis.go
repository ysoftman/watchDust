package main

import (
	"log"
	"strings"
)

func analDustInfo(jsonDustInfo *dustinfoResp) string {
	// 미세먼지 정보 전체 데이터 확인용
	// if jsonStr, err := json.Marshal(&jsonDustInfo.Response); err == nil {
	// 	log.Println(string(jsonStr))
	// }
	dustinfoMsg := ""
	imageURL := ""
	markedCode := make(map[string]bool)
	for _, val := range jsonDustInfo.Response.Body.Items {
		// 이미지가 없는 예보 자료는 스킵
		if len(val.ImageURL1) == 0 || len(val.ImageURL4) == 0 {
			continue
		}
		// 앞쪽의 아이템이 최신정보, 이후 중복 정보는 스킵
		if markedCode[val.InformCode] {
			continue
		}
		switch val.InformCode {
		case "PM10":
			imageURL = val.ImageURL1
		case "PM25":
			imageURL = val.ImageURL4
		}

		infoGrade := ""
		for i, v := range strings.Split(strings.ReplaceAll(val.InformGrade, " : ", ":"), ",") {
			infoGrade += v
			if (i+1)%5 == 0 {
				infoGrade += "\n"
			} else {
				infoGrade += "  "
			}
		}
		dustinfoMsg += val.DataTime + "\n" +
			val.InformCode + "\n" +
			strings.TrimLeft(val.InformOverall, " ○") + "\n" +
			infoGrade + "\n" +
			imageURL + "\n\n"
		markedCode[val.InformCode] = true
	}
	log.Println(dustinfoMsg)

	return dustinfoMsg
}
