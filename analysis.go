package main

import (
	"log"
)

func analDustInfo(jsonDustInfo *dustinfoResp, stationName string) string {
	dustinfoMsg := ""
	markedCode := make(map[string]bool)
	for _, val := range jsonDustInfo.Response.Body.Items {
		// 이미지가 없는 예보 자료는 스킵
		if len(val.ImageURL1) == 0 {
			continue
		}
		// 앞쪽의 아이템이 최신정보, 이후 중복 정보는 제거
		if markedCode[val.InformCode] {
			continue
		}
		dustinfoMsg += val.DataTime + "\n" +
			val.InformCode + " " + val.InformOverall + "\n" +
			val.InformGrade + "\n" +
			"이미지:(" + val.ImageURL6 + ")" + "\n\n"
		markedCode[val.InformCode] = true
	}
	log.Println(dustinfoMsg)

	return dustinfoMsg
}
