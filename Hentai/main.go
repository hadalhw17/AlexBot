package Hentai

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateLink() string{
	responce:= "https://danbooru.donmai.us/posts/"
	return responce+strconv.Itoa(random(2687405,2688326))
}
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}