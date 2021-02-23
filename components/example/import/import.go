// Copyright (c) 2017 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package example

import (
	. "github.com/phsm/pandora/components/example"
	"github.com/phsm/pandora/core/register"
)

func Import() {
	register.Provider("example", NewProvider, DefaultProviderConfig)
	register.Gun("example", NewGun)
}
