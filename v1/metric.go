package v1

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"
)

type Metrics struct {
	responseTimes map[string][]float64 // Map的key是接口名称，value对应接口请求的响应时间或时间戳；
	timestamps    map[string][]float64
	executor      *time.Ticker
	mu            sync.Mutex
}

func NewMetrics() *Metrics {
	return &Metrics{
		responseTimes: make(map[string][]float64),
		timestamps:    make(map[string][]float64),
	}
}

func (m *Metrics) RecordResponseTime(apiName string, responseTime float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.responseTimes[apiName]; !exists {
		m.responseTimes[apiName] = []float64{}
	}
	m.responseTimes[apiName] = append(m.responseTimes[apiName], responseTime)
}

func (m *Metrics) RecordTimestamp(apiName string, timestamp float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.timestamps[apiName]; !exists {
		m.timestamps[apiName] = []float64{}
	}
	m.timestamps[apiName] = append(m.timestamps[apiName], timestamp)
}

func (m *Metrics) StartRepeatedReport(period time.Duration) {
	m.executor = time.NewTicker(period)
	go func() {
		for range m.executor.C {
			stats := make(map[string]map[string]float64)
			m.mu.Lock()
			for apiName, apiRespTimes := range m.responseTimes {
				if _, exists := stats[apiName]; !exists {
					stats[apiName] = make(map[string]float64)
				}
				stats[apiName]["max"] = max(apiRespTimes)
				stats[apiName]["avg"] = avg(apiRespTimes)
			}

			for apiName, apiTimestamps := range m.timestamps {
				if _, exists := stats[apiName]; !exists {
					stats[apiName] = make(map[string]float64)
				}
				stats[apiName]["count"] = float64(len(apiTimestamps))
			}
			m.mu.Unlock()

			jsonStats, err := json.Marshal(stats)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				continue
			}
			fmt.Println(string(jsonStats))
		}
	}()
}

func (m *Metrics) Stop() {
	if m.executor != nil {
		m.executor.Stop()
	}
}

func max(dataset []float64) float64 {
	maxValue := math.Inf(-1)
	for _, value := range dataset {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func avg(dataset []float64) float64 {
	var sum float64
	for _, value := range dataset {
		sum += value
	}
	return sum / float64(len(dataset))
}
