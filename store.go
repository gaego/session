// Copyright 2012 The AEGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package session provides an interface for Sessions. Currently using
Gorilla sessions

*/
package session

import (
	"github.com/gorilla/sessions"
	//"config"
)

// Store is a gorilla/session store.
var Store = sessions.NewCookieStore([]byte("123456789"))

//var Store = sessions.NewCookieStore([]byte(config.SecretKey))
