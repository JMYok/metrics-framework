package model

type RequestInfo struct {
	ApiName      string
	ResponseTime float64 //响应时间
	Timestamp    float64 //触发时间戳
}

type RequestStat struct {
	MaxResponseTime  float64
	MinResponseTime  float64
	AvgResponseTime  float64
	P999ResponseTime float64
	P99ResponseTime  float64
	Count            int64
	Tps              int64
}
