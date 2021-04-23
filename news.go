package main
import (
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"encoding/json"
	"io/ioutil"
	"crypto/tls"
//	"time"
)



func initnews(irccon *tls.Conn){
	log.Println("NEWS: Starting...")
	//	Config file check
	if _, err := os.Stat("config.json"); err == nil{
		_, err := os.Open("config.json") 
		if err != nil {log.Fatal(err)}
		log.Println("NEWS: I found a config.json! parsing...")
	}else{log.Fatal("NEWS: I can't find a config file! create a config.json and fill it, exiting...")}

	//	Parsing config.json
	type urlfeed struct{
		Url string
		Tag string
		}
    data, err := ioutil.ReadFile("./config.json")
     if err != nil {log.Fatal(err)}
		
	var urlfeeds []urlfeed
	err = json.Unmarshal([]byte(data), &urlfeeds)
    if err != nil {log.Fatal(err)}

	var feed *gofeed.Feed
		log.Println("NEWS: Parsing...")	
		fp := gofeed.NewParser()

	for i := 0 ; i < len(urlfeeds) ; i++{
		feed, err = fp.ParseURL(urlfeeds[i].Url)
	    if err != nil {log.Fatal(err)}
			for j :=0 ; j < feed.Len() ; j++{
				//privmsg(recipient, content, irccon)
				//irccon.Write([]byte("PRIVMSG #ex1 :\x0304["+urlfeeds[i].Tag+"] "+feed.Items[j].Published+"\r\n"))
				//irccon.Write([]byte("PRIVMSG #ex1 :\x16\x0304"+feed.Items[j].Title+"\r\n"))
				//irccon.Write([]byte("PRIVMSG #ex1 :\x1D"+feed.Items[j].Link+"\r\n"))
				log.Println("["+urlfeeds[i].Tag+"] "+feed.Items[j].Published)
				log.Println(feed.Items[j].Title)
				log.Println(feed.Items[j].Link)
				//time.Sleep(3*time.Second)
			}
	}
	log.Println("Parsed!")
}

func news(sender string,recipient string, message string, irccon *tls.Conn){
}

func newspm(sender string, message string, irccon *tls.Conn){
		  if message[:4] == "list" {
		privmsg(sender, "[general,europe,middleeast,politics,business]", irccon)
		privmsg(sender, "[electronics,android,security,linux]", irccon)
		privmsg(sender, "[alpine,debian,arch,gentoo,openbsd]", irccon)
	}else if message[:9] == "subscribe" {
		
		privmsg(sender, "[alpine,debian,arch,gentoo,openbsd]", irccon)		
	}


}
