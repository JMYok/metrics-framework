package main

import "time"

func main() {
	metrics := NewMetrics()
	metrics.RecordResponseTime("API1", 100.5)
	metrics.RecordTimestamp("API1", 1633072800)
	metrics.StartRepeatedReport(5 * time.Second)

	time.Sleep(15 * time.Second) // Let the program run for a while to see the output
	metrics.Stop()
}
