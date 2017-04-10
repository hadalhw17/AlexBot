package main

import (
	"flag"
	"fmt"
	"strings"
	"bufio"
	"os"
	"math/rand"

	"time"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/PuloV/ics-golang"

)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
	readLines("links.txt")
}

var token string
var lines []string
var anounced = false
var event []string
var eightballResponses = []string{
	"Most definitely yes",
	"For sure",
	"As I see it, yes",
	"My sources say yes",
	"Yes",
	"According to BBC, your mother is a whore",
	"Most likely",
	"Signs point to yes",
	"Perhaps",
	"Maybe",
	"Reply hazy, try again",
	"Ask again later ",
	"Better not tell you now",
	" Cannot predict now",
	"Concentrate and ask again",
	"Not sure",
	"It is uncertain",
	"Ask me again later",
	"Don't count on it",
	"Probably not",
	"Very doubtful",
	"Most likely no",
	"Nope",
	"No",
	"My sources say no",
	"Dont even think about it",
	"Definitely no",
	"NO, YOU WHORE!!",
	"NO - It may cause disease contraction",
}
var wankOpions = []string{
	"Amateur",
	"American",
	"Anal Sex",
	"Hentai",
	"BDSM",
	"Beach Sex",
	"Blowjob",
	"Creampie",
	"Deapthroat",
	"Ebony",
	"German",
	"Hardcore",
	"Indian",
	"Milf",
	"Orgy",
	"Public",
	"Redhead",
	"Retro",
	"Russian",
	"School",
	"Shemale",
	"Jap",
	"Teacher",
	"Latin",
	"Solo Girls",
	"Gay",
}
var coin = []string{
	"HEADS",
	"TAILS",
}
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
	//dg.GuildMemberRoleRemove("238048572413575168","299602264039882752","295907725022593024")
	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Register guildCreate as a callback for the guildCreate events.
	dg.AddHandler(guildCreate)

	dg.AddHandler(newMember)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
	ticker:=time.NewTicker(time.Hour * 12)
	var msg string
	var num int
	msg,num = readCal("courseworks.ics")
	go func(){
		for range ticker.C {
			if num!= 0 {
				dg.ChannelMessageSend("238048572413575168", "@everyone you have a lot of courseworks to do!"+"\n"+msg)
			}else {
				dg.ChannelMessageSend("238048572413575168", " no courseworks!! You are free!")
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
	var text = m.ContentWithMentionsReplaced()
	text=strings.Replace(text,"@everyone","",-1)
	if strings.HasPrefix(m.Content, "!info") {

		s.ChannelMessageSend(m.ChannelID, "Hello, I am Just a testing bot, created due to improve my master's Go skills.")
	}

	if strings.HasPrefix(m.Content, "!wank") {

		s.ChannelMessageSend(m.ChannelID, wankWheel(m))
	}

	if strings.HasPrefix(m.Content, "!hentai") && strings.Compare(m.ChannelID, "295887491809017858")==0 {

		k := rand.NewSource(time.Now().Unix())
		r := rand.New(k) // initialize local pseudorandom generator
		s.ChannelMessageSend(m.ChannelID, lines[r.Intn(len(lines))])
		fmt.Println(string(len(lines)))
	}else if strings.HasPrefix(m.Content, "!hentai") && strings.Compare(m.ChannelID, "295887491809017858")!=0{
		s.ChannelMessageSend(m.ChannelID, "**@"+ m.Author.Username +"** will be banned cuz hentai should be posted at #nsfw")
		s.GuildMemberRoleAdd("238048572413575168",m.Author.ID,"295907725022593024")
		s.ChannelMessageSend(m.ChannelID,"You are now @Naughty Faggot")
		time.Sleep(time.Minute*5)
		s.GuildMemberRoleRemove("238048572413575168",m.Author.ID,"295907725022593024")
		s.ChannelMessageSend(m.ChannelID,"**@"+ m.Author.Username +"** you are not a @Naughty Faggot any more, but be careful next time. " + m.Author.Token)
	}
	if strings.HasPrefix(m.Content, "!help"){
		if strings.Contains(m.Content, "meme"){
			s.ChannelMessageSend(m.ChannelID, "To generate meme, please type !meme <meme_name> <top_text> <bot_text>\n" +
				"Instead of spaces in <top_text> and <bot_text> use '_'\n" +
				"List of available memes:\nfry\nkermit\nafraid\naag\nblb\nkeanu\nbd\n" +
				"Example: !memes fry top_text_here bot_text_here")
		}else {
			s.ChannelMessageSend(m.ChannelID, "```Available commands: \n!info\tIt will tell you more about bot\n!hentai\tTells why did I chose programming"+
				"\n!wank\tProbably the best feature here.\n!deadline\tShow the list of all deadlines.\n!timetable\tShow our timetable\n!meme\tGet help by typing" +
				" !help meme\n!google\tGoogles you any question```")
		}
	}
	var num int
	var msg string
	msg,num = readCal("courseworks.ics")
	if strings.HasPrefix(m.Content, "!deadline"){
		s.ChannelMessageSend(m.ChannelID, "You have "+ strconv.Itoa(num-1) + " courseworks to do:\n"+msg)
	}
	if strings.HasPrefix(m.Content, "!timetable"){
		s.ChannelMessageSend(m.ChannelID, "Here you go, wanker http://imgur.com/a/jx8N2")
	}
	if strings.HasPrefix(m.Content, "!8ball"){
		s.ChannelMessageSend(m.ChannelID, eightball(text))
	}
	if strings.HasPrefix(m.Content, "@everyone you have a lot of courseworks to do!"){
		pinLastMessage(m,s)
	}
	if strings.HasPrefix(m.Content, "!meme"){
		//fmt.Print(getMemes("","",""))
		s.ChannelMessageSend(m.ChannelID, getMemes(text))
	}
	if strings.HasPrefix(m.Content, "!google"){
		fmt.Print("Hey "+googleForMe(text))
		s.ChannelMessageSend(m.ChannelID,googleForMe(text))
	}
	if strings.HasPrefix(m.Content, "!flip"){
		s.ChannelMessageSend(m.ChannelID, flip())
	}

}
// This function will be called (due to AddHandler above) every time a new
// guild is joined.+
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			//_, _ = s.ChannelMessageSend(channel.ID, "Kawaii is ready! Type !info to get some usefull shit about me.")
			return
		}
	}
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}
func newMember(s *discordgo.Session, event *discordgo.GuildMemberAdd){
	s.ChannelMessageSend("238048572413575168", "New member joined. Welcome @" + event.Member.User.Username)
	fmt.Print("New member joined. Welcome *@" + event.Member.User.Username+"*")

}
func readCal(path string) (string, int){
	var a = ""
	var i = 1
	parser := ics.New()
	parserChan := parser.GetInputChan()
	parserChan <- path

	// wait to kill the main goroute
	parser.Wait()
	cal, err := parser.GetCalendars()
	if err == nil {
		for _, calendar := range cal {
			event := calendar.GetEvents()
			a= a+ fmt.Sprintf("```")
			for _, events := range event {
				// print event info (event implements Stringer interface)
				if time.Now().Before(events.GetEnd()){
					a=a+fmt.Sprintf("%d) %s\t DEADLINE: %d/%d/%d\n", i,events.GetSummary(), events.GetEnd().Day(),events.GetEnd().Month(),events.GetEnd().Year())
					i++
				}

			}
			a = a + fmt.Sprintf("```")
		}
	}
	return a, i
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

func eightball(text string) string {
	answer := eightballResponses[rand.Intn(len(eightballResponses))]

	if len(text) > 7 {
		question := text[7:]

		return fmt.Sprintf(":question:`Question:` *%s* \n:8ball:`8Ball answer:` **%s**", question, answer)
	}

	return answer
}

func pinLastMessage(m *discordgo.MessageCreate, s *discordgo.Session){
	s.ChannelMessagePin(m.ChannelID, m.ID)
}

func getMemes(text string) (resp string){


	id := strings.Fields(strings.TrimSpace(text))
	id[1] = strings.Replace(id[1],"_", "-",1)
	id[2] = strings.Replace(id[2],"_", "-",1)
	resp = "https://memegen.link/"+id[1]+"/"+id[2]+"/"+id[3]+".jpg"
	return resp
}

func googleForMe(text string) string{
	tmp:= text[8:]
	link:=strings.Replace(tmp," ", "+", 1000)
	request:="http://lmgtfy.com/?q="+link
	return request

}

func wankWheel(m *discordgo.MessageCreate) string{
	if(strings.Compare(m.Author.ID,"238046128292102145")==0){
		return "Time for some Gay Porn, man"
	}else{
		k := rand.NewSource(time.Now().Unix())
		r := rand.New(k) // initialize local pseudorandom generator
		return "Time for some "+ wankOpions[r.Intn(len(lines))]+" Porn, man"
	}
}
func flip() string{
	k := rand.NewSource(time.Now().Unix())
	r := rand.New(k) // initialize local pseudorandom generator
	return ":regional_indicator_f: :regional_indicator_l: :regional_indicator_i: :regional_indicator_p:: `" + coin[r.Intn(len(coin))] + "`"
}
