package commands

import(
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
	"math/rand"
	"os"
	"bufio"
	"strconv"
	"github.com/PuloV/ics-golang"
	"fmt"
)
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

var coin = []string{
	"HEADS",
	"TAILS",
}

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
				"\n!wank\tProbably the best feature here.\n!deadline\tShow the list of all deadlines.\n!timetable\tShow our timetable\n!meme\tGet help by typing" +
				" !help meme\n!google\tGoogles you any question```"
		}

		return reply, true
	}
	if strings.HasPrefix(m.Content, "!info") {

		reply = "Hello, I am Just a testing bot, created due to improve my master's Go skills."
		return reply, true
	}
	if strings.HasPrefix(m.Content, "!wank") {

		reply =  wankWheel(m)
		return reply, true
	}
	if strings.HasPrefix(m.Content, "!hentai") && strings.Compare(m.ChannelID, "295887491809017858")==0 {

		k := rand.NewSource(time.Now().Unix())
		r := rand.New(k) // initialize local pseudorandom generator
		reply = lines[r.Intn(len(lines))]
		return reply, true
	}else if strings.HasPrefix(m.Content, "!hentai") && strings.Compare(m.ChannelID, "295887491809017858")!=0{
		s.ChannelMessageSend(m.ChannelID, "**@"+ m.Author.Username +"** will be banned cuz hentai should be posted at #nsfw")
		s.GuildMemberRoleAdd("238048572413575168",m.Author.ID,"295907725022593024")
		s.ChannelMessageSend(m.ChannelID,"You are now @Naughty Faggot")
		time.Sleep(time.Minute*5)
		s.GuildMemberRoleRemove("238048572413575168",m.Author.ID,"295907725022593024")
		s.ChannelMessageSend(m.ChannelID,"**@"+ m.Author.Username +"** you are not a @Naughty Faggot any more, but be careful next time. " + m.Author.Token)

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
		s.ChannelMessageSend(m.ChannelID, eightball(text))
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
		s.ChannelMessageSend(m.ChannelID, flip())
	}

	return reply, false
}
func wankWheel(m *discordgo.MessageCreate) string{
	if(strings.Compare(m.Author.ID,"238046128292102145")==0){
		return "Time for some Gay Porn, man"
	}else{
		k := rand.NewSource(time.Now().Unix())
		r := rand.New(k) // initialize local pseudorandom generator
		return "Time for some "+ wankOpions[r.Intn(len(wankOpions))]+" Porn, man"
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
					a=a+fmt.Sprintf("%d) %s\t DEADLINE: %d/%d/%d\n", i,events.GetSummary(), events.GetEnd().Day(),events.GetEnd().Month(),events.GetEnd().Year())
					i++
				}

			}
			a = a + fmt.Sprintf(CODE_HIGHLIGHT)
		}
	}
	return a, i
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

func flip() string{
	k := rand.NewSource(time.Now().Unix())
	r := rand.New(k) // initialize local pseudorandom generator
	return ":regional_indicator_f: :regional_indicator_l: :regional_indicator_i: :regional_indicator_p:: `" + coin[r.Intn(len(coin))] + "`"
}