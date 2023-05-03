package util

import (
	"math/rand"
	"time"
)

func RetryFunc(trialTimes, trialInterval int, process func() error) (times int, err error) {
	for trialTimes > times {
		times++
		if err = process(); err == nil {
			return times, nil
		}
		// sleep for random time
		time.Sleep(time.Duration(rand.Intn(trialInterval)) * time.Millisecond)
	}
	return times, err
}
