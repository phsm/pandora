// Copyright (c) 2017 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package aggregator

import (
	"context"

	"github.com/phsm/pandora/core"
	"github.com/phsm/pandora/core/coreutil"
)

// NewDiscard returns Aggregator that just throws reported ammo away.
func NewDiscard() core.Aggregator {
	return discard{}
}

type discard struct{}

func (discard) Run(ctx context.Context, _ core.AggregatorDeps) error {
	<-ctx.Done()
	return nil
}

func (discard) Report(s core.Sample) {
	coreutil.ReturnSampleIfBorrowed(s)
}
