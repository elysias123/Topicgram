package i18n

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*DefaultTranslator) BotFather_Error_InvalidToken() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Invalid Bot Token")

	return builder.Build()
}

func (*DefaultTranslator) BotFather_Error_BotExists() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Bot exists")

	return builder.Build()
}

func (*DefaultTranslator) BotFather_Error_BotNotExists() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Bot not exists")

	return builder.Build()
}

func (*DefaultTranslator) Error_PermissionDenied() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Permission denied")

	return builder.Build()
}

func (*DefaultTranslator) Error_AdminRequired() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("Please promote this bot to an administrator and grant the following permissions")
	builder.EOL()
	builder.Textln("Permission required:")
	builder.Textln("- Delete Messages")
	builder.Textln("- Pin Messages")
	builder.Textln("- Manage topics")

	return builder.Build()
}

func (*DefaultTranslator) Error_NoForwardGroup() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("This bot cannot currently forward messages because it does not have a forwarding group set up")

	return builder.Build()
}
