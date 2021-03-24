package main
import (
	"log"
	"crypto/tls"
	"strings"
)

//Respond to PINGs
func pong(linelen int,stream string,conn *tls.Conn){
	conn.Write([]byte("PONG "+stream[strings.Index(stream," :")+1:linelen]+"\r\n"))
	log.Print("PONG "+stream[strings.Index(stream," :")+1:linelen])
}
//Respond to PRIVMSG
func handleprivmsg(linelen int,stream string,conn *tls.Conn){
	var m ircmessage
	m.sender=stream[1: strings.Index(stream,"!")]
	m.recipient=stream[strings.Index(stream,"PRIVMSG ")+7:strings.Index(stream," :")]
	m.content=stream[strings.Index(stream," :")+2:linelen]
}
//Send PRIVMSG
func privmsg(recipient string, content string,conn *tls.Conn){
	log.Print("PRIVMSG "+recipient+" :"+content+"\r\n")
	_,err:= conn.Write( []byte("PRIVMSG "+recipient+" :"+content+"\r\n") )	
	if err != nil {log.Fatalln(err)}

}
//Execute command
func irccommand(command string, params string,conn *tls.Conn){
	_,err:= conn.Write([]byte(command+" "+params+"\r\n"))	
	if err != nil {log.Fatalln(err)}

}
//Join channel
func joinchan(irccon *tls.Conn, chanlist []ircchan){
	for i:=0;i<len(chanlist);i++ {
					irccommand("JOIN",chanlist[i].name,irccon)
					privmsg(chanlist[i].name,"Ready.",irccon)
					log.Println("GarbageBot: I'm in "+chanlist[i].name)
					}//TODO Fill out NAMES from NAMES list
	
}

