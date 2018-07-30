/*
Copyright 2018 Frederic Branczyk All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"net/http/httptest"
	"testing"
)

func TestRewrite(t *testing.T) {
	r, err := NewRewriteRule("/federate\\?namespace=([a-z0-9]([-a-z0-9]*[a-z0-9])?)", "/federate?match[]={namespace=$1}")
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("", "/federate?namespace=default", nil)
	err = r.Execute(req)
	if err != nil {
		t.Fatal(req)
	}

	if req.URL.Path != "/federate" {
		t.Fatal("Expected path to be rewritten to /federate, but got: ", req.URL.Path)
	}

	if req.URL.RawQuery != "match[]={namespace=default}" {
		t.Fatal("Expected query to be rewritten to match[]={namespace=default}, but got: ", req.URL.RawQuery)
	}
}
