package main

import (
	"interface-metrics/collector"
	"interface-metrics/model"
	"interface-metrics/reporter"
	"interface-metrics/storage"
	"time"
)

func main() {
	metricsStorage := storage.NewInMemoryMetricsStorage()
	consoleReporter := reporter.NewConsoleReporter(metricsStorage)
	consoleReporter.StartRepeatedReport(5, 60)

	metricsCollector := collector.NewMetricsCollector(metricsStorage)
	metricsCollector.RecordRequest(model.RequestInfo{
		ApiName:      "register",
		ResponseTime: 123,
		Timestamp:    1736943889368,
	})
	metricsCollector.RecordRequest(model.RequestInfo{
		ApiName:      "register",
		ResponseTime: 223,
		Timestamp:    1736943889368,
	})
	metricsCollector.RecordRequest(model.RequestInfo{
		ApiName:      "register",
		ResponseTime: 323,
		Timestamp:    1736943889368,
	})
	metricsCollector.RecordRequest(model.RequestInfo{
		ApiName:      "login",
		ResponseTime: 23,
		Timestamp:    1736943889368,
	})
	metricsCollector.RecordRequest(model.RequestInfo{
		ApiName:      "login",
		ResponseTime: 1223,
		Timestamp:    1736943889368,
	})

	time.Sleep(15 * time.Second)
	consoleReporter.Stop()
}
