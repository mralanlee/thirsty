package main

type Commands interface {
	auth() commands
}

// ADD command and flag router
