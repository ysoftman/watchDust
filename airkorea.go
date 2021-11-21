package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/urlfetch"
)

// GetEncURL : URL 인코딩
func GetEncURL(str string) string {
	t := &url.URL{Path: str}
	encurl := t.String()
	fmt.Printf("encode url %s => %s\n", str, encurl)
	return encurl
}

func getAirKoreaURL() string {
	airKoreaURL := conf.OpenapiAirkorea.URL + "?numOfRows=" + strconv.Itoa(conf.OpenapiAirkorea.NumOfRows) +
		"&pageNo=" + strconv.Itoa(conf.OpenapiAirkorea.PageNo) +
		"&searchDate=" + time.Now().UTC().Format("2006-01-02") +
		"&returnType=json" +
		"&serviceKey=" + conf.OpenapiAirkorea.Servicekey
	fmt.Println("airKoreaURL:", airKoreaURL)
	return airKoreaURL
}

func openapiAirKoreaGAE(r *http.Request) *dustinfoResp {
	url := getAirKoreaURL()
	// appengine 에서는 기본 http client 를 할 수 없다.
	// google.golang.org/appengine/v2/urlfetch 를 사용해야 한다.
	// http.DefaultTransport and http.DefaultClient are not available in App Engine. See https://cloud.google.com/appengine/docs/go/urlfetch/
	// resp, err := http.Get(url)
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	client.Timeout = time.Second * 3
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err.Error())
		return &dustinfoResp{}
	}
	defer resp.Body.Close()

	// 응답결과 출력
	// ioutil.ReadAll 로 resp.Body 읽고 나면 resp.Body 내용은 사라진다.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can't read resp.Body")
		return &dustinfoResp{}
	}
	// bodystring := string(body)
	// log.Println(bodystring)
	jsonDustInfo := &dustinfoResp{}
	json.Unmarshal([]byte(body), &jsonDustInfo)

	return jsonDustInfo
}

// 공공데이터 airkorea 로 부터 미세먼지 정보 파악
func openapiAirKorea() *dustinfoResp {
	client := http.Client{
		Timeout: time.Second * 3,
	}
	url := getAirKoreaURL()
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err.Error())
		return &dustinfoResp{}
	}
	defer resp.Body.Close()

	// 응답결과 출력
	// ioutil.ReadAll 로 resp.Body 읽고 나면 resp.Body 내용은 사라진다.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("can't read resp.Body")
		return &dustinfoResp{}
	}
	// bodystring := string(body)
	// log.Println(jsonDustInfo)
	jsonDustInfo := &dustinfoResp{}
	json.Unmarshal([]byte(body), &jsonDustInfo)

	return jsonDustInfo
}
