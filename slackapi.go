package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func getSlackURL(slackchannel, msg string) string {
	sc := conf.SlackAPI.Channel
	if len(slackchannel) > 0 {
		sc = slackchannel
	}
	return "token=" + conf.SlackAPI.Token +
		"&channel=" + sc +
		"&username=" + conf.SlackAPI.Username +
		"&text=" + msg
}

func sendToSlackGAE(r *http.Request, slackchannel, msg string) {
	content := getSlackURL(slackchannel, msg)
	reqBody := bytes.NewBufferString(content)

	// appengine 에서는 기본 http client 를 할 수 없다.
	// google.golang.org/appengine/urlfetch 를 사용해야 하나.
	// http.DefaultTransport and http.DefaultClient are not available in App Engine. See https://cloud.google.com/appengine/docs/go/urlfetch/
	// resp, err := http.Post("https://slack.com/api/chat.postMessage", "application/x-www-form-urlencoded", reqBody)
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Post("https://slack.com/api/chat.postMessage", "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

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

func sendToSlack(slackchannel, msg string) {
	content := getSlackURL(slackchannel, msg)
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
