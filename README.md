# Watch-Dust

## References

- data 공공포털 openapi 한국환경공단 대기오염정보 조회 서비스(airkorea)

  - airkorea 지역명 검색 : <https://www.airkorea.or.kr/web/realSearch>
  - data.go.kr 로그인해서 활용기간 2년씩 연장 필요 <https://www.data.go.kr/iim/api/selectAPIAcountView.do>
  - 사용예

  ```bash
  http://apis.data.go.kr/B552584/ArpltnInforInqireSvc/getMinuDustFrcstDspth?numOfRows=4&pageNo=1&searchDate=2021-10-15&returnType=json&serviceKey=aaa
  ```

- slack 알림

  - 채널 메시지 보내기 api : <https://api.slack.com/methods/chat.postMessage>
  - 토근 발급 : <https://api.slack.com/custom-integrations/legacy-tokens>
  - 사용예

  ```bash
  curl -X POST https://slack.com/api/chat.postMessage -d "token=aaaaa&channel=dustinfo&username=watchDust bot&text=미세먼지 정보입니다."
  ```

## 빌드 및 실행

```bash
# get packages
go get github.com/BurntSushi/toml
go get github.com/robfig/cron
go get github.com/PuerkitoBio/goquery
go get google.golang.org/appengine/v2

# build
go build

# execute
# default:gae(google app engin) 환경
./watchdust

# 일반 서버 환경
./watchdust -servertype normal

# 테스트로 한번 실행하고 종료
./watchdust -servertype test
```

## google app engine 사용

```bash
# gcloud 설치 - mac
brew install google-cloud-sdk
# 또는 수동 설치
cd ~/workspace
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-365.0.0-darwin-x86_64.tar.gz
rm -rf google-cloud-sdk
tar zxvf google-cloud-sdk-365.0.0-darwin-x86_64.tar.gz
# 컴포넌트 설치, golang 패키지 설치
gcloud components install app-engine-go

# gcloud 설치 - ubuntu
export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)"
echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
sudo apt-get update && sudo apt-get install google-cloud-sdk
sudo apt-get install google-cloud-sdk-app-engine-python google-cloud-sdk-app-engine-go google-cloud-sdk-datastore-emulator


# log
https://cloud.google.com/appengine/docs/standard/go/logs/
# cron.yaml
https://cloud.google.com/appengine/docs/standard/go/config/cron
# app.yaml
https://cloud.google.com/appengine/docs/standard/go/config/appref


# google cloud 올리기전에 로컬에서 테스트 해볼 수 있다.
# 아래 명령을 실행해두면 .go 소스 수정때마다 자동 빌드 된다.
# google-cloud-sdk.tar.gz 로 설치했을 경우
~/workspace/google-cloud-sdk/bin/dev_appserver.py app.yaml --port 9999

# gcloud 인증(브라우저 열리고 로그인)
# gcloud auth login
# appengine/v2 (runtime: go116) 을 배포하기 위해선 현재 beta 를 사용해야 한다.
gcloud beta auth login

# google cloud 초기화
# url 링크 후 verification code 확인하여 입력
# 기존 프로젝트 또는 신규 프로젝트 생성 선택
# Compute Region and Zone 선택
gcloud init

# glcoud 구글 app engine 에 배포하기
# 배포 종료시 접속 가능한 url 이 표시된다.
# 배포전 아래 내용이 출력된다. 이상이 있다면 gcloud init 로 다시 설정하자.
# descriptor:      [/Users/ysoftman/workspace/watchDust/app.yaml]
# source:          [/Users/ysoftman/workspace/watchDust]
# target project:  [watchdust]
# target service:  [default]
# target version:  [20190416t141405]
# target url:      [https://watchdust.appspot.com]
# --verion 버전 명시
# --promote 현재 배포한 버전이 모든 트랙픽(100%)을 받도록 한다. 기존 버전의 인스턴스는 트랙픽 0% 이 된다.
# --stop-previous-version 새버전이 올라가면 기존 버전은 stop 하도록 한다.
# gcloud app deploy ./app.yaml --version 20211119-1 --promote --stop-previous-version
# appengine/v2 (runtime: go116) 을 배포하기 위해선 현재 beta 를 사용해야 한다.
gcloud beta app deploy ./app.yaml --version 20211122-1 --promote --stop-previous-version

# 크론 작업 cron.yaml
gcloud app deploy cron.yaml

# 배포가 완료되면 확인
https://watchdust.appspot.com

# 배포 후 접속 URL 확인 하기
gcloud app browse

# 앱 로그 확인
https://console.cloud.google.com/logs/viewer?project=watchdust
gcloud app logs tail -s default
```
