# dev

A simple static file server for local development. I got tired of `python3 -m http.server`.

## Features

- Support for `404.html`
- On-the-fly directory listing whenever no `index.html` present

## Usage

```
Usage: dev [options]

A simple static file server for local development.

Options:
  -dir string
        Directory to serve files from (default ".")
  -port int
        Port to bind server to (default 8000)
```
