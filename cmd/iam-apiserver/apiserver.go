// Copyright 2023 NgSinCung <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"GoPracticalDevelopmentDemo/internel/apiserver"
	"math/rand"
	"time"

	_ "go.uber.org/automaxprocs"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	apiserver.NewApp("iam-apiserver").Run()
}
