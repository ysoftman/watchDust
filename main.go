package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v11"
	"github.com/robfig/cron"
	"google.golang.org/appengine/v2"

	appenginelog "google.golang.org/appengine/v2/log"
)

const configFileName = "watchDustConfig.toml"

// fineDustMsg = ""
var (
	conf  serverConfig
	isGAE bool
)

func loadConfig() {
	if _, err := toml.DecodeFile(configFileName, &conf); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := env.Parse(&conf); err != nil {
		log.Println(err)
		os.Exit(2)
	}
	log.Printf("%v was loaded\n%+v\n", configFileName, conf)
}

func main() {
	loadConfig()

	// 20190406 google compute engine 무료 기간 만료
	// App Engine에서 애플리케이션이 배포되어 있는 로컬 파일 시스템은 쓸 수 없습니다.
	// google.golang.org/appengine/v2/log 으로 로깅 가능하다.
	// f, err := os.OpenFile("wd.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("can't open log file")
	// }
	// defer f.Close()
	// log.SetOutput(f)
	// log.Println("start Watch-Dust")

	serverType := flag.String("servertype", "gae", "test|normal|gae(google app engin)")
	flag.Parse()
	isGAE = (*serverType == "gae")
	log.Println("servertype :", *serverType)
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/watchDust", handlerWatchingDust)
	switch *serverType {
	case "normal":
		// 일반 서버 환경으로 운영시
		go cronWatchingDust()
		log.Printf("Starting HTTP server on port %d\n", conf.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil); err != nil {
			log.Fatal(err)
		}
	case "gae":
		// GAE(google app engine) 환경으로 운영시
		appengine.Main()
	}
}

func SetCommonResponseHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "get")
}

func logInfo(r *http.Request, msg string) {
	if isGAE {
		ctx := appengine.NewContext(r)
		appenginelog.Infof(ctx, msg)
	} else {
		log.Println(msg)
	}
}

func airKorea(r *http.Request) *dustinfoResp {
	if isGAE {
		return openapiAirKoreaGAE(r)
	}
	return openapiAirKorea()
}

func slack(r *http.Request, channel, msg string) (error, string) {
	if isGAE {
		return sendToSlackGAE(r, channel, msg)
	}
	return sendToSlack(channel, msg)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	logInfo(r, "/ 요청 처리")
	out := `
디폴트 미세먼지 정보
https://watchdust.appspot.com/watchDust

dustinfo 슬랙 채널에 미세먼지 정보 발송
https://watchdust.appspot.com/watchDust?slack=dustinfo

github
https://github.com/ysoftman/watchDust
`
	SetCommonResponseHeader(w)
	if _, err := fmt.Fprintln(w, out); err != nil {
		log.Println(err)
	}
}

func handlerWatchingDust(w http.ResponseWriter, r *http.Request) {
	logInfo(r, "/watchDust 요청 처리")
	query := r.URL.Query()

	log.Println("---------- openapiAirKoreaGAE")
	airResult := airKorea(r)
	dustinfomsg := analDustInfo(airResult)
	out := dustinfomsg

	log.Println("---------- slack channel:", query.Get("slack"))
	if len(query.Get("slack")) > 0 {
		out += "slack channel = " + query.Get("slack")
		err, respMsg := slack(r, query.Get("slack"), dustinfomsg)
		if err != nil {
			log.Println(err)
			out += err.Error()
		} else {
			out += "\n" + respMsg
		}
	}
	SetCommonResponseHeader(w)
	if _, err := fmt.Fprintln(w, out); err != nil {
		log.Println(err)
	}
}

func cronWatchingDust() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for {
	// 		select {
	// 		case <-time.After(time.Duration(conf.WatchIntervalHour) * time.Hour):
	// 			// fineDustMsg = fineDustSearch())
	// 			analDustInfo(airKorea(nil))
	// 		}
	// 	}
	// }()
	// wg.Wait()

	// analDustInfo(airKorea(nil))

	c := cron.New()
	// c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	// every second
	// c.AddFunc("* * * * * *", func() { analDustInfo(airKorea(nil)) })
	// every minute
	// c.AddFunc("0 */1 * * * *", func() { analDustInfo(airKorea(nil)) })
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	// 9-21/3 : 9~21시 사이 3시간 간격으로 => 9 12 15 18 21시
	if err := c.AddFunc("0 0 9-21/"+strconv.Itoa(conf.WatchIntervalHour)+" * * *", func() {
		airResult := airKorea(nil)
		dustinfomsg := analDustInfo(airResult)
		if err, _ := slack(nil, conf.SlackAPI.Channel, dustinfomsg); err != nil {
			log.Println(err)
		}
	}); err != nil {
		log.Println(err)
	}
	c.Start()
	select {}
	// for {
	// 	select {
	// 	case <-time.After(10 * time.Second):
	// 	}
	// }
}
