package reporter

import (
	"encoding/json"
	"fmt"
	"interface-metrics/aggregator"
	"interface-metrics/model"
	"interface-metrics/storage"
	"time"
)

type ConsoleReporter struct {
	metricsStorage storage.MetricsStorage
	executor       *time.Ticker
	stopChan       chan bool
}

func NewConsoleReporter(metricsStorage storage.MetricsStorage) *ConsoleReporter {
	return &ConsoleReporter{
		metricsStorage: metricsStorage,
		stopChan:       make(chan bool),
	}
}

func (cr *ConsoleReporter) StartRepeatedReport(periodInSeconds, durationInSeconds int64) {
	cr.executor = time.NewTicker(time.Duration(periodInSeconds) * time.Second)
	go func() {
		for {
			select {
			case <-cr.executor.C:
				durationInMillis := durationInSeconds * 1000
				endTimeInMillis := time.Now().UnixNano() / 1e6
				startTimeInMillis := endTimeInMillis - durationInMillis //计算开始时间

				// 获取某个时间范围内的所有接口请求信息
				requestInfos := cr.metricsStorage.GetRequestInfosByTimeRange(startTimeInMillis, endTimeInMillis)
				stats := make(map[string]*model.RequestStat)
				for apiName, reqInfos := range requestInfos {
					reqStat := aggregator.Aggregate(reqInfos, durationInMillis)
					stats[apiName] = reqStat
				}
				fmt.Printf("Time Span: [%d, %d]\n", startTimeInMillis, endTimeInMillis)
				jsonStats, _ := json.MarshalIndent(stats, "", "  ")
				fmt.Println(string(jsonStats))
			case <-cr.stopChan:
				cr.executor.Stop()
				return
			}
		}
	}()
}

func (cr *ConsoleReporter) Stop() {
	close(cr.stopChan)
}
