package timer

import (
	"time"

	"demo-ws/common/log"
)

type Timer interface {
	Capture(moment string)
	Print()
}

type timer struct {
	prev   time.Time
	timing map[string]time.Duration
}

func NewTimer() Timer {
	return &timer{prev: time.Now(), timing: map[string]time.Duration{}}
}

func (r *timer) Capture(moment string) {
	r.timing[moment] = time.Since(r.prev)
	r.prev = time.Now()
}

func (r *timer) Print() {
	log.Infof("--- Timing ---")
	for k, v := range r.timing {
		log.Infof("%v: %v", k, v)
	}
	log.Infof("--------------")
}
