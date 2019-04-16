package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

// GetEncURL : URL 인코딩
func GetEncURL(str string) string {
	t := &url.URL{Path: str}
	encurl := t.String()
	fmt.Printf("encode url(%s) = %s\n", str, encurl)
	return encurl
}

func getAirKoreaURL(stationName string) string {
	sn := conf.OpenapiAirkorea.StationName
	if len(stationName) > 0 {
		// sn = stationName
	}
	return "http://openapi.airkorea.or.kr/openapi/services/rest/ArpltnInforInqireSvc/getMsrstnAcctoRltmMesureDnsty?numOfRows=" + strconv.Itoa(conf.OpenapiAirkorea.NumOfRows) +
		"&pageNo=" + strconv.Itoa(conf.OpenapiAirkorea.PageNo) +
		"&stationName=" + GetEncURL(sn) +
		"&dataTerm=" + conf.OpenapiAirkorea.DataTerm +
		"&ver=" + conf.OpenapiAirkorea.Ver +
		"&_returnType=json" +
		"&serviceKey=" + conf.OpenapiAirkorea.Servicekey
}

func openapiAirKoreaGAE(r *http.Request, stationName string) *dustinfoResp {
	url := getAirKoreaURL(stationName)
	// appengine 에서는 기본 http client 를 할 수 없다.
	// google.golang.org/appengine/urlfetch 를 사용해야 하나.
	// http.DefaultTransport and http.DefaultClient are not available in App Engine. See https://cloud.google.com/appengine/docs/go/urlfetch/
	// resp, err := http.Get(url)
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get(url)
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
	bodystring := string(body)
	log.Println(bodystring)
	jsonDustInfo := &dustinfoResp{}
	json.Unmarshal([]byte(body), &jsonDustInfo)

	return jsonDustInfo
}

// 공공데이터 airkorea 로 부터 미세먼지 정보 파악
func openapiAirKorea(stationName string) *dustinfoResp {
	url := getAirKoreaURL(stationName)
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
	bodystring := string(body)
	log.Println(bodystring)
	jsonDustInfo := &dustinfoResp{}
	json.Unmarshal([]byte(body), &jsonDustInfo)

	return jsonDustInfo
}
