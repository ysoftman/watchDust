package main

import (
	"log"
	"slices"
	"strings"
)

// 대기오염 예보 통보코드
const (
	codePM10 = "PM10"
	codePM25 = "PM25"
	codeO3   = "O3"
)

// codeAliases : URL 매칭용 통보코드 표기 후보 (PM25는 PM2P5 표기로도 등장)
var codeAliases = map[string][]string{
	codePM10: {codePM10},
	codePM25: {codePM25, "PM2P5"},
	codeO3:   {codeO3},
}

func analDustInfo(jsonDustInfo *dustinfoResp) string {
	var dustinfoMsg strings.Builder
	markedCode := make(map[string]bool)
	for _, val := range jsonDustInfo.Response.Body.Items {
		// 앞쪽의 아이템이 최신정보, 이후 중복 정보는 스킵
		if markedCode[val.InformCode] {
			continue
		}

		matchedURLs := matchedImageURLs(val)
		// 통보코드와 매칭되는 이미지 URL이 하나도 없는 아이템은 스킵
		if len(matchedURLs) == 0 {
			continue
		}

		dustinfoMsg.WriteString(val.DataTime + "\n")
		dustinfoMsg.WriteString(val.InformCode + "\n")
		dustinfoMsg.WriteString(strings.TrimLeft(val.InformOverall, " ○") + "\n")
		dustinfoMsg.WriteString(formatInformGrade(val.InformGrade) + "\n")
		for _, line := range matchedURLs {
			dustinfoMsg.WriteString(line + "\n")
		}
		dustinfoMsg.WriteString("\n")
		markedCode[val.InformCode] = true
	}

	if dustinfoMsg.Len() == 0 {
		return "미세먼지 예보 데이터가 없습니다.\n"
	}
	log.Println(dustinfoMsg.String())
	return dustinfoMsg.String()
}

// formatInformGrade : "서울 : 보통, 부산 : 보통, ..." 형태 문자열을 5개씩 줄바꿈하여 정리
func formatInformGrade(s string) string {
	var b strings.Builder
	for i, v := range strings.Split(strings.ReplaceAll(s, " : ", ":"), ",") {
		b.WriteString(v)
		if (i+1)%5 == 0 {
			b.WriteString("\n")
		} else {
			b.WriteString("  ")
		}
	}
	return b.String()
}

// matchedImageURLs : informCode에 매칭되는 imageUrl 라인 목록을 반환
func matchedImageURLs(it dustinfoItem) []string {
	aliases, ok := codeAliases[it.InformCode]
	if !ok {
		aliases = []string{it.InformCode}
	}
	var lines []string
	for _, url := range it.ImageURLs() {
		if len(url) == 0 {
			continue
		}
		if !slices.ContainsFunc(aliases, func(alias string) bool {
			return strings.Contains(url, alias)
		}) {
			continue
		}
		lines = append(lines, url)
	}
	return lines
}
