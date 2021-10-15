package main

type serverConfig struct {
	Title             string `toml:"Title"`
	WatchIntervalHour int    `toml:"WatchIntervalHour"`
	OpenapiAirkorea   struct {
		URL        string `toml:"url"`
		Servicekey string `toml:"servicekey"`
		NumOfRows  int    `toml:"numOfRows"`
		PageNo     int    `toml:"pageNo"`
		DataTerm   string `toml:"dataTerm"`
	} `toml:"openapi_airkorea"`
	SlackAPI struct {
		Token    string `toml:"token"`
		Channel  string `toml:"channel"`
		Username string `toml:"username"`
	} `toml:"slack_api"`
}

type dustinfoResp struct {
	Response struct {
		Body struct {
			TotalCount int `json:"totalCount"`
			Items      []struct {
				ImageURL4     string      `json:"imageUrl4"`
				InformCode    string      `json:"informCode"`
				ImageURL5     string      `json:"imageUrl5"`
				ImageURL6     string      `json:"imageUrl6"`
				ActionKnack   interface{} `json:"actionKnack"`
				InformCause   string      `json:"informCause"`
				InformOverall string      `json:"informOverall"`
				InformData    string      `json:"informData"`
				InformGrade   string      `json:"informGrade"`
				DataTime      string      `json:"dataTime"`
				ImageURL3     string      `json:"imageUrl3"`
				ImageURL2     string      `json:"imageUrl2"`
				ImageURL1     string      `json:"imageUrl1"`
			} `json:"items"`
			PageNo    int `json:"pageNo"`
			NumOfRows int `json:"numOfRows"`
		} `json:"body"`
		Header struct {
			ResultMsg  string `json:"resultMsg"`
			ResultCode string `json:"resultCode"`
		} `json:"header"`
	} `json:"response"`
}
