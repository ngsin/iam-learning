// Copyright 2023 Xiansong Wu <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import cliflag "github.com/marmotedu/component-base/pkg/cli/flag"

type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
}