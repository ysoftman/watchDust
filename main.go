package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/robfig/cron"
	"google.golang.org/appengine"
	appenginelog "google.golang.org/appengine/log"
)

const configFileName = "watchDustConfig.toml"

var fineDustMsg = ""

var conf serverConfig

func loadConfig() {
	// 파일로 부터 파싱해서 conf 로 저장하기
	_, err := toml.DecodeFile(configFileName, &conf)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("%v was loaded\n%v\n", configFileName, conf)
}

func main() {
	loadConfig()

	// 20190406 google compute engine 무료 기간 만료
	// App Engine에서 애플리케이션이 배포되어 있는 로컬 파일 시스템은 쓸 수 없습니다.
	// google.golang.org/appengine/log 으로 로깅 가능하다.
	// f, err := os.OpenFile("wd.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("can't open log file")
	// }
	// defer f.Close()
	// log.SetOutput(f)
	// log.Println("start Watch-Dust")

	serverType := flag.String("servertype", "gae", "test|noraml|gae(google app engin)")
	flag.Parse()
	log.Println("servertype :", *serverType)
	if *serverType == "test" {
		// 연결 확인만 하고 종료
		airReuslt := openapiAirKorea("")
		analDustInfo(airReuslt, "")
	} else if *serverType == "normal" {
		// 일반 서버 환경으로 운영시
		watchingDust()
	} else if *serverType == "gae" {
		// GAE(google app engine) 환경으로 운영시
		http.HandleFunc("/", handlerIndex)
		http.HandleFunc("/watchDust", handlerWatchingDust)
		appengine.Main()
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/ 요청 처리")
	out := `
디폴트 측정소 미세먼지 정보
https://watchdust.appspot.com/watchDust

dustinfo 슬랙 채널에 미세먼지 정보 발송
https://watchdust.appspot.com/watchDust?slack=dustinfo
https://watchdust.appspot.com/watchDust?slack=dustinfo
https://watchdust.appspot.com/watchDust?slack=dustinfo
https://watchdust.appspot.com/watchDust?slack=dustinfo

측정소(station) 정보 참고
https://www.airkorea.or.kr/web/stationInfo
https://www.airkorea.or.kr/web/realSearch

github
https://github.com/ysoftman/watchDust
`
	fmt.Fprintln(w, out)
}

func handlerWatchingDust(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	appenginelog.Infof(ctx, "/watchDust 요청 처리")
	query := r.URL.Query()

	log.Println("----------", query.Get("station"))
	airResult := openapiAirKoreaGAE(r, query.Get("station"))
	dustinfomsg := analDustInfo(airResult, query.Get("station"))
	out := dustinfomsg

	log.Println("----------", query.Get("slack"))
	if len(query.Get("slack")) > 0 {
		sendToSlackGAE(r, query.Get("slack"), dustinfomsg)
		out += "slack channel = " + query.Get("slack")
	}
	fmt.Fprintln(w, out)
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

	// analDustInfo(openapiAirKorea())

	c := cron.New()
	// c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	// every second
	// c.AddFunc("* * * * * *", func() { analDustInfo(openapiAirKorea()) })
	// every minute
	// c.AddFunc("0 */1 * * * *", func() { analDustInfo(openapiAirKorea()) })
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	// 9-21/3 : 9~21시 사이 3시간 간격으로 => 9 12 15 18 21시
	c.AddFunc("0 0 9-21/"+strconv.Itoa(conf.WatchIntervalHour)+" * * *", func() {
		airReuslt := openapiAirKorea("")
		dustinfomsg := analDustInfo(airReuslt, "")
		sendToSlack(conf.SlackAPI.Channel, dustinfomsg)
	})
	c.Start()
	for {
		select {
		case <-time.After(10 * time.Second):
		}
	}
}
