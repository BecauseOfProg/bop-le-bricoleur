module github.com/BecauseOfProg/bop-le-bricoleur

go 1.15

require (
	github.com/bwmarrin/discordgo v0.23.0
	github.com/exybore/goweather v0.1.1
	github.com/theovidal/onyxcord v0.1.0
	go.mongodb.org/mongo-driver v1.4.4
)

replace github.com/theovidal/onyxcord => ../../theovidal/onyxcord

replace github.com/bwmarrin/discordgo => github.com/FedorLap2006/discordgo v0.22.1-0.20210217184539-8718e2d37898
