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
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
)

type RewriteRule struct {
	regex *regexp.Regexp
	to    string
}

func NewRewriteRule(fromExpr, to string) (*RewriteRule, error) {
	regex, err := regexp.Compile(fromExpr)
	if err != nil {
		return nil, err
	}

	return &RewriteRule{
		regex: regex,
		to:    to,
	}, nil
}

func (r *RewriteRule) Execute(req *http.Request) error {
	uri := req.URL.RequestURI()

	// On no match, no need to execute further.
	if !r.regex.MatchString(uri) {
		fmt.Println("regex goes not match")
		return nil
	}

	// Copying the regex avoids lock contention on concurrent requests.
	regex := r.regex.Copy()
	match := regex.FindStringSubmatchIndex(uri)
	result := regex.ExpandString([]byte{}, r.to, uri, match)
	u, err := url.Parse(path.Clean(string(result)))
	if err != nil {
		return err
	}

	req.URL.Path = u.Path
	req.URL.RawPath = u.RawPath
	req.URL.RawQuery = u.RawQuery

	return nil
}
