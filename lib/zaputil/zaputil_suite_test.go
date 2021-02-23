// Copyright (c) 2018 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package zaputil

import (
	"testing"

	"github.com/phsm/pandora/lib/ginkgoutil"
)

func TestZaputilSuite(t *testing.T) {
	ginkgoutil.RunSuite(t, "Zaputil Suite")
}
