# gotelegrambot

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/gotelegrambot.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/gotelegrambot/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gotelegrambot/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gotelegrambot)](https://goreportcard.com/report/github.com/jjideenschmiede/gotelegrambot) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gotelegrambot?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gotelegrambot) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/gotelegrambot) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

Here you can find our library for telegram bot's. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.

## Install

```console
go get github.com/jjideenschmiede/gotelegrambot
```

## How to use?

Currently we have the following functions covered:

- [Messages](https://github.com/jjideenschmiede/gotelegrambot#messages)
- [Updates](https://github.com/jjideenschmiede/gotelegrambot#updates)
- [Webhooks](https://github.com/jjideenschmiede/gotelegrambot#webhooks)


## Messages

### Create a message

If you want to send a message via a Telegram bot, then you can do it as follows.

```go
// Create a new message
message, err := gotelegrambot.CreateMessage("[J&J Notification][🟠 Warning] Es sind 55 Einträge in der API verzeichnet wurden.", "-1234567", "Markdown", "14241124214:ASDJSKALFSIfjewqrfew234123")
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(message)
}
```

## Updates

### Get updates

If you want to view the updates, you can do that with the following function.

```go
// Get updates
updates, err := gotelegrambot.Updates("14241124214:ASDJSKALFSIfjewqrfew234123")
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updates)
}
```

## Webhooks

### Set a webhook

If you want to create a new webhook, you can use the following function.

```go
// Set webhook
setWebhook, err := gotelegrambot.SetWebhook("https://test.de/webhook", "14241124214:ASDJSKALFSIfjewqrfew234123")
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(setWebhook)
}
```
