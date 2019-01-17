# Watch-Dust :smile:

## 개발 환경

- data 공공포털 openapi 한국환경공단_대기오염정보 조회 서비스(airkorea)
  - 데이터 갱신 주기
    - 실시간 정보 : 10분(매 시간 시간자료 갱신은 20분 전후로 반영됨)
    - 대기질 예보 정보 : 매 시간 22분, 57분
  - 사용예 : http://openapi.airkorea.or.kr/openapi/services/rest/ArpltnInforInqireSvc/getMsrstnAcctoRltmMesureDnsty?numOfRows=10&pageNo=1&stationName=수내동&dataTerm=DAILY&ver=1.3&_returnType=json&serviceKey=aaaaa

- slack 알림
  - 채널 메시지 보내기 api : https://api.slack.com/methods/chat.postMessage
  - 토근 발급 : https://api.slack.com/custom-integrations/legacy-tokens
  - 사용예 : curl -X POST https://slack.com/api/chat.postMessage -d "token=aaaaa&channel=dustinfo&username=watchDust bot&text=미세먼지 정보입니다."

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
./watchdust
```
