// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

// Handler provides a custom HttpHandler for responding to HTTP requests.
package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

// ServeHandler is a custom HttpHandler with 404.html support and automatic directory browser generation.
type ServeHandler struct {
	Directory string // The root directory to serve content from.
}

func (m *ServeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(m.Directory, r.URL.Path)

	filePath = filepath.Clean(filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		notFoundPath := filepath.Join(m.Directory, "404.html")
		if content, err := os.ReadFile(notFoundPath); err == nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			w.Write(content)
			return
		}

		http.NotFound(w, r)
		return
	}

	if info, err := os.Stat(filePath); err == nil && info.IsDir() {
		indexPath := filepath.Join(filePath, "index.html")
		if _, err := os.Stat(indexPath); err == nil {
			http.ServeFile(w, r, indexPath)
			return
		}

		http.FileServer(http.Dir(m.Directory)).ServeHTTP(w, r)
		return
	}

	http.ServeFile(w, r, filePath)
}
