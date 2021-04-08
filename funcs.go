package main

import (
	"crypto/tls"
	"log"
	"strings"
)

func handleline(linelen int, stream string, servername string, irccon *tls.Conn) {
	log.Print(stream[:linelen])
	if string(stream[:4]) == "PING" {
		pong(linelen, string(stream), irccon)
	}

	if strings.Contains(stream[:2+len(servername)], ":"+servername+" ") {
		if strings.Contains(stream[1:linelen], servername+" 001") {
			handlecode(001, "", linelen, irccon)
		} else if strings.Contains(stream[1:linelen], servername+" 433") {
			handlecode(433, "", linelen, irccon)
		}
	}

	if strings.Contains(stream[:linelen], " PRIVMSG Garbage :") {
		sen, _, mes := handleprivmsg(linelen, stream)
		modulespm(sen, mes, irccon)
	} else if strings.Contains(stream[:linelen], " PRIVMSG #") {
		sen, rec, mes := handleprivmsg(linelen, stream)
		modules(sen, rec, mes, irccon)
	}

}

//Respond to server codes
func handlecode(code int, stream string, linelen int, irccon *tls.Conn) {
	switch code {
	case 001:
		ircraw("VERSION", "Bot", irccon)
		joinchan(irccon, chanlist)
	case 433:
		log.Fatal("GarbageBot: it seems that someone is sqatting this nick, change it and reconnect.")
	}
}

//Connect to a server
func ircconnect(server string, port string, chanlist []ircchan, user string, nick string) (*tls.Conn, string) {
	log.Println("GarbageBot: Initializing connection.")
	stream := make([]byte, 256)
	tlsconf := &tls.Config{}
	irccon, err := tls.Dial("tcp", server+":"+port, tlsconf)
	if err != nil {
		log.Fatalln(err)
	}

	linelen, err := irccon.Read(stream)
	if err != nil {
		log.Fatal(err)
	}
	servername := string(stream[1:strings.Index(string(stream[:linelen]), " ")])

	ircraw("USER", user, irccon)
	ircraw("NICK", nick, irccon)

	return irccon, servername
}

//Maintain the connection, main infinite loop.
func ircmaintain(chanlist []ircchan, servername string, irccon *tls.Conn) {
	defer irccon.Close()
	stream := make([]byte, 256)
	for {
		linelen, err := irccon.Read(stream)
		if err != nil {
			log.Fatal(err)
		}
		handleline(linelen, string(stream), servername, irccon)
	}
}

//Respond to PINGs
func pong(linelen int, stream string, conn *tls.Conn) {
	conn.Write([]byte("PONG " + stream[strings.Index(stream, " :")+1:linelen] + "\r\n"))
	log.Print("PONG " + stream[strings.Index(stream, " :")+1:linelen])
}

//Respond to PRIVMSG
func handleprivmsg(linelen int, stream string) (sender string, recipient string, content string) {
	var m ircmessage
	m.sender = stream[1:strings.Index(stream, "!")]
	m.recipient = stream[strings.Index(stream, "PRIVMSG ")+7 : strings.Index(stream, " :")]
	m.content = stream[strings.Index(stream, " :")+2 : linelen]
	return m.sender, m.recipient, m.content
}

//Send PRIVMSG
func privmsg(recipient string, content string, conn *tls.Conn) {
	log.Print("PRIVMSG " + recipient + " :" + content + "\r\n")
	_, err := conn.Write([]byte("PRIVMSG " + recipient + " :" + content + "\r\n"))
	if err != nil {
		log.Fatalln(err)
	}
}

//Execute command
func ircraw(command string, params string, conn *tls.Conn) {
	_, err := conn.Write([]byte(command + " " + params + "\r\n"))
	if err != nil {
		log.Fatalln(err)
	}
}

//Join channel
func joinchan(chanlist []ircchan, irccon *tls.Conn,) {
	for i := 0; i < len(chanlist); i++ {
		ircraw("JOIN", chanlist[i].name, irccon)
		log.Println("GarbageBot: I'm in " + chanlist[i].name)
	}
}
