package main

import "crypto/tls"

func modules(sender string, recipient string, message string, irccon *tls.Conn) {
	report(sender, recipient, message, irccon)
	gpt(sender, recipient, message, irccon)
	
}

func modulespm(sender string, message string, irccon *tls.Conn) {
	reportpm(sender, message, irccon)
	gptpm(sender, message, irccon)
}
