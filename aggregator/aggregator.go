package aggregator

import (
	"interface-metrics/model"
	"sort"
)

func Aggregate(requestInfos []model.RequestInfo, durationInMillis int64) *model.RequestStat {
	var maxRespTime float64 = -1 << 63 // Equivalent to Double.MIN_VALUE
	var minRespTime float64 = 1 << 63  // Equivalent to Double.MAX_VALUE
	var avgRespTime float64
	var p999RespTime float64
	var p99RespTime float64
	var sumRespTime float64
	var count int64

	for _, requestInfo := range requestInfos {
		count++
		respTime := requestInfo.ResponseTime
		if maxRespTime < respTime {
			maxRespTime = respTime
		}
		if minRespTime > respTime {
			minRespTime = respTime
		}
		sumRespTime += respTime
	}

	if count != 0 {
		avgRespTime = sumRespTime / float64(count)
	}

	tps := count * 1000 / durationInMillis

	// Sort the requestInfos by response time
	sort.Slice(requestInfos, func(i, j int) bool {
		return requestInfos[i].ResponseTime < requestInfos[j].ResponseTime
	})

	if count != 0 {
		idx999 := int(float64(count) * 0.999) //获得排序位于前 99.9% 位置的相应时间
		idx99 := int(float64(count) * 0.99)   //获得排序位于前 99% 位置的相应时间
		p999RespTime = requestInfos[idx999].ResponseTime
		p99RespTime = requestInfos[idx99].ResponseTime
	}

	requestStat := &model.RequestStat{
		MaxResponseTime:  maxRespTime,
		MinResponseTime:  minRespTime,
		AvgResponseTime:  avgRespTime,
		P999ResponseTime: p999RespTime,
		P99ResponseTime:  p99RespTime,
		Count:            count,
		Tps:              tps,
	}

	return requestStat
}
