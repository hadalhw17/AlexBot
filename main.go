package main

import (
	"flag"
	"fmt"
	"strings"
	"bufio"
	"os"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
	readLines("links.txt")
}

var token string
var lines []string
func main() {
	if token == "" {
		fmt.Println("No token provided. Please run: airhorn -t <bot token>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(guildCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("KawaiiBot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	_ = s.UpdateStatus(0, "!info")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!info") {

		s.ChannelMessageSend(m.ChannelID, "Hello, I am Just a testing bot, created due to improve my master's Go skills.")
	}

	if strings.HasPrefix(m.Content, "!wank") {

		s.ChannelMessageSend(m.ChannelID, "Suck my dee------ick, bitch ")
	}

	if strings.HasPrefix(m.Content, "!hentai"){

		k := rand.NewSource(time.Now().Unix())
		r := rand.New(k) // initialize local pseudorandom generator
		s.ChannelMessageSend(m.ChannelID, lines[r.Intn(len(lines))])
		fmt.Println(string(len(lines)))
	}
	if strings.HasPrefix(m.Content, "!help"){
		s.ChannelMessageSend(m.ChannelID, "Available commands: \n!info\tIt will tell you more about bot\n!hentai\tTells why did I chose programming")
	}
}

// This function will be called (due to AddHandler above) every time a new
// guild is joined.
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Kawaii is ready! Type !info to get some usefull shit about me.")
			return
		}
	}
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
