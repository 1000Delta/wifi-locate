package utils

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("call failed until timeout")

// CallUntilNoError will do callback and catch error,
// then sleep current goroutine by delay(will be up from initDelay to maxDelay),
// until the callback return nil.
//
// initDelay will by up by 2 power until it great than or equal than maxDelay.
func CallUntilNoError(callback func(currentDelay time.Duration) error, initDelay time.Duration, maxDelay time.Duration) {
	delay := initDelay
	for {
		err := callback(delay)
		if err == nil {
			return
		}
		// 延时当前线程
		time.Sleep(time.Second * delay)
		if delay < maxDelay {
			delay *= 2
		} else if delay > maxDelay {
			delay = maxDelay
		}
	}
}

// CallUntilNoErrorWithTimeout means call CallUntilNoError
// but add a timeout that doing will not be endless
func CallUntilNoErrorWithTimeout(
	callback func(currentDelay time.Duration) error,
	initDelay time.Duration, maxDelay time.Duration, timeout time.Duration) error {
	outTicker := time.NewTicker(timeout * time.Second)
	doneChan := make(chan struct{})

	go CallUntilNoError(func(currentDelay time.Duration) error {
		err := callback(currentDelay)
		if err != nil {
			return err
		}
		// 标记执行完成
		doneChan <- struct{}{}
		return nil
	}, initDelay, maxDelay)

	select {
	case <-outTicker.C:
		return ErrTimeout
	case <-doneChan:
		return nil
	}
}
