package main

import "crypto/tls"

func report(sender string, recipient string, message string, irccon *tls.Conn) {
	if len(message) >= len(prefix+"bots") {
		if message[:5] == prefix+"bots" {
			privmsg(recipient, "Reporting in. [Go]: https://github.com/Azrotronik/GarbageBot/tree/gptbot", irccon)
		}
	}
}

func reportpm(sender string, message string, irccon *tls.Conn) {
	if len(message) >= len(prefix+"test") {
		if message[:4] == "test" {
			privmsg(sender, "Hey!", irccon)
		}
	}
}
