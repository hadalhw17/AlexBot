package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
//	"math/rand"
	"os"
	"bufio"
	"strconv"
	"github.com/PuloV/ics-golang"
	"github.com/hadalhw17/AlexBot/games"
	"github.com/hadalhw17/AlexBot/Hentai"
	"fmt"
)

const NSFW = "295887491809017858"
const GUILDID = "238048572413575168"
const MUTE_ROLE = "295907725022593024"


var CODE_HIGHLIGHT = "```"
var lines []string
func init(){
	readLines("links.txt")
}
func ParceForCommands(s *discordgo.Session, m *discordgo.MessageCreate) (string, bool) {
	var reply string
	var num int
	var msg string
	var text = m.ContentWithMentionsReplaced()
	text=strings.Replace(text,"@everyone","",-1)

	if strings.HasPrefix(m.Content, "!help") {
		if strings.Contains(m.Content, "meme") {
			s.ChannelMessageSend(m.ChannelID, "To generate meme, please type !meme <meme_name> <top_text> <bot_text>\n"+
				"Instead of spaces in <top_text> and <bot_text> use '_'\n"+
				"List of available memes:\nfry\nkermit\nafraid\naag\nblb\nkeanu\nbd\n"+
				"Example: !memes fry top_text_here bot_text_here")
		}else{
			reply = "```Available commands: \n!info\tIt will tell you more about bot\n!hentai\tTells why did I chose programming" +
				"\n!wank\tProbably the best feature here.\n!deadline\tShow the list of all deadlines." +
				"\n!timetable\tShow our timetable\n!meme\tGet help by typing" +
				" !help meme\n!google\tGoogles you any question```"
		}

		return reply, true
	}
	if strings.HasPrefix(m.Content, "!info") {

		reply = "Hello, I am Just a testing bot, created due to improve my master's Go skills."
		return reply, true
	}
	if strings.HasPrefix(m.Content, "!wank") {

		reply =  games.WankWheel(m)
		return reply, true
	}
	if strings.HasPrefix(m.Content, "!hentai") && strings.Compare(m.ChannelID, NSFW)==0 {

		reply = Hentai.GenerateLink()
		fmt.Print("hello")
		return reply, true
	}else if strings.HasPrefix(m.Content, "!hentai") && strings.Compare(m.ChannelID, NSFW)!=0{
		s.ChannelMessageSend(m.ChannelID, "**@"+ m.Author.Username +"** will be banned cuz hentai should be posted at #nsfw")
		s.GuildMemberRoleAdd(GUILDID,m.Author.ID,MUTE_ROLE)
		s.ChannelMessageSend(m.ChannelID,"You are now @Naughty Faggot")
		time.Sleep(time.Minute*5)
		s.GuildMemberRoleRemove(GUILDID,m.Author.ID,MUTE_ROLE)
		s.ChannelMessageSend(m.ChannelID,"**@"+ m.Author.Username +"** you are not a @Naughty Faggot" +
			" any more, but be careful next time. " + m.Author.Token)
		return "",true
	}

	if strings.HasPrefix(m.Content, "!deadline"){
		msg,num = ReadCal("courseworks.ics")
		reply = "You have "+ strconv.Itoa(num-1) + " courseworks to do:\n"+msg
		return reply, true
	}

	if strings.HasPrefix(m.Content, "!timetable"){
		reply = "Here you go, wanker http://imgur.com/a/jx8N2"
		return reply,true
	}

	if strings.HasPrefix(m.Content, "!8ball"){
		s.ChannelMessageSend(m.ChannelID, games.Eightball(text))
	}

	if strings.HasPrefix(m.Content, "@everyone you have a lot of courseworks to do!"){
		pinLastMessage(m,s)

		return "", true
	}


	if strings.HasPrefix(m.Content, "!meme"){
		//fmt.Print(getMemes("","",""))
		reply = getMemes(text)
		return reply, true
	}

	if strings.HasPrefix(m.Content, "!google"){
		reply = googleForMe(text)
		return reply, true
	}

	if strings.HasPrefix(m.Content, "!flip"){
		s.ChannelMessageSend(m.ChannelID, games.Flip())
	}

	return reply, false
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

func ReadCal(path string) (string, int){
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
			a= a+ fmt.Sprintf(CODE_HIGHLIGHT)
			for _, events := range event {
				// print event info (event implements Stringer interface)
				if time.Now().Before(events.GetEnd()){
					a=a+fmt.Sprintf("%d) %s\t DEADLINE: %d/%d/%d\n",
						i,events.GetSummary(), events.GetEnd().Day(),events.GetEnd().Month(),events.GetEnd().Year())
					i++
				}

			}
			a = a + fmt.Sprintf(CODE_HIGHLIGHT)
		}
	}
	return a, i
}

func pinLastMessage(m *discordgo.MessageCreate, s *discordgo.Session){
	s.ChannelMessagePin(m.ChannelID, m.ID)
}

func getMemes(text string) (resp string){
	id := strings.Fields(strings.TrimSpace(text))
	if len(text)<6{
		resp = "To generate a meme, you need to type `!meme <name> <top_text> <bot_text>"
	}else{
		if len(id)<4 {
			return "Look !help meme to get help or just call ambulance and tell that you are mentally ill"
		}
		id[1] = strings.Replace(id[1],"_", "-",1)
		id[2] = strings.Replace(id[2],"_", "-",1)
		resp = "https://memegen.link/"+id[1]+"/"+id[2]+"/"+id[3]+".jpg"
	}

	return resp
}

func googleForMe(text string) string{
	tmp:= text[8:]
	link:=strings.Replace(tmp," ", "+", 1000)
	request:="http://lmgtfy.com/?q="+link
	return request

}