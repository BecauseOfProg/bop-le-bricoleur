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
		ListenInDM: true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			sentenceNumber := rand.Intn(len(pingSentences))
			_ = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: pingSentences[sentenceNumber],
				},
			})
			return
		},
	}
}
