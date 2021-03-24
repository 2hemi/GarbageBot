package main
import (
	"log"
	"crypto/tls"
	"strings"
)

type ircmessage struct{
	 sender     string
	 content    string
	 recipient  string
}							
type ircserver  struct{
	 hostname   string
	 port       string
	 channels []ircchan
}
type ircchan    struct{
	 name       string
	 users    []string	
}               
type ircbotid   struct{
	 user	    string
	 nick 	    string
	 rname      string
}

func main(){
	var chanlist  []ircchan	
 chanlist = append(chanlist, ircchan{name:"#ex1"})
 chanlist = append(chanlist, ircchan{name:"#ex2"})
 				  
	exserver    :=ircserver{
				  hostname:"irc.rizon.net",
				  port:    "6697",
				  channels:chanlist} 
	
	exbot		:=ircbotid{
				  user: "GarbageBot",
				  nick: "GarbageBot",
				  rname:"GarbageBot"}

	maincon :=ircconnect(exserver.hostname,exserver.port,exbot.user+" "+exbot.user+" "+exbot.user+" "+exbot.rname,exbot.nick)
	ircmaintain(maincon,exserver.channels)

}

//Connect
func ircconnect(server string,port string,user string,nick string) *tls.Conn{
	log.Println("Gossipbot: Initializing connection.")
    tlsconf := &tls.Config{}
	conn, err := tls.Dial("tcp", server+":"+port,tlsconf)
	if err != nil {log.Fatalln(err)}	
    irccommand("USER",user,conn)
	irccommand("NICK",nick,conn)

   	return conn
}
//Maintain connection, main infinite loop.
func ircmaintain(irccon *tls.Conn,chanlist []ircchan){
	defer irccon.Close()
	stream := make([]byte, 256)
		
	for {
			linelen, err := irccon.Read(stream)
			if err != nil {log.Fatal(err)}
		log.Print(string(stream[:linelen]))
			
			if string(stream[:4]) =="PING" { 				
				pong(linelen,string(stream),irccon)

			}else if strings.Contains(string(stream[:linelen]), " 372 "){
			//	Completely ignore everything in the MOTD as it breaks stuff.
			
			}else if strings.Contains(string(stream[:linelen]), " PRIVMSG #"){ //TODO handle PMs and Messages SEPERATELY
					handleprivmsg(linelen,string(stream),irccon)

			}else if strings.Contains(string(stream[:linelen]), " 001 "){
					irccommand("VERSION","Bot",irccon)
					joinchan(irccon,chanlist)
				
			}
		}	
}
