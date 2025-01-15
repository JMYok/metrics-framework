package model

type RequestInfo struct {
	ApiName      string
	ResponseTime float64
	Timestamp    float64
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
