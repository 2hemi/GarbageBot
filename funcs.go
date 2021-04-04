package main
import (
	"log"
	"crypto/tls"
	"strings"
)

func handleline(linelen int, stream string, servername string, irccon *tls.Conn){
	log.Print(stream[:linelen])

	if string(stream[:4]) == "PING" {pong(linelen,string(stream),irccon)}
	
	if strings.Contains(stream[:2+len(servername)],":"+servername+" "){
		if strings.Contains(stream[2+len(servername):linelen],"001 "){
			handlecode(001,"",linelen,irccon)
		}else if strings.Contains(stream[2+len(servername):linelen],"433 "){
			handlecode(433,"",linelen,irccon)
		}else if strings.Contains(stream[2+len(servername):linelen],"353 "){
			handlecode(353,stream,linelen,irccon)
		}
		
	}

/*	if strings.Contains(stream[2+len(servername):linelen],"PRIVMSG GarbageBot :"){
		handleprivmsg(linelen,stream)
	}*/

}
//Respond to server codes
func handlecode(code int,stream string,linelen int, irccon *tls.Conn){
   switch code {
      	case 001:
			irccommand("VERSION","Bot",irccon)
      		joinchan(irccon,chanlist);      	
      	case 433:
      		log.Fatal("GarbageBot: it seems that someone is sqatting this nick, change it and reconnect.");
		case 353:
      		for strings.Contains(stream[2+len(servername):linelen],"366 "){
				//TODO irccon read and slice
			};
   }
}
//Connect to a server
func ircconnect(server string,port string, chanlist []ircchan,user string,nick string) (*tls.Conn,string){
	log.Println("GarbageBot: Initializing connection.")
	stream := make([]byte, 256)	
    tlsconf := &tls.Config{}
	irccon, err := tls.Dial("tcp", server+":"+port,tlsconf)
	if err != nil {log.Fatalln(err)}	

	linelen, err := irccon.Read(stream)
	if err != nil {log.Fatal(err)}
	servername := string(stream[1:strings.Index(string(stream[:linelen])," ")])	

	irccommand("USER",user,irccon)
	irccommand("NICK",nick,irccon)

   	return irccon, servername
}
//Maintain the connection, main infinite loop.
func ircmaintain(irccon *tls.Conn,chanlist []ircchan, servername string){
	defer irccon.Close()
	stream := make([]byte, 256)	
	for {
		linelen, err := irccon.Read(stream)
		if err != nil {log.Fatal(err)}
		handleline(linelen, string(stream), servername, irccon)		
	}	
}

//Respond to PINGs
func pong(linelen int,stream string,conn *tls.Conn){
	conn.Write([]byte("PONG "+stream[strings.Index(stream," :")+1:linelen]+"\r\n"))
	log.Print("PONG "+stream[strings.Index(stream," :")+1:linelen])
}
//Respond to PRIVMSG
func handleprivmsg(linelen int,stream string){
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
	}
	
}

