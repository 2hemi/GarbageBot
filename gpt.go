package main

import (
	"crypto/tls"
	"github.com/imroc/req"
	"log"
	"encoding/json"
	"strings"
)

var uri="Changeme!"
	
func gpt(sender string, recipient string, message string, irccon *tls.Conn) {

	if len(message) >= len(prefix+"QA") {
			if message[:3] == prefix+"QA" {
				privmsg(recipient, gptcomplete(message[3:len(message)-2]), irccon)
			}
		}


	if len(message) >= len(prefix+"complete") {
		if message[:9] == prefix+"complete" {
			privmsg(recipient, gptcomplete(message[10:len(message)-2]), irccon)
		}
	}
}

func gptpm(sender string, message string, irccon *tls.Conn) {
	if len(message) >= len("QA") {
			if message[:2] == "QA" {
				privmsg(sender, gptcompleteqa(message[2:len(message)-2]), irccon)
			}
	}

	if len(message) >= len("complete") {
		if message[:8] == "complete" {
			privmsg(sender, gptcompleteqa(message[9:len(message)-2]), irccon)
		}
	}
}

func gptcomplete(message string) string {

	dictrequest :=`{"prompt": "`+message+`", "seed": 0, "stream": false, "temperature": 1, "top_k": 1000, "top_p": 0.7}`

	r, err := req.Post(uri,"",dictrequest)
	if err !=nil{log.Fatal(err)}


	type res struct{
		Text string
		Reached_end bool
		Total_tokens int
	}

	var response res
	var body string
	body, err=r.ToString()
		if err != nil {log.Fatal(err)}

	json.Unmarshal([]byte(body),&response)
		if err != nil {log.Fatal(err)}
	
log.Print(body)
return message+" "+strings.ReplaceAll(response.Text,"\n"," ")
}

func gptcompleteqa(message string) string {

	dictrequest :=`{"prompt": "Q:`+message+`\nA:", "seed": 0, "stream": false, "temperature": 1, "top_k": 1000, "top_p": 0.7}`

	r, err := req.Post(uri,"",dictrequest)
	if err !=nil{log.Fatal(err)}


	type res struct{
		Text string
		Reached_end bool
		Total_tokens int
	}

	var response res
	var body string
	body, err=r.ToString()
		if err != nil {log.Fatal(err)}

	json.Unmarshal([]byte(body),&response)
		if err != nil {log.Fatal(err)}
	
log.Print(body)
return strings.ReplaceAll(response.Text,"\n"," ")
}
