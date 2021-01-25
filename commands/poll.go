package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord"
)

var pollChoices = map[string][]string{
	"shapes":         strings.Split("🔴 🟤 🟠 🟣 🟡 🔵 🟢 ⚫ ⚪ 🟥 🟫 🟧 🟪 🟨 🟦 🟩 ⬛ ⬜ 🔶 🔺", " "),
	"numbers":        strings.Split("0️⃣ 1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣ 8️⃣ 9️⃣ 🇦 🇧 🇨 🇩 🇪 🇫 🇬 🇭 🇮", " "),
	"letters":        strings.Split("🇦 🇧 🇨 🇩 🇪 🇫 🇬 🇭 🇮 🇯 🇰 🇱 🇲 🇳 🇴 🇵 🇶 🇷 🇸 🇹", " "),
	"food":           strings.Split("🍎 🍍 🍇 🥐 🥗 🥪 🍕 🥓 🍜 🥘 🍧 🍩 🍰 🍬 🍭 ☕ 🧃 🍵 🍾 🍸", " "),
	"faces":          strings.Split("😄 😋 😎 😂 🥰 😎 🤔 🙄 😑 🤨 😮 😴 😛 😤 🤑 😭 😨 🥵 🥶 😷", " "),
	"animals":        strings.Split("🐔 🐴 🐸 🐷 🐗 🐰 🐹 🦊 🐶 🐼 🦓 🐁 🐘 🐢 🐍 🐳 🦐 🐠 🦢 🦜", " "),
	"transportation": strings.Split("🚗 🚓 🚌 🚚 🚜 🚅 🚋 🚇 🚠 ✈ 🚁 🚀 🚢 🛹 🚲 🛴 🛵 🚑 🚒 🦽", " "),
}

func Poll() *onyxcord.Command {
	return &onyxcord.Command{
		Description:    "Organiser un vote",
		Usage:          "poll <template>,<question>,[choix...]",
		Category:       "management",
		Show:           true,
		ListenInPublic: true,
		Execute: func(arguments []string, bot onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			err = bot.Client.ChannelMessageDelete(message.ChannelID, message.ID)
			if err != nil {
				return
			}

			if len(arguments) < 3 {
				return errors.New("Il n'y a pas assez d'arguments")
			}
			if len(arguments[2:]) > 20 || len(arguments[2:]) == 0 {
				return errors.New("Le nombre de réponses doit être compris entre 1 et 22")
			}

			var template []string
			if arguments[0] == "" {
				template = pollChoices["letters"]
			} else {
				var ok bool
				template, ok = pollChoices[arguments[0]]
				if !ok {
					return errors.New(
						"Le modèle de choix est invalide. " +
							"Les choix possibles sont : `shapes`, `numbers`, `letters`, `food`, `faces`, `transportation`",
					)
				}
			}

			var choices string
			for index, value := range arguments[2:] {
				choices += fmt.Sprintf("%s %s\n", template[index], value)
			}

			userAvatar := message.Author.AvatarURL("64")
			poll := discordgo.MessageEmbed{
				Title:       fmt.Sprintf("**📊 Sondage :** %s", arguments[1]),
				Description: choices,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    message.Author.Username,
					IconURL: userAvatar,
				},
			}

			sentPoll, err := bot.Client.ChannelMessageSendEmbed(
				message.ChannelID,
				onyxcord.MakeEmbed(
					bot.Config,
					&poll,
				),
			)
			if err != nil {
				panic(err)
				return
			}

			for index := range arguments[2:] {
				err = bot.Client.MessageReactionAdd(message.ChannelID, sentPoll.ID, template[index])
				if err != nil {
					return
				}
			}

			return
		},
	}
}
