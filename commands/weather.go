package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/exybore/goweather"
	"github.com/theovidal/onyxcord"
)

func Weather() *onyxcord.Command {
	return &onyxcord.Command{
		Description:    "Obtenir la météo à un endroit précis",
		Usage:          "weather <localisation>",
		Show:           true,
		ListenInPublic: true,
		ListenInDM:     true,
		Execute: func(arguments []string, bot onyxcord.Bot, message *discordgo.MessageCreate) (err error) {
			weatherAPI, err := goweather.NewAPI(os.Getenv("OPENWEATHERMAP_KEY"), "fr", "metric")
			if err != nil {
				return
			}

			location := strings.Join(arguments, ",")
			weather, err := weatherAPI.Current(location)
			if err != nil {
				return errors.New(":satellite_orbital: Cette localisation est inconnue")
			}

			_, err = bot.Client.ChannelMessageSendEmbed(
				message.ChannelID,
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
			)
			return
		},
	}
}
