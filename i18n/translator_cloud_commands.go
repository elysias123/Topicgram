package i18n

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	"gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*DefaultTranslator) BotFather_CommandDescription_AddBot() string {
	return "Add a new bot"
}

func (*DefaultTranslator) BotFather_CommandDescription_DelBot() string {
	return "Remove a bot"
}

func (*DefaultTranslator) BotFather_CommandDescription_MyBots() string {
	return "List my bots"
}

func (*DefaultTranslator) CommandDescription_SetGroup() string {
	return "Set this group as forward group"
}

func (*DefaultTranslator) BotFather_CommandUsage_AddBot() (string, []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Usage: ")
	builder.Code("/addbot <bot_token>")

	return builder.Build()
}

func (*DefaultTranslator) BotFather_CommandUsage_DelBot() (string, []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Usage: ")
	builder.Code("/delbot <bot_id>")

	return builder.Build()
}

func (*DefaultTranslator) BotFather_CommandUsage_Notice() (string, []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Usage: ")
	builder.Code("/notice <user_id> <text>")

	return builder.Build()
}

func (*DefaultTranslator) BotFather_CommandUsage_NoticeAll() (string, []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Usage: ")
	builder.Code("/notice_all <text>")

	return builder.Build()
}
