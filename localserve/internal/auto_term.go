package internal

import (
	"localserve/localserve/internal/tuned_log"
	"os"
	"sync"
	"syscall"
	"time"
)

type AutoTerminateWatch struct {
	MaxIdleTimeout time.Duration
	currentTime    time.Duration
	mux            sync.Mutex
	startOnce      sync.Once
	TermChan       chan os.Signal
}

func (atw *AutoTerminateWatch) DecTimerBySec() bool {
	time.Sleep(1 * time.Second)
	atw.mux.Lock()
	defer atw.mux.Unlock()
	atw.currentTime -= time.Second
	if atw.currentTime <= 0 {
		return true
	}
	return false
}

func (atw *AutoTerminateWatch) ResetTimer() {
	atw.mux.Lock()
	defer atw.mux.Unlock()
	atw.currentTime = atw.MaxIdleTimeout
}

func (atw *AutoTerminateWatch) theLoop() {
	go func() {
		atw.ResetTimer()
		for {
			if atw.DecTimerBySec() {
				break
			}
		}
		tunedLogger := tuned_log.GetDefaultLogger()
		tuned_log.InfoPrintToUser("\nAuto-Terminating ...", tunedLogger)
		tuned_log.CloseDefaultLogger()
		atw.TermChan <- syscall.SIGTERM
	}()
}

func (atw *AutoTerminateWatch) StartTimerOnce() {
	atw.startOnce.Do(atw.theLoop)
}
