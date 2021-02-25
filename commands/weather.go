package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/exybore/goweather"
	"github.com/theovidal/onyxcord"
)

func Weather() *onyxcord.Command {
	return &onyxcord.Command{
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(bot *onyxcord.Bot, interaction *discordgo.InteractionCreate) (err error) {
			weatherAPI, err := goweather.NewAPI(os.Getenv("OPENWEATHERMAP_KEY"), "fr", "metric")
			if err != nil {
				log.Panicln("‼ Error creating the weather API:", err)
			}

			location := interaction.Data.Options[0].StringValue()
			fmt.Println(os.Getenv("OPENWEATHERMAP_KEY"))
			weather, err := weatherAPI.Current(location)
			if err != nil {
				fmt.Println(err)
				return errors.New(":satellite_orbital: Cette localisation est inconnue")
			}

			_ = bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Embeds: []*discordgo.MessageEmbed{
						onyxcord.MakeEmbed(
							bot.Config,
							&discordgo.MessageEmbed{
								Title: fmt.Sprintf("%s :flag_%s:", weather.City.Name, strings.ToLower(weather.City.Country)),
								Description: fmt.Sprintf("**%s**\n\n"+
									":thermometer: Température : %.1f°C\n"+
									":droplet: Humidité : %.1f%%\n"+
									":cloud: Nuages : %.1f%%\n"+
									":dash: Vent : %.1f km/h",
									strings.Title(weather.Conditions.Description), weather.Conditions.Temperature,
									weather.Conditions.Humidity, weather.Conditions.Clouds,
									weather.Conditions.WindSpeed*3.6),
								Thumbnail: &discordgo.MessageEmbedThumbnail{
									URL: weather.Conditions.IconURL,
								},
							},
						),
					},
				},
			})
			return
		},
	}
}
