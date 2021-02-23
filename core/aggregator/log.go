// Copyright (c) 2017 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package aggregator

import (
	"context"

	"go.uber.org/zap"

	"github.com/phsm/pandora/core"
)

func NewLog() core.Aggregator {
	return &logging{sink: make(chan core.Sample, 128)}
}

type logging struct {
	sink chan core.Sample
	log  *zap.SugaredLogger
}

func (l *logging) Report(sample core.Sample) {
	l.sink <- sample
}

func (l *logging) Run(ctx context.Context, deps core.AggregatorDeps) error {
	l.log = deps.Log.Sugar()
loop:
	for {
		select {
		case sample := <-l.sink:
			l.handle(sample)
		case <-ctx.Done():
			break loop
		}
	}
	for {
		// Context is done, but we should read all data from sink.
		select {
		case r := <-l.sink:
			l.handle(r)
		default:
			return nil
		}
	}
}

func (l *logging) handle(sample core.Sample) {
	l.log.Info("Sample reported: ", sample)
}
