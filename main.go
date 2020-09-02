package main

import (
	"github.com/hatchify/errors"
	"github.com/vroomy/common"
	"github.com/yookoala/gofast"
)

var (
	connFactory gofast.ConnFactory
)

const (
	// ErrInvalidFastCGIAddress is returned when the FastCGI address is missing from the Vroomy Environment
	ErrInvalidFastCGIAddress = errors.Error("invalid fast-cgi-addr, cannot be empty")
	// ErrInvalidArguments is returned when an invalid number of arguments is provided
	ErrInvalidArguments = errors.Error("invalid number of arguments, expected filename")
)

// Init is called when Vroomy initializes the plugin
func Init(env map[string]string) (err error) {
	// Set the FastCGI address from the configuration
	address := env["fast-cgi-addr"]
	if len(address) == 0 {
		err = ErrInvalidFastCGIAddress
		return
	}

	// Initialize connection factory for our FastCGI addresss
	connFactory = gofast.SimpleConnFactory("unix", address)
	return
}

// Handler will handle FastCGI requests, it takes the following argument:
//	- filename (e.g. /var/www/html/index.php)
func Handler(args ...string) (h common.Handler, err error) {
	var filename string
	// Get the filename from the provided arguments
	if filename, err = getFilename(args); err != nil {
		return
	}

	// Initialize gofast endpoint for the provided filename
	endpoint := gofast.NewFileEndpoint(filename)
	// Initialize gofast session
	session := endpoint(gofast.BasicSession)
	// Initialize gofast client
	client := gofast.SimpleClientFactory(connFactory, 0)
	// Initialize gofast handler
	handler := gofast.NewHandler(session, client)

	h = newHandler(handler)
	return
}
