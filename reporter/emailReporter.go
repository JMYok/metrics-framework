package reporter

import (
	"fmt"
	"interface-metrics/aggregator"
	"interface-metrics/model"
	"interface-metrics/storage"
	"time"
)

type EmailSender struct{}

func (es *EmailSender) Send(toAddresses []string, subject, body string) error {
	// Implement email sending logic here
	return nil
}

type EmailReporter struct {
	metricsStorage storage.MetricsStorage
	emailSender    *EmailSender
	toAddresses    []string
}

func NewEmailReporter(metricsStorage storage.MetricsStorage, emailSender *EmailSender) *EmailReporter {
	return &EmailReporter{
		metricsStorage: metricsStorage,
		emailSender:    emailSender,
		toAddresses:    []string{},
	}
}

func (er *EmailReporter) AddToAddress(address string) {
	er.toAddresses = append(er.toAddresses, address)
}

func (er *EmailReporter) StartDailyReport() {
	nextRun := getNextMidnight()
	time.AfterFunc(nextRun.Sub(time.Now()), func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				durationInMillis := int64(86400 * 1000)
				endTimeInMillis := time.Now().UnixNano() / 1e6
				startTimeInMillis := endTimeInMillis - durationInMillis
				requestInfos := er.metricsStorage.GetRequestInfosByTimeRange(startTimeInMillis, endTimeInMillis)
				stats := make(map[string]*model.RequestStat)
				for apiName, reqInfos := range requestInfos {
					reqStat := aggregator.Aggregate(reqInfos, durationInMillis)
					stats[apiName] = reqStat
				}
				htmlBody := formatAsHTML(stats)
				subject := "Daily Report"
				er.emailSender.Send(er.toAddresses, subject, htmlBody)
			}
		}
	})
}

func getNextMidnight() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
}
func formatAsHTML(stats map[string]*model.RequestStat) string {
	// Implement HTML formatting logic here
	return fmt.Sprintf("%v", stats)
}
