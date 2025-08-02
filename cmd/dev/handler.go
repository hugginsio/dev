// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"net/http"
	"os"
	"path/filepath"
)

type devHandler struct {
	dir string
}

func (d *devHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(d.dir, r.URL.Path)

	filePath = filepath.Clean(filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		notFoundPath := filepath.Join(d.dir, "404.html")
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

		http.FileServer(http.Dir(d.dir)).ServeHTTP(w, r)
		return
	}

	http.ServeFile(w, r, filePath)
}
