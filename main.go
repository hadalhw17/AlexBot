package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hadalhw17/AlexBot/commands"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

const ANNOUNCEID  = "238048572413575168"
var token string

func main() {
	if token == "" {
		fmt.Println("No token provided. Please run: AlexBot -t <bot token>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}
	//dg.GuildMemberRoleRemove("238048572413575168","299602264039882752","295907725022593024")

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(guildCreate)

	//Register newMember as a callback for the GuildMemberAdd
	dg.AddHandler(newMember)
	//
	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
	//Every 12 hours it will remind about courseworks
	ticker:=time.NewTicker(time.Hour * 12)
	var msg string
	var num int
	msg,num = commands.ReadCal("courseworks.ics")
	go func(){
		for range ticker.C {
			if num!= 0 {
				dg.ChannelMessageSend(ANNOUNCEID, "@everyone you have a lot of courseworks to do!"+"\n"+msg)
			}else {
				dg.ChannelMessageSend(ANNOUNCEID, " no courseworks!! You are free!")
			}
		}

	}()
	fmt.Print("Bot is running.... Thanks god!!")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	_ = s.UpdateStatus(0, "@AlexBot !info|HadalHW17")

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		// Ignore self
		return
	}

	reply, isExist:= commands.ParceForCommands(s,m)
	if isExist{
		s.ChannelMessageSend(m.ChannelID, reply)
	}
	return
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.+
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
		_, _ = s.ChannelMessageSend(ANNOUNCEID, "AlexBot is ready! Type !help to get some usefull shit about me.")
			return
		}
	}
}

// This function will be called (due to AddHandler above) every time a new
// member is joined.+
func newMember(s *discordgo.Session, event *discordgo.GuildMemberAdd){
	s.ChannelMessageSend(ANNOUNCEID, "New member joined. Welcome @" + event.Member.User.Username)

}
