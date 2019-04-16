package main

import (
	"log"
	"strconv"
)

func analDustInfo(jsonDustInfo *dustinfoResp, stationName string) string {

	// dataTime		측정일
	// mangName		측정망 정보
	// so2Value		아황산가스 농도(ppm)
	// coValue		일산화탄소 농도(ppm)
	// o3Value		오존 농도(ppm)
	// no2Value		이산화질소 농도(ppm)
	// pm10Value	미세먼지(PM10) 농도(㎍/㎥)
	// pm10Value24	미세먼지(PM10) 24시간예측이동농도(㎍/㎥)
	// pm25Value	미세먼지(PM2.5) 농도(㎍/㎥)
	// pm25Value24	미세먼지(PM2.5) 24시간예측이동농도(㎍/㎥)
	// khaiValue	통합대기환경수치
	// khaiGrade	통합대기환경지수
	// so2Grade		아황산가스 지수
	// coGrade		일산화탄소 지수
	// o3Grade		오존 지수
	// no2Grade		이산화질소 지수
	// pm10Grade	미세먼지(PM10) 24시간 등급
	// pm25Grade	미세먼지(PM2.5) 24시간 등급
	// pm10Grade1h	미세먼지(PM10) 1시간 등급
	// pm25Grade1h	미세먼지(PM2.5) 1시간 등급

	sn := conf.OpenapiAirkorea.StationName
	if len(stationName) > 0 {
		sn = stationName
	}
	if len(jsonDustInfo.List) < 1 {
		log.Printf("%s 측정소 정보가 없습니다.", sn)
		return ""
	}
	i := 0
	dustinfoMsg := "측정시간: " + jsonDustInfo.List[i].DataTime + "\n" +
		"측정소: " + sn + "\n" +
		"아황산가스: " + jsonDustInfo.List[i].So2Value + "ppm(" + toGradeStr(jsonDustInfo.List[i].So2Grade) + ")" + "\n" +
		"일산화탄소: " + jsonDustInfo.List[i].CoValue + "ppm(" + toGradeStr(jsonDustInfo.List[i].CoGrade) + ")" + "\n" +
		"오존: " + jsonDustInfo.List[i].O3Value + "ppm(" + toGradeStr(jsonDustInfo.List[i].O3Grade) + ")" + "\n" +
		"이산화질소: " + jsonDustInfo.List[i].No2Value + "ppm(" + toGradeStr(jsonDustInfo.List[i].No2Grade) + ")" + "\n" +
		// "미세먼지(pm10): " + jsonDustInfo.List[i].Pm10Value + "(" + toGradeStr(jsonDustInfo.List[i].Pm10Grade) + ")" + "\n" +
		// "초미세먼지(pm25): " + jsonDustInfo.List[i].Pm25Value + "(" + toGradeStr(jsonDustInfo.List[i].Pm25Grade) + ")" + "\n"
		"미세먼지(pm10): " + jsonDustInfo.List[i].Pm10Value + "㎍/㎥(" + toWHOPM10GradeStr(jsonDustInfo.List[i].Pm10Value) + ")" + "\n" +
		"초미세먼지(pm25): " + jsonDustInfo.List[i].Pm25Value + "㎍/㎥(" + toWHOPM25GradeStr(jsonDustInfo.List[i].Pm25Value) + ")" + "\n"
	log.Println(dustinfoMsg)

	return dustinfoMsg
}

func toWHOPM10GradeStr(value string) string {
	nValue, err := strconv.Atoi(value)
	if err != nil {
		log.Println(err)
		return "_"
	}
	if nValue <= 15 {
		return "최고:shuffle_parrot:"
	} else if nValue <= 30 {
		return "좋음:party_parrot:"
	} else if nValue <= 40 {
		return "양호:pikachu:"
	} else if nValue <= 50 {
		return "보통:smile:"
	} else if nValue <= 75 {
		return "나쁨:scream:"
	} else if nValue <= 100 {
		return "상당히나쁨:angry:"
	} else if nValue <= 150 {
		return "매우나쁨:angryy:"
	}
	return "최악:angryyy:"
}

func toWHOPM25GradeStr(value string) string {
	nValue, err := strconv.Atoi(value)
	if err != nil {
		log.Println(err)
		return "_"
	}
	if nValue <= 8 {
		return "최고:shuffle_parrot:"
	} else if nValue <= 15 {
		return "좋음:party_parrot:"
	} else if nValue <= 20 {
		return "양호:pikachu:"
	} else if nValue <= 25 {
		return "보통:smile:"
	} else if nValue <= 37 {
		return "나쁨:scream:"
	} else if nValue <= 50 {
		return "상당히나쁨:angry:"
	} else if nValue <= 75 {
		return "매우나쁨:angryy:"
	}
	return "최악:angryyy:"
}

func toGradeStr(grade string) string {
	if grade == "1" {
		return "좋음:party_parrot:"
	} else if grade == "2" {
		return "보통:smile:"
	} else if grade == "3" {
		return "나쁨:scream:"
	} else if grade == "4" {
		return "매우나쁨:angryy:"
	}
	return "_"
}
