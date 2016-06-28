# go-base58

A Base58 encoding and decoding package for [Go](https://golang.org/).

## What?

Base58 allows you to represent a numeric value with fewer characters, useful
for short URLs among other things. Flickr is one the biggest sites that makes
use of it for short photo URLs.

For example `6857269519` becomes `brXijP` when Base58 encoded, and hence the
Flickr short URL is: `http://flic.kr/p/brXijP`

## Installation

    go get https://github.com/jimeh/go-base58

## Usage

```go
import "github.com/jimeh/go-base58"

uid := base58.Encode(6857269519)
fmt.Println(uid) // Prints: brXijP

num, err := base58.Decode('brXijP');
fmt.Println(num) // Prints: 6857269519
```

## Credit

This package is more or less a port of the [Base58][gem] Ruby Gem by the same
name. Which as far as I can tell, is based on [this article][article] by Flickr
staff.

[gem]: https://github.com/dougal/base58
[article]: https://www.flickr.com/groups/51035612836@N01/discuss/72157616713786392/

## License

Released under the MIT license. Copyright (c) 2016 Jim Myhrberg.
