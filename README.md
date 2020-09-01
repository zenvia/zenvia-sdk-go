# Zenvia SDK for Go

This SDK for [Golang](https://golang.org/) was created based on the [Zenvia](https://www.zenvia.com/) [API](https://zenvia.github.io/zenvia-openapi-spec/).

## How to test

1) Clone this repo
2) Define zenvia credentials and a number to send test messages:

```
export ZENVIA_USER="USER"
export ZENVIA_PASSWORD="PASSWORD"
export VALID_PHONE="PHONE_TO_SEND_TEST_MESSAGES"
```

3) go test
