package runner

import (
	"errors"
	"time"
)

// Heavily inspired by
// https://github.com/matryer/runner
// https://medium.com/@matryer/stopping-goroutines-golang-1bf28799c1cb

type StopChan <-chan struct{}
type Finish func()

type IRunner interface {
	Wait() error
	Stop()
}

type Runner struct {
	stop    chan struct{}
	stopped chan struct{}
}

func newRunner() *Runner {
	return &Runner{stop: make(chan struct{}, 0), stopped: make(chan struct{}, 0)}
}

// Stop will signal the runner to finish the given goroutine.
// It is not possible to reuse a stopped runner
func (r *Runner) Stop() {
	if r.stop != nil {
		close(r.stop)
	}
}

// Wait will wait for 2 seconds for the runner to stop
// otherwise it returns an error
func (r *Runner) Wait() (err error) {
	select {
	case <-r.stopped:
		return nil
	case <-time.After(2 * time.Second):
		return errors.New("Failed to close in time")
	}
}

// Go starts the given function in a goroutine
// and adds the ability to stop it gracefully.
//
// In your runner listen for the stop channel to be closed,
// then call finish() to stop the runner.
//
// Example:
//	task := Go(func(stop StopChan, finish Finish) {
// 		for {
// 			select {
// 			case _, ok := <-stop:
// 				if !ok {
// 					finish()
// 					return
// 				}
// 			}
// 		}
//	}
func Go(f func(s StopChan, f Finish)) (r *Runner) {
	r = newRunner()

	finish := func() {
		if r.stopped != nil {
			close(r.stopped)
		}
	}
	go f(r.stop, finish)

	return
}
