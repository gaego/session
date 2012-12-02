// Copyright 2012 The GAEGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package session

import (
	"github.com/gaego/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() {
}

func teardown() {
	context.Close()
}

func TestGetStore(t *testing.T) {
	setup()
	defer teardown()

	r, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	w := httptest.NewRecorder()
	c := context.NewContext(r)

	// create the store
	tStore, err := GetStore(c)
	if err != nil {
		t.Errorf(`err getting store: %q, want nil`, err)
	}

	// Get or Create a new session using the store
	s, err := tStore.Get(r, "t1")
	if err != nil {
		t.Errorf(`err getting session: %q, want nil`, err)
	}

	// Confirm api
	s.Values["val1"] = "example value"
	err = s.Save(r, w)
	if err != nil {
		t.Errorf(`err saving session: %q, want nil`, err)
	}
}
