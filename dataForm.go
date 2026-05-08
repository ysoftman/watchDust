package main

type serverConfig struct {
	Title             string `toml:"Title"`
	WatchIntervalHour int    `toml:"WatchIntervalHour"`
	Port              int    `toml:"Port" env:"WATCHDUST_PORT" envDefault:"8080"`
	OpenapiAirkorea   struct {
		URL        string `toml:"url"`
		Servicekey string `toml:"servicekey" env:"WATCHDUST_OPENAPIAIRKOREA_SERVICE_KEY" envDefault:""`
		NumOfRows  int    `toml:"numOfRows"`
		PageNo     int    `toml:"pageNo"`
		DataTerm   string `toml:"dataTerm"`
	} `toml:"openapi_airkorea"`
	SlackAPI struct {
		Token    string `toml:"token" env:"WATCHDUST_SLACKAPI_TOKEN" envDefault:""`
		Channel  string `toml:"channel"`
		Username string `toml:"username"`
	} `toml:"slack_api"`
}

// dustinfoItem : 대기오염예보통보 조회 응답 아이템
type dustinfoItem struct {
	DataTime      string `json:"dataTime"`      // 통보시간
	InformCode    string `json:"informCode"`    // 통보코드 (PM10, PM25, O3)
	InformOverall string `json:"informOverall"` // 예보개황
	InformCause   string `json:"informCause"`   // 발생원인
	InformGrade   string `json:"informGrade"`   // 예보등급
	ActionKnack   string `json:"actionKnack"`   // 행동요령
	ImageURL1     string `json:"imageUrl1"`     // 시간대별 예측모델 결과사진 (6:00, 12:00, 18:00, 24:00 KST)
	ImageURL2     string `json:"imageUrl2"`     // 시간대별 예측모델 결과사진 (6:00, 12:00, 18:00, 24:00 KST)
	ImageURL3     string `json:"imageUrl3"`     // 시간대별 예측모델 결과사진 (6:00, 12:00, 18:00, 24:00 KST)
	ImageURL4     string `json:"imageUrl4"`     // 시간대별 예측모델 결과사진 (6:00, 12:00, 18:00, 24:00 KST)
	ImageURL5     string `json:"imageUrl5"`     // 시간대별 예측모델 결과사진 (6:00, 12:00, 18:00, 24:00 KST)
	ImageURL6     string `json:"imageUrl6"`     // 시간대별 예측모델 결과사진 (6:00, 12:00, 18:00, 24:00 KST)
	ImageURL7     string `json:"imageUrl7"`     // PM10 한반도 대기질 예측모델결과 애니메이션
	ImageURL8     string `json:"imageUrl8"`     // PM2.5 한반도 대기질 예측모델결과 애니메이션
	ImageURL9     string `json:"imageUrl9"`     // O3 한반도 대기질 예측모델결과 애니메이션
	InformData    string `json:"informData"`    // 예측통보시간
}

// ImageURLs : imageUrl1~9 를 슬라이스로 반환 (인덱스 0이 imageUrl1)
func (it dustinfoItem) ImageURLs() []string {
	return []string{
		it.ImageURL1, it.ImageURL2, it.ImageURL3,
		it.ImageURL4, it.ImageURL5, it.ImageURL6,
		it.ImageURL7, it.ImageURL8, it.ImageURL9,
	}
}

// dustinfoResp : data.go.kr 한국환경공단 대기오염예보통보 조회(getMinuDustFrcstDspth) 응답
// 참고: https://www.data.go.kr/data/15073861/openapi.do
type dustinfoResp struct {
	Response struct {
		Header struct {
			ResultCode string `json:"resultCode"` // 결과코드
			ResultMsg  string `json:"resultMsg"`  // 결과메시지
		} `json:"header"`
		Body struct {
			TotalCount int            `json:"totalCount"` // 전체 결과 수
			PageNo     int            `json:"pageNo"`     // 페이지 번호
			NumOfRows  int            `json:"numOfRows"`  // 한 페이지 결과 수
			Items      []dustinfoItem `json:"items"`
		} `json:"body"`
	} `json:"response"`
}
