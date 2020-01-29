/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package middleware

import (
	"net/http"
	"strings"
)

// CORS is middleware to handle CORS
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		if len(strings.Trim(origin, " ")) == 0 {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		ac := r.Header.Get("access-control-request-headers")
		if len(strings.Trim(ac, " ")) == 0 {
			ac = "*"
		}
		w.Header().Set("access-control-allow-headers", ac)
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// TrimTrailingSlash removes the trailing slash
func TrimTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
