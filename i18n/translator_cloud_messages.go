package i18n

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*DefaultTranslator) BotFather_Welcome() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("Welcome to use Topicgram!")
	builder.Textln("Our notification channel @TopicgramMessager")
	builder.EOL()
	builder.Text("Now, you can use /addbot to add a new bot")

	return builder.Build()
}

func (*DefaultTranslator) BotFather_BotAdded() (text string, entities []botapi.MessageEntity) {
    builder := formatter.NewBuilder()

    builder.Textln("Bot successfully added")
    builder.EOL()
    builder.Textln("Add this bot to your group as an administrator, then use /setgroup to set forward group")
    builder.Italic("Warning! New group owner are anonymous by default. You must use current account or Permission Denied")

    return builder.Build()
}

func (*DefaultTranslator) BotFather_BotRemoved() (text string, entities []botapi.MessageEntity) {
    builder := formatter.NewBuilder()

    builder.Text("Bot successfully removed")

    return builder.Build()
}

func (*DefaultTranslator) BotFather_OwnBots(usernames []string, ids []int64) (text string, entities []botapi.MessageEntity) {
    builder := formatter.NewBuilder()

    builder.Textln("You own these bots:")
    for i, username := range usernames {
        builder.Textf("- @%s (", username).Code(ids[i]).Textln(")")
    }

    return builder.Build()
}