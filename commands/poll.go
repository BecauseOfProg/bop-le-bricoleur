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
		ListenInPublic: true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			answers := strings.Split(interaction.Data.Options[2].StringValue(), ",")
			if len(answers) > 20 || len(answers) == 0 {
				return errors.New("Le nombre de réponses doit être compris entre 1 et 22")
			}

			template := pollChoices[interaction.Data.Options[1].StringValue()]

			var choices string
			for index, value := range answers {
				choices += fmt.Sprintf("%s %s\n", template[index], value)
			}

			userAvatar := interaction.Member.User.AvatarURL("64")
			poll := discordgo.MessageEmbed{
				Title:       fmt.Sprintf("**📊 Sondage :** %s", interaction.Data.Options[0].StringValue()),
				Description: choices,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    interaction.Member.User.Username,
					IconURL: userAvatar,
				},
			}

			sentPoll, err := bot.Client.ChannelMessageSendEmbed(
				interaction.ChannelID,
				onyxcord.MakeEmbed(
					bot.Config,
					&poll,
				),
			)
			if err != nil {
				panic(err)
				return
			}

			for index := range answers {
				err = bot.Client.MessageReactionAdd(interaction.ChannelID, sentPoll.ID, template[index])
				if err != nil {
					return
				}
			}

			return
		},
	}
}
