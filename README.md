# 미세 먼지 알림 :smile:

## 개발 환경

- 미세먼지 정보 출처 : data 공공포털 openapi 한국환경공단_대기오염정보 조회 서비스(airkorea)
  https://www.data.go.kr/subMain.jsp#/L3B1YnIvcG90L215cC9Jcm9zTXlQYWdlL29wZW5EZXZEZXRhaWxQYWdlJEBeMDgyTTAwMDAxMzBeTTAwMDAxMzUkQF5wdWJsaWNEYXRhRGV0YWlsUGs9dWRkaTo3MDkxMTBlNy1kN2IxLTQ0MjEtOTBiYS04NGE2OWY5ODBjYWJfMjAxNjA4MDgxMTE0JEBecHJjdXNlUmVxc3RTZXFObz0zODMzNDExJEBecmVxc3RTdGVwQ29kZT1TVENEMDE=


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
