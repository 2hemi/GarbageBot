# GarbageBot
a very slim and simple boilerplate IRC bot.

### Why this over anything else?
This bot is meant to slim (<200LOC), easy to read and use all while having the necessary abstractions to make life easier.

### How it works
This bot reads a line and scans it for triggers, your custom modules execute on each trigger, keeping PMs and channel messages separate, then it goes for the next line.

Create your module.go and drop it in the modules directory, add your function() or functionpm() in modules.go, your trigger must be written inside those functions.

for examples, look at report.go.

### IRC functions to use
Most of these are self explanatory
>func ircconnect(server string, port string, channellist []ircchan, user string, nick string) (*tls.Conn, string)

Starts a TLS connection to a server, identifies and returns the connection pointer

>func ircmaintain(chanlist []ircchan, servername string, irccon *tls.Conn)

maintains the connection, reads lines and sends them for processing

>func pong(linelen int, stream string, conn *tls.Conn)

Responds to PINGs

>func privmsg(recipient string, content string, conn *tls.Conn) 

Sends a message to user/channel

>func handleprivmsg(linelen int, stream string) (sender string, recipient string, content string)

Reads and dissects a message, returns the sender, the recipient and the message content.

>func ircraw(command string, params string, conn *tls.Conn) 

Sends raw data to the server, useful for MODE and other commands.

>func joinchan(chanlist []ircchan, irccon *tls.Conn) 

Joins all the channels in the chanlist[].
