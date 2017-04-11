package games

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"math/rand"
	"time"
	"strings"
)

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

func Eightball(text string) string {
	answer := eightballResponses[rand.Intn(len(eightballResponses))]

	if len(text) > 7 {
		question := text[7:]

		return fmt.Sprintf(":question:`Question:` *%s* \n:8ball:`8Ball answer:` **%s**", question, answer)
	}

	return answer
}

func Flip() string{
	k := rand.NewSource(time.Now().Unix())
	r := rand.New(k) // initialize local pseudorandom generator
	return ":regional_indicator_f: :regional_indicator_l: :regional_indicator_i: :regional_indicator_p:: `" + coin[r.Intn(len(coin))] + "`"
}

func WankWheel(m *discordgo.MessageCreate) string{
	if(strings.Compare(m.Author.ID,"238046128292102145")==0){
		return "Time for some Gay Porn, man"
	}else{
		k := rand.NewSource(time.Now().Unix())
		r := rand.New(k) // initialize local pseudorandom generator
		return "Time for some "+ wankOpions[r.Intn(len(wankOpions))]+" Porn, man"
	}
}
