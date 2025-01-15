package storage

import "interface-metrics/model"

type InMemoryMetricsStorage struct {
	data map[string][]*model.RequestInfo
}

func NewInMemoryMetricsStorage() *InMemoryMetricsStorage {
	return &InMemoryMetricsStorage{
		data: make(map[string][]*model.RequestInfo),
	}
}

func (i *InMemoryMetricsStorage) SaveRequestInfo(info model.RequestInfo) {
	if info.ApiName == "" {
		return
	}
	if _, ok := i.data[info.ApiName]; !ok {
		i.data[info.ApiName] = make([]*model.RequestInfo, 0)
	}
	i.data[info.ApiName] = append(i.data[info.ApiName], &info)
}

func (i *InMemoryMetricsStorage) GetRequestInfos(apiName string, startTimeInMillis, endTimeInMillis int64) []model.RequestInfo {
	if _, ok := i.data[apiName]; !ok {
		return nil
	}
	var result []model.RequestInfo
	for _, info := range i.data[apiName] {
		if info.Timestamp >= float64(startTimeInMillis) && info.Timestamp <= float64(endTimeInMillis) {
			result = append(result, *info)
		}
	}
	return result
}

func (i *InMemoryMetricsStorage) GetRequestInfosByTimeRange(startTimeInMillis, endTimeInMillis int64) map[string][]model.RequestInfo {
	result := make(map[string][]model.RequestInfo)
	for apiName, infos := range i.data {
		for _, info := range infos {
			if info.Timestamp >= float64(startTimeInMillis) && info.Timestamp <= float64(endTimeInMillis) {
				result[apiName] = append(result[apiName], *info)
			}
		}
	}
	return result
}
