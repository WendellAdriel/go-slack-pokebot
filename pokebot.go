package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/WendellAdriel/go-slack-pokebot/commands"
	"github.com/WendellAdriel/go-slack-pokebot/config"
	"github.com/nlopes/slack"
)

func main() {
	chSender := make(chan slack.OutgoingMessage)
	chReceiver := make(chan slack.SlackEvent)
	var botId string

	api := slack.New(config.GetToken())
	api.SetDebug(false)

	wsAPI, err := api.StartRTM("", "http://example.com")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	go wsAPI.HandleIncomingEvents(chReceiver)
	go wsAPI.Keepalive(20 * time.Second)

	go func(wsAPI *slack.SlackWS, chSender chan slack.OutgoingMessage) {
		for {
			select {
			case msg := <-chSender:
				wsAPI.SendMessage(&msg)
			}
		}
	}(wsAPI, chSender)

	for {
		select {
		case msg := <-chReceiver:
			switch ev := msg.Data.(type) {
			case *slack.PresenceChangeEvent:
				botId = ev.UserId
			case *slack.MessageEvent:
				if isBotMessage(ev, botId) {
					buildReply(ev, api)
				}
			case *slack.SlackWSError:
				error := msg.Data.(*slack.SlackWSError)
				fmt.Printf("Error: %d - %s\n", error.Code, error.Msg)
			default:
				// Ignore
			}
		}
	}
}

func isBotMessage(ev *slack.MessageEvent, botId string) bool {
	if ev.Type == "message" && strings.HasPrefix(ev.Text, "<@"+botId+">") {
		return true
	}
	return false
}

func buildReply(ev *slack.MessageEvent, api *slack.Slack) {
	messageArray := strings.Fields(ev.Text)
	response := buildResponseText(messageArray)
	sendMessage(ev, api, response)
}

func buildResponseText(messageArray []string) string {
	command := messageArray[1]
	var response string

	switch command {
	case "help":
		response = commands.ExecHelpCommand()
	case "pokemon":
		commandParameter := messageArray[2]
		response = commands.ExecPokemonCommand(commandParameter)
	default:
		response = commands.ExecDefaultCommand()
	}

	return response
}

func sendMessage(ev *slack.MessageEvent, api *slack.Slack, response string) {
	_, _, chanID, err := api.OpenIMChannel(ev.UserId)
	if err != nil {
		log.Fatal("Error opening IM: ", err)
	}

	params := slack.PostMessageParameters{}
	channelId, timestamp, err := api.PostMessage(chanID, response, params)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelId, timestamp)
}
