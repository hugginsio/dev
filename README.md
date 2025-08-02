# dev

A better static file server for local development.

Alternatively: I got tired of using `python3 -m http.server`.

## Features

- Support for `404.html`
- On-the-fly directory listing whenever no `index.html` present
- No more `OSError: Address already in use`: it will retry with every port up to 65535

## Installation

`dev` is available in Homebrew:

```
brew install hugginsio/tap/dev
```

Alternatively, you can download binaries for your platform from [the latest release](https://github.com/hugginsio/dev/releases/latest).

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
