package main

import "crypto/tls"

func modules(sender string, recipient string, message string, irccon *tls.Conn) {
	report(sender, recipient, message, irccon)
	news(sender,recipient,message,irccon)

}

func modulespm(sender string, message string, irccon *tls.Conn) {
	reportpm(sender, message, irccon)
	newspm(sender, message, irccon)
}
