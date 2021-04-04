package main

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
	 users    []ircuser	
}               
type ircbotid   struct{
	 user	    string
	 nick 	    string
	 rname      string
}
  
type ircuser    struct{
	 nick 	   	string
	 rname      string
	 prefix		byte
}

var (chanlist  	[]ircchan	  
	exserver	ircserver
	exbot		ircbotid
	servername	string)

func main(){

	chanlist = append(chanlist, ircchan{name:"#ex1"})
	chanlist = append(chanlist, ircchan{name:"#ex2"})
	exserver    :=ircserver{
				  hostname:"irc.rizon.net",
				  port:    "6697",
				  channels:chanlist} 
	exbot		:=ircbotid{
				  user: "Garbage",
				  nick: "Garbage",
				  rname:"Garbage"}
	maincon,servername :=ircconnect(exserver.hostname,exserver.port, exserver.channels, exbot.user+" "+exbot.user+" "+exbot.user+" "+exbot.rname,exbot.nick)
	ircmaintain(maincon,exserver.channels,servername)
}

