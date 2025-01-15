package storage

import "interface-metrics/model"

type MetricsStorage interface {
	SaveRequestInfo(info model.RequestInfo)
	GetRequestInfos(apiName string, startTimeInMillis, endTimeInMillis int64) []model.RequestInfo
	GetRequestInfosByTimeRange(startTimeInMillis, endTimeInMillis int64) map[string][]model.RequestInfo
}
