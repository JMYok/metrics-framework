package collector

import (
	"interface-metrics/model"
	"interface-metrics/storage"
)

type MetricsCollector struct {
	metricsStorage storage.MetricsStorage //基于接口而非实现
}

func NewMetricsCollector(metricsStorage storage.MetricsStorage) *MetricsCollector {
	return &MetricsCollector{
		metricsStorage: metricsStorage,
	}
}

func (m *MetricsCollector) RecordRequest(info model.RequestInfo) {
	if info == (model.RequestInfo{}) || info.ApiName == "" {
		return
	}
	m.metricsStorage.SaveRequestInfo(info)
}
