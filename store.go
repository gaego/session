// Copyright 2012 The AEGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package session provides an interface for Sessions. Currently using
Gorilla sessions

*/
package session

import (
	"errors"
	"sync"

	"appengine"
	"github.com/gaego/config"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var UnknownError = errors.New("an unknown error occured that caused initialization to fail (this could have been initialization failing in another goroutine)")

// This is the store that is used. Note that this can and will be nil until the first successful call of GetStore.
var store sessions.Store

// This contains the default configuration for the SessionKeys Config. The keys for this are generated on initialization.
var defaultConfig = make(map[string]string)

func init() {
	authenticationKeyBytes := securecookie.GenerateRandomKey(64)
	defaultConfig["AuthenticationKey"] = string(authenticationKeyBytes)

	encryptionKeyBytes := securecookie.GenerateRandomKey(32)
	defaultConfig["EncryptionKey"] = string(encryptionKeyBytes)
}

// This is the sync.Once used for initializing the sessions.Store
var initializer = new(sync.Once)

func GetStore(c appengine.Context) (sessions.Store, error) {
	var initializerError error
	initializer.Do(func() {
		storeConfig, err := config.GetOrInsert(c, "SessionKeys", defaultConfig)
		if err != nil {
			initializerError = err
			initializer = new(sync.Once) // We create a new initializer if we error out because instances of sync.Once can not be reset any other way.
			return
		}

		store = sessions.NewCookieStore([]byte(storeConfig.Values["AuthenticationKey"]), []byte(storeConfig.Values["EncryptionKey"]))
	})
	if initializerError != nil { // We've got an error from the initializer.
		return nil, initializerError
	}
	if store == nil { // This more than likely means that store failed to initialize in another goroutine.
		return nil, UnknownError
	}

	return store, nil
}
