# Kaban

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)

Kaban is a simple tool for manipulating sprite sheet images.

## Features

- [ ] Create sprite sheet image from individual images
- [x] Unpack sprite sheet image into individual images

## Installation

This package can be utilized either as a CLI tool or as a library.

### As a CLI tool

```sh
go install github.com/aethiopicuschan/kaban@latest
kaban -v
```

### As a library

You can import following package to your project.

```go
import (
  "github.com/aethiopicuschan/kaban/detection"
  "github.com/aethiopicuschan/kaban/types"
)
```

Then run `go mod tidy`.
