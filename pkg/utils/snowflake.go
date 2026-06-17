package utils

import (
	"time"
)

type SnowflakeIDGenerator struct {
	timestamp int64
	workerID  int64
	counter   int64
}

func NewSnowflakeIDGenerator(workerID int64) *SnowflakeIDGenerator {
	return &SnowflakeIDGenerator{
		timestamp: time.Now().UnixNano() / 1000000,
		workerID:  workerID,
		counter:   0,
	}
}

func (s *SnowflakeIDGenerator) Generate() int64 {
	timestamp := time.Now().UnixNano() / 1000000

	if timestamp < s.timestamp {
		panic("Clock moved backwards")
	}

	if timestamp == s.timestamp {
		s.counter++
		if s.counter > 1023 {
			time.Sleep(1 * time.Millisecond)
			return s.Generate()
		}
	} else {
		s.counter = 0
	}

	s.timestamp = timestamp

	id := (timestamp-1288834974657)*1024*1024 + s.workerID*1024 + s.counter
	return id
}
