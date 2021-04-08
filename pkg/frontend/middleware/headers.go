package middleware

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net/http"
	"strings"
)

func Headers(isDevelopmentMode bool) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			if strings.EqualFold(r.Header.Get("X-Ms-Return-Client-Request-Id"), "true") {
				w.Header().Set("X-Ms-Client-Request-Id", r.Header.Get("X-Ms-Client-Request-Id"))
			}

			if isDevelopmentMode {
				r.Header.Set("Referer", "https://localhost:8443"+r.URL.String())
			}

			h.ServeHTTP(w, r)
		})
	}
}
