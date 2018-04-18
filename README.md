# Watch-Dust :smile: ~

## 개발 환경

- 미세먼지 정보 출처 : data 공공포털 openapi 한국환경공단_대기오염정보 조회 서비스(airkorea)

- slack 알림
  - 채널 메시지 보내기 api : https://api.slack.com/methods/chat.postMessage
  - 토근 발급 : https://api.slack.com/custom-integrations/legacy-tokens
  - 사용예 : curl -X POST https://slack.com/api/chat.postMessage -d "token=aaaaa&channel=dustinfo&username=watchDust bott&text=미세먼지 정보입니다."

## Perequisite (watchDustConfig.toml)

- openapi_airkorea.serverkey = https://www.data.go.kr/subMain.jsp#/L3B1YnIvcG90L215cC9Jcm9zTXlQYWdlL29wZW5EZXZEZXRhaWxQYWdlJEBeMDgyTTAwMDAxMzBeTTAwMDAxMzUkQF5wdWJsaWNEYXRhRGV0YWlsUGs9dWRkaTo3MDkxMTBlNy1kN2IxLTQ0MjEtOTBiYS04NGE2OWY5ODBjYWJfMjAxNjA4MDgxMTE0JEBecHJjdXNlUmVxc3RTZXFObz0zODMzNDExJEBecmVxc3RTdGVwQ29kZT1TVENEMDE=

- slack_api.token = https://api.slack.com/custom-integrations/legacy-tokens

## 빌드 및 실행

```bash
# get packages
go get github.com/BurntSushi/toml
go get github.com/robfig/cron
go get github.com/PuerkitoBio/goquery

# build
go build

# execute
./watchdust openapi_airkorea_servicekey
```
