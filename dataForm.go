package main

type serverConfig struct {
	Title             string `toml:"Title"`
	WatchIntervalHour int    `toml:"WatchIntervalHour"`
	OpenapiAirkorea   struct {
		Servicekey  string `toml:"servicekey"`
		NumOfRows   int    `toml:"numOfRows"`
		PageNo      int    `toml:"pageNo"`
		StationName string `toml:"stationName"`
		DataTerm    string `toml:"dataTerm"`
		Ver         string `toml:"ver"`
	} `toml:"openapi_airkorea"`
	SlackAPI struct {
		Token    string `toml:"token"`
		Channel  string `toml:"channel"`
		Username string `toml:"username"`
	} `toml:"slack_api"`
}

type dustinfoResp struct {
	List []struct {
		ReturnType  string `json:"_returnType"`
		CoGrade     string `json:"coGrade"`
		CoValue     string `json:"coValue"`
		DataTerm    string `json:"dataTerm"`
		DataTime    string `json:"dataTime"`
		KhaiGrade   string `json:"khaiGrade"`
		KhaiValue   string `json:"khaiValue"`
		MangName    string `json:"mangName"`
		No2Grade    string `json:"no2Grade"`
		No2Value    string `json:"no2Value"`
		NumOfRows   string `json:"numOfRows"`
		O3Grade     string `json:"o3Grade"`
		O3Value     string `json:"o3Value"`
		PageNo      string `json:"pageNo"`
		Pm10Grade   string `json:"pm10Grade"`
		Pm10Grade1H string `json:"pm10Grade1h"`
		Pm10Value   string `json:"pm10Value"`
		Pm10Value24 string `json:"pm10Value24"`
		Pm25Grade   string `json:"pm25Grade"`
		Pm25Grade1H string `json:"pm25Grade1h"`
		Pm25Value   string `json:"pm25Value"`
		Pm25Value24 string `json:"pm25Value24"`
		ResultCode  string `json:"resultCode"`
		ResultMsg   string `json:"resultMsg"`
		Rnum        int    `json:"rnum"`
		ServiceKey  string `json:"serviceKey"`
		SidoName    string `json:"sidoName"`
		So2Grade    string `json:"so2Grade"`
		So2Value    string `json:"so2Value"`
		StationCode string `json:"stationCode"`
		StationName string `json:"stationName"`
		TotalCount  string `json:"totalCount"`
		Ver         string `json:"ver"`
	} `json:"list"`
	Parm struct {
		ReturnType  string `json:"_returnType"`
		CoGrade     string `json:"coGrade"`
		CoValue     string `json:"coValue"`
		DataTerm    string `json:"dataTerm"`
		DataTime    string `json:"dataTime"`
		KhaiGrade   string `json:"khaiGrade"`
		KhaiValue   string `json:"khaiValue"`
		MangName    string `json:"mangName"`
		No2Grade    string `json:"no2Grade"`
		No2Value    string `json:"no2Value"`
		NumOfRows   string `json:"numOfRows"`
		O3Grade     string `json:"o3Grade"`
		O3Value     string `json:"o3Value"`
		PageNo      string `json:"pageNo"`
		Pm10Grade   string `json:"pm10Grade"`
		Pm10Grade1H string `json:"pm10Grade1h"`
		Pm10Value   string `json:"pm10Value"`
		Pm10Value24 string `json:"pm10Value24"`
		Pm25Grade   string `json:"pm25Grade"`
		Pm25Grade1H string `json:"pm25Grade1h"`
		Pm25Value   string `json:"pm25Value"`
		Pm25Value24 string `json:"pm25Value24"`
		ResultCode  string `json:"resultCode"`
		ResultMsg   string `json:"resultMsg"`
		Rnum        int    `json:"rnum"`
		ServiceKey  string `json:"serviceKey"`
		SidoName    string `json:"sidoName"`
		So2Grade    string `json:"so2Grade"`
		So2Value    string `json:"so2Value"`
		StationCode string `json:"stationCode"`
		StationName string `json:"stationName"`
		TotalCount  string `json:"totalCount"`
		Ver         string `json:"ver"`
	} `json:"parm"`
	ArpltnInforInqireSvcVo struct {
		ReturnType  string `json:"_returnType"`
		CoGrade     string `json:"coGrade"`
		CoValue     string `json:"coValue"`
		DataTerm    string `json:"dataTerm"`
		DataTime    string `json:"dataTime"`
		KhaiGrade   string `json:"khaiGrade"`
		KhaiValue   string `json:"khaiValue"`
		MangName    string `json:"mangName"`
		No2Grade    string `json:"no2Grade"`
		No2Value    string `json:"no2Value"`
		NumOfRows   string `json:"numOfRows"`
		O3Grade     string `json:"o3Grade"`
		O3Value     string `json:"o3Value"`
		PageNo      string `json:"pageNo"`
		Pm10Grade   string `json:"pm10Grade"`
		Pm10Grade1H string `json:"pm10Grade1h"`
		Pm10Value   string `json:"pm10Value"`
		Pm10Value24 string `json:"pm10Value24"`
		Pm25Grade   string `json:"pm25Grade"`
		Pm25Grade1H string `json:"pm25Grade1h"`
		Pm25Value   string `json:"pm25Value"`
		Pm25Value24 string `json:"pm25Value24"`
		ResultCode  string `json:"resultCode"`
		ResultMsg   string `json:"resultMsg"`
		Rnum        int    `json:"rnum"`
		ServiceKey  string `json:"serviceKey"`
		SidoName    string `json:"sidoName"`
		So2Grade    string `json:"so2Grade"`
		So2Value    string `json:"so2Value"`
		StationCode string `json:"stationCode"`
		StationName string `json:"stationName"`
		TotalCount  string `json:"totalCount"`
		Ver         string `json:"ver"`
	} `json:"ArpltnInforInqireSvcVo"`
	TotalCount int `json:"totalCount"`
}
