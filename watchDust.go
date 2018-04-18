package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
)

var fineDustMsg = ""

var conf serverConfig

func loadConfig() {
	// 파일로 부터 파싱해서 conf 로 저장하기
	_, err := toml.DecodeFile("watchDustConfig.toml", &conf)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("%v\n", conf)
}

func main() {
	// fineDustMsg = fineDustSearch()
	// if len(os.Args) != 2 {
	// 	fmt.Printf("ex) %s [airkorea service key]\n", os.Args[0])
	// 	os.Exit(1)
	// }
	loadConfig()
	watchingDust()
}

func watchingDust() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for {
	// 		select {
	// 		case <-time.After(time.Duration(conf.WatchIntervalHour) * time.Hour):
	// 			// fineDustMsg = fineDustSearch())
	// 			analDustInfo(openapiAirKorea())
	// 		}
	// 	}
	// }()
	// wg.Wait()

	analDustInfo(openapiAirKorea())

	c := cron.New()
	// c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	// every second
	// c.AddFunc("* * * * * *", func() { analDustInfo(openapiAirKorea()) })
	// every minute
	// c.AddFunc("0 */1 * * * *", func() { analDustInfo(openapiAirKorea()) })
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("0 0 */"+strconv.Itoa(conf.WatchIntervalHour)+" * * *", func() { analDustInfo(openapiAirKorea()) })
	c.Start()
	for {
	}
}

// 다음 검색으로 미세먼지 파악하기
// 현재 지역의 미세먼지 정보는 javascript 로 파싱할 수 없어, 경기 지역만 본다.
func fineDustSearch() string {
	// Request the HTML page.
	res, err := http.Get("https://search.daum.net/search?w=tot&ie=UTF-8&q=%EB%AF%B8%EC%84%B8%EB%A8%BC%EC%A7%80")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		errmsg := "status code error: (" + strconv.Itoa(res.StatusCode) + ")" + res.Status
		log.Println(errmsg)
		return errmsg
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}

	// #airPollutionNColl > div.coll_cont > div > div.wrap_whole > div.cont_map.bg_map > div.map_region > ul > li.city_03 > a > span
	titregion := ""
	screenout := ""
	txtstate := ""
	selector := "#airPollutionNColl div.coll_cont div div.wrap_whole div.cont_map.bg_map div.map_region ul  li.city_03 a"
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		titregion = s.Find("em.tit_region").Text()
		screenout = s.Find("span.screen_out").Text()
		txtstate = s.Find("span.txt_state").Text()
		// fmt.Printf("%d: %s %s %s\n", i, titregion, screenout, txtstate)
	})
	msg := titregion + "지역 미세먼지는 " + screenout + "(" + txtstate + ")" + " 입니다.\n"
	fmt.Print(msg)
	return msg
}

// 공공데이터 airkorea 로 부터 미세먼지 정보 파악
func openapiAirKorea() *dustinfoResp {
	url := "http://openapi.airkorea.or.kr/openapi/services/rest/ArpltnInforInqireSvc/getMsrstnAcctoRltmMesureDnsty?numOfRows=" + strconv.Itoa(conf.OpenapiAirkorea.NumOfRows) +
		"&pageNo=" + strconv.Itoa(conf.OpenapiAirkorea.PageNo) +
		"&startPage=1" +
		"&stationName=" + GetEncURL(conf.OpenapiAirkorea.StationName) +
		"&dataTerm=" + conf.OpenapiAirkorea.DataTerm +
		"&ver=" + conf.OpenapiAirkorea.Ver +
		"&_returnType=json" +
		"&serviceKey=" + conf.OpenapiAirkorea.Servicekey

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	// 응답결과 출력
	// ioutil.ReadAll 로 resp.Body 읽고 나면 resp.Body 내용은 사라진다.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can't read resp.Body")
	}

	// bodystring := string(body)
	// fmt.Println(bodystring)

	jsonDustInfo := &dustinfoResp{}
	json.Unmarshal([]byte(body), &jsonDustInfo)

	return jsonDustInfo
}

func analDustInfo(jsonDustInfo *dustinfoResp) {

	// dataTime	측정일
	// mangName	측정망 정보
	// so2Value	아황산가스 농도
	// coValue		일산화탄소 농도
	// o3Value		오존 농도
	// no2Value	이산화질소 농도
	// pm10Value	미세먼지(PM10) 농도
	// pm10Value24	미세먼지(PM10) 24시간예측이동농도
	// pm25Value	미세먼지(PM2.5) 농도
	// pm25Value24	미세먼지(PM2.5) 24시간예측이동농도
	// khaiValue	통합대기환경수치
	// khaiGrade	통합대기환경지수
	// so2Grade	아황산가스 지수
	// coGrade		일산화탄소 지수
	// o3Grade		오존 지수
	// no2Grade	이산화질소 지수
	// pm10Grade	미세먼지(PM10) 24시간 등급
	// pm25Grade	미세먼지(PM2.5) 24시간 등급
	// pm10Grade1h	미세먼지(PM10) 1시간 등급
	// pm25Grade1h	미세먼지(PM2.5) 1시간 등급
	if len(jsonDustInfo.List) < 1 {
		log.Printf("%s 측정소 정보가 없습니다.", conf.OpenapiAirkorea.StationName)
		return
	}
	i := 0
	dustinfoMsg := "측정시간: " + jsonDustInfo.List[i].DataTime + "\n" +
		"측정소: " + conf.OpenapiAirkorea.StationName + "\n" +
		"아황산가스: " + jsonDustInfo.List[i].So2Value + "(" + toGradeStr(jsonDustInfo.List[i].So2Grade) + ")" + "\n" +
		"일산화탄소: " + jsonDustInfo.List[i].CoValue + "(" + toGradeStr(jsonDustInfo.List[i].CoGrade) + ")" + "\n" +
		"오존: " + jsonDustInfo.List[i].O3Value + "(" + toGradeStr(jsonDustInfo.List[i].O3Grade) + ")" + "\n" +
		"이산화질소: " + jsonDustInfo.List[i].No2Value + "(" + toGradeStr(jsonDustInfo.List[i].No2Grade) + ")" + "\n" +
		"미세먼지(pm10): " + jsonDustInfo.List[i].Pm10Value + "(" + toGradeStr(jsonDustInfo.List[i].Pm10Grade) + ")" + "\n" +
		"초미세먼지(pm25): " + jsonDustInfo.List[i].Pm25Value + "(" + toGradeStr(jsonDustInfo.List[i].Pm25Grade) + ")" + "\n"

	log.Println(dustinfoMsg)

	sendToSlack(dustinfoMsg)
}

func toGradeStr(grade string) string {
	if grade == "1" {
		return "좋음:party_parrot:"
	} else if grade == "2" {
		return "보통:smile:"
	} else if grade == "3" {
		return "나쁨:angry:"
	} else if grade == "4" {
		return "매우나쁨:angryyy:"
	}
	return "_"
}

// GetEncURL : URL 인코딩
func GetEncURL(str string) string {
	t := &url.URL{Path: str}
	encurl := t.String()
	fmt.Printf("encode url(%s) = %s\n", str, encurl)
	return encurl
}

func sendToSlack(msg string) {
	content := "token=" + conf.SlackAPI.Token +
		"&channel=" + conf.SlackAPI.Channel +
		"&username=" + conf.SlackAPI.Username +
		"&text=" + msg
	// log.Println(content)
	reqBody := bytes.NewBufferString(content)
	resp, err := http.Post("https://slack.com/api/chat.postMessage", "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("%s ... send to slack is success\n", string(respBody))
	}
}
