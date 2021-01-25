package commands

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

var pingSentences = []string{
	"Pong ! :smirk:",
	"BONJOUR",
	"gildas il sait trop bien grimper putain, théo est jaloux de ouf",
	"t'façon tout le monde sait que Java ça pue",
	"https://github.com/theovidal Meilleur profil GitHub du monde :wink:",
	`On dit "iPhone dix" et non "iPhone ixe" :rage:`,
	"*prend la voix de kernoeb* Heyyy",
	"jaaj",
	"Mais j'ai l'impression de rendre l'aléatoire encore plus aléatoire",
	"OK boomer",
}

func Ping() *onyxcord.Command {
	return &onyxcord.Command{
		Description: "Tester si le robot répond correctement",
		Usage:       "ping",
		ListenInDM:  true,
		Execute: func(arguments []string, bot onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			sentenceNumber := rand.Intn(len(pingSentences))
			_, err = bot.Client.ChannelMessageSend(message.ChannelID, pingSentences[sentenceNumber])
			return
		},
	}
}
