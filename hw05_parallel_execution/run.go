package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := &sync.WaitGroup{}
	ignoreErrors := false
	if m < 0 {
		ignoreErrors = true
	}
	var errorsCount int32
	guard := make(chan struct{}, n)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, task := range tasks {
		wg.Add(1)
		guard <- struct{}{}
		go func(t Task) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				if t() != nil {
					if atomic.AddInt32(&errorsCount, 1) >= int32(m) && !ignoreErrors {
						cancel()
						return
					}
				}
			}
		}(task)
	}
	wg.Wait()
	if errorsCount >= int32(m) && !ignoreErrors {
		return ErrErrorsLimitExceeded
	}
	return nil
}
