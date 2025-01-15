package storage

type RedisStorage struct {
}

func (r *RedisStorage) RecordMetric(key string, value float64) error {
	// Implement the method
	return nil
}

func (r *RedisStorage) GetMetric(key string) (float64, error) {
	// Implement the method
	return 0, nil
}
