package main

import (
	"github.com/Hatch1fy/httpserve"
	"github.com/yookoala/gofast"
)

func getFilename(args []string) (filename string, err error) {
	// Switch on number of args
	switch len(args) {
	case 1:
		// Set filename
		filename = args[0]

	default:
		// Invalid number of argument, return error
		err = ErrInvalidArguments
	}

	return
}

func newHandler(handler gofast.Handler) httpserve.Handler {
	// Wrap gofast.Handler with httpserve Handler
	return func(ctx *httpserve.Context) (res httpserve.Response) {
		// Call handler.ServeHTTP and pass it the writer and request
		handler.ServeHTTP(ctx.Writer, ctx.Request)
		return
	}
}
