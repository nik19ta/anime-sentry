package animegoorg

import (
	"anime-bot-schedule/pkg/message"
	repositories_subscribe "anime-bot-schedule/repositories/subscribe"
	parsing "anime-bot-schedule/services/parser/animego.org"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var LINK_PATTERN = `^https://animego.org/anime/.*$`

func Handle(update tgbotapi.Update) message.NewMessage {
	data, err := parsing.Fetch(update.Message.Text)

	if err != nil {
		msg := message.NewMessage{
			Text: "Произошла ошибка :(",
		}

		return msg
	}

	if len(*data.Title) == 0 {
		msg := message.NewMessage{
			Text:  "Мы не нашли такого аниме",
			Photo: "https://animego.org/animego/images/404.gif",
		}

		return msg
	}

	var lastEpisod parsing.Episod

	if data.Episods[0].Relized {
		lastEpisod = data.Episods[0]
	} else if data.Episods[1].Relized {
		lastEpisod = data.Episods[1]
	} else if data.Episods[2].Relized {
		lastEpisod = data.Episods[2]
	}

	err = repositories_subscribe.SubscribeToAnime(update.Message.Chat.ID, update.Message.Text,
		*data.Title, *data.Image, lastEpisod.Number)

	if err != nil {
		if err.Error() == "you are already subscribed to this anime" {
			msg := message.NewMessage{
				Text: "Вы уже подписанны на это аниме!",
			}

			return msg
		}

		msg := message.NewMessage{
			Text: "Произошла неизвестная ошибка :(",
		}
		return msg
	}

	messageText := fmt.Sprintf("%s\n\nАниме сохраненно, вы будете получать уведомления когда выйдут новые серии. \n\n%s (%s) выйдет %s.",
		*data.Title, lastEpisod.Number, lastEpisod.Title, lastEpisod.Date)

	newMsg := message.NewMessage{
		Text: messageText,
	}

	if data.Image != nil && *data.Image != "" {
		newMsg.Photo = *data.Image
	}

	return newMsg
}
