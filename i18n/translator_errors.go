package i18n

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*DefaultTranslator) Error() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("An error has occurred, please try again")

	return builder.Build()
}

func (*DefaultTranslator) Error_Database() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Failed to connect database, please try again")

	return builder.Build()
}

func (*DefaultTranslator) Error_TopicRequired() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Please enable topic mode in this group")

	return builder.Build()
}

func (*DefaultTranslator) Error_UnknwonCommand() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Unknown command")

	return builder.Build()
}

func (*DefaultTranslator) Error_UnsupportedMessage() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Unsupported message")

	return builder.Build()
}

func (*DefaultTranslator) Error_ForwardForbidden() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("This bot cannot currently forward messages because forwarding was restricted")

	return builder.Build()
}

func (*DefaultTranslator) Error_FailedToCreateTopic(forUser bool) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	if forUser {
		builder.Text("Failed to create topic, please try again")
	} else {
		builder.Text("Failed to create topic")
	}

	return builder.Build()
}

func (*DefaultTranslator) Error_FailedToEdit() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("This message can not be edited")

	return builder.Build()
}
