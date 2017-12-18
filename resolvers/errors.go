package nemgraphql

import (
	"errors"
	"log"
	"strings"

	nemrequests "github.com/wzulfikar/go-nem-client/requests"
)

// nemError checks (and return) if given error is originated from NEM network.
// Usage:
// 	if nemErr := nemError(r.Client, err); nemErr != nil {
// 		return nil, nemErr
// 	}
func nemError(c *nemrequests.Client, err error) error {
	strErr := err.Error()
	switch {
	case strings.Contains(strErr, "connection refused"):
		return errors.New("NEM server error: failed to connect to " + c.Endpoint)
	case strings.Contains(strErr, "timeout"):
		return errors.New("NEM server error: connection timeout")
	default:
		return nil
	}
}

func serverError(msg string) error {
	return errors.New(msg)
}

// Handle error and return the correct error msg (or use fallback msg).
// Usage:
// 	if err != nil {
// 		return nil, errorHandler(r.Client, err, "Oops! Something went wrong :(")
// 	}
func errorHandler(c *nemrequests.Client, err error, fallbackMsg string) error {
	log.Println("[ERROR]", err)
	if nemErr := nemError(c, err); nemErr != nil {
		return nemErr
	}
	return serverError(fallbackMsg)
}
