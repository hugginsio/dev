# dev

A better static file server for local development.

Alternatively: I got tired of using `python3 -m http.server`.

## Features

- Support for `404.html`
- On-the-fly directory listing whenever no `index.html` present
- No more `OSError: Address already in use`: it will retry with every port up to 65535

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
