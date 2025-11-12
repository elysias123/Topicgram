package i18n

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*DefaultTranslator) CommandDescription_Ban() string {
	return "Ban user"
}

func (*DefaultTranslator) CommandDescription_Unban() string {
	return "Unban user"
}

func (*DefaultTranslator) CommandDescription_Terminate() string {
	return "Terminate current session"
}

func (*DefaultTranslator) CommandDescription_Verify() string {
	return "Set verification mode"
}

func (*DefaultTranslator) CommandUsage_Ban() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Usage: ")
	builder.Code("/ban <user_id>")

	return builder.Build()
}

func (*DefaultTranslator) CommandUsage_Unban() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Usage: ")
	builder.Code("/unban <user_id>")

	return builder.Build()
}

func (*DefaultTranslator) CommandUsage_Terminate() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Usage: ")
	builder.Code("/terminate <user_id>")

	return builder.Build()
}

func (*DefaultTranslator) CommandUsage_Verify() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Usage: ")
	builder.Code("/verify <calculate|off>")

	return builder.Build()
}
