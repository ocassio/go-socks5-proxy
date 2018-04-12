package main

import (
	"log"
	"os"

	"github.com/armon/go-socks5"
)

func main() {

	creadentials := socks5.StaticCredentials{
		os.Getenv("USER"): os.Getenv("PASSWORD"),
	}
	authenticator := socks5.UserPassAuthenticator{Credentials: creadentials}

	// Create a SOCKS5 server
	config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{authenticator},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}
	server, err := socks5.New(config)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
		panic(err)
	}
}
