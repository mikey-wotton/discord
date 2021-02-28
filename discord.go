package discord

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type ChannelID string

type GuildID string

const (
	MikePublicChannelID         ChannelID = "806546506852597793"
	MikeShotsAndLeagueChannelID ChannelID = "806904335681978400"
	MikePublicGuildID           GuildID   = "806546506852597790"
	TyberiumPublicChannelID     ChannelID = "271377209124454401"
	TyberiumShotBotChannelID    ChannelID = "810889364082786304"
)

type Discord struct {
	session *discordgo.Session
	Config  Config
}

type Config struct {
	BotToken       string
	ChannelID      ChannelID
	GuildID        GuildID
	TimeBetweenTTS time.Duration
}

func New(c Config) (*Discord, error) {
	session, err := discordgo.New("Bot " + c.BotToken)
	if err != nil {
		return nil, err
	}

	session.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentsDirectMessages | discordgo.IntentsDirectMessageReactions

	d := &Discord{
		session: session,
		Config:  c,
	}

	if err := session.Open(); err != nil {
		return nil, err
	}

	return d, nil
}

func (d Discord) ListenFor(handler interface{}) {
	d.session.AddHandler(handler)
}

func (d Discord) Close() error {
	return d.session.Close()
}

func (d Discord) SendMessage(message string) error {
	_, err := d.session.ChannelMessageSend(string(d.Config.ChannelID), message)
	return err
}

func (d Discord) SendTTSMessage(message string) error {
	_, err := d.session.ChannelMessageSendTTS(string(d.Config.ChannelID), message)
	if err != nil {
		log.Fatal(err)
	}

	<-time.Tick(d.Config.TimeBetweenTTS)

	return nil
}

func tagGet(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.Contains(m.Content, "-poro") {
		return
	}

	args := strings.Split(m.Content, " ")

	fmt.Println(len(args))

	switch len(args) {
	case 3:
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "Shotbot Round: robot wars",
			Description: "Here are the round results!",
			Color:       200,
			Footer: &discordgo.MessageEmbedFooter{
				Text:         "brought to you by the bummers",
				IconURL:      "",
				ProxyIconURL: "",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Kirky",
					Value:  "2",
					Inline: true,
				},
				{
					Name:   "Dayle",
					Value:  "1",
					Inline: true,
				},
				{
					Name:   "Mikey",
					Value:  "1 ",
					Inline: true,
				},
				{
					Name:   "Gemima",
					Value:  "1",
					Inline: true,
				},
				{
					Name:  "\u200B",
					Value: "\u200B",
				},
				{
					Name:   "Field 1 Name",
					Value:  "Field 1 Value",
					Inline: true,
				},
				{
					Name:   "Field 3 Name",
					Value:  "Field 3 Value",
					Inline: true,
				},
				{
					Name:   "Field 1 Name",
					Value:  "Field 1 Value",
					Inline: true,
				},
				{
					Name:   "Field 1 Name",
					Value:  "Field 1 Value",
					Inline: true,
				},
				{
					Name:   "Field 3 Name",
					Value:  "Field 3 Value",
					Inline: true,
				},
				{
					Name:   "Field 1 Name",
					Value:  "Field 1 Value",
					Inline: true,
				},
				{
					Name:   "Field 1 Name",
					Value:  "Field 1 Value",
					Inline: true,
				},
				{
					Name:   "Field 3 Name",
					Value:  "Field 3 Value",
					Inline: true,
				},
				{
					Name:   "Field 1 Name",
					Value:  "Field 1 Value",
					Inline: true,
				},
				{
					Name:   "Field 1 Name",
					Value:  "Field 1 Value",
					Inline: true,
				},
				{
					Name:   "Field 3 Name",
					Value:  "Field 3 Value",
					Inline: true,
				},
				{
					Name:   "Field 1 Name",
					Value:  "Field 1 Value",
					Inline: true,
				},
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	case 4:

	default:
		if _, err := s.ChannelMessageSend(m.ChannelID, "Pong!"); err != nil {
			log.Fatal(err)
		}
	}

}
