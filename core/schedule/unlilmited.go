// Copyright (c) 2017 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package schedule

import (
	"time"

	"github.com/phsm/pandora/core"
)

// NewUnlimited returns schedule that generates unlimited ops for passed duration.
func NewUnlimited(duration time.Duration) core.Schedule {
	return &unlimitedSchedule{duration: duration}
}

type UnlimitedConfig struct {
	Duration time.Duration `validate:"min-time=1ms"`
}

func NewUnlimitedConf(conf UnlimitedConfig) core.Schedule {
	return NewUnlimited(conf.Duration)
}

type unlimitedSchedule struct {
	duration time.Duration

	StartSync
	finish time.Time
}

func (s *unlimitedSchedule) Start(startAt time.Time) {
	s.MarkStarted()
	s.startOnce.Do(func() {
		s.finish = startAt.Add(s.duration)
	})
}

func (s *unlimitedSchedule) Next() (tx time.Time, ok bool) {
	s.startOnce.Do(func() {
		s.MarkStarted()
		s.finish = time.Now().Add(s.duration)
	})
	now := time.Now()
	if now.Before(s.finish) {
		return now, true
	}
	return s.finish, false
}

func (s *unlimitedSchedule) Left() int {
	if !s.IsStarted() || time.Now().Before(s.finish) {
		return -1
	}
	return 0
}
