# go-isemoji
[![go reportcard](https://goreportcard.com/badge/github.com/makeworld-the-better-one/go-isemoji)](https://goreportcard.com/report/github.com/makeworld-the-better-one/go-isemoji)
[![GoDoc](https://godoc.org/github.com/makeworld-the-better-one/go-isemoji?status.svg)](https://godoc.org/github.com/makeworld-the-better-one/go-isemoji)
[![license](https://img.shields.io/github/license/makeworld-the-better-one/go-isemoji)](./LICENSE)


Go library to test if a string is an emoji.

## Usage

```go
isemoji.IsEmoji("🤗")    // True
isemoji.IsEmoji("test")  // False
isemoji.IsEmoji("🙇🏼‍♂️🤗")  // False, because there are multiple emojis
isemoji.Name("👨🏼‍🦰")       // "man: medium-light skin tone, red hair"
```

Find the full documentation on [godoc](https://godoc.org/github.com/makeworld-the-better-one/go-isemoji).

## License
This project is under the MIT License.
