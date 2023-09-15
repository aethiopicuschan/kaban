# Kaban

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/kaban.svg)](https://pkg.go.dev/github.com/aethiopicuschan/kaban)

Kaban is a simple tool for manipulating sprite sheet images.

## Features

- Create sprite sheet image from individual images
- Unpack sprite sheet image into individual images


## Example

### Unpack

```sh
kaban unpack ./example/example.png
```

![example.png](/example/example.png)
→
![0_1__67_126.png](/example/0_1__67_126.png)
&
![76_0__128_128.png](example/76_0__128_128.png)

### Pack

```sh
kaban pack ./example/0_1__67_126.png ./example/76_0__128_128.png -o ./example/packed.png
```

![0_1__67_126.png](/example/0_1__67_126.png)
+
![76_0__128_128.png](example/76_0__128_128.png)
→
![packed.png](/example/packed.png)

## Installation

This package can be utilized either as a CLI tool or as a library.

### As a CLI tool

```sh
go install github.com/aethiopicuschan/kaban@latest
kaban -h
```

### As a library

```sh
go get github.com/aethiopicuschan/kaban
```

```go
import (
	"github.com/aethiopicuschan/kaban/detection"
	"github.com/aethiopicuschan/kaban/merge"
)
```
