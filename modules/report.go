package main

import "crypto/tls"

func report(sender string, recipient string, message string, irccon *tls.Conn) {
	if message[:5] == prefix+"bots" {
		privmsg(recipient, "Reporting in. [Go]", irccon)
	}
}

func reportpm(sender string, message string, irccon *tls.Conn) {
	if message[:4] == "test" {
		privmsg(sender, "Hey!", irccon)
	}
}
