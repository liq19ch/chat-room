package main

import "time"

type message struct {
	Name      string
	Message   string
	when      time.Time
	AvatarURL string
}
