package i18n

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*DefaultTranslator) Welcome(isPro bool, ad string) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("Hello!")
	builder.EOL()
	builder.Textln("You can contact us using this bot")

	if !isPro {
		builder.EOL()
		builder.Italicln("This bot was made by @TopicgramMessager")

		if ad != "" {
			builder.Textln("────────────")
			builder.Italicln("Advertise:")
			builder.Italic(ad)
		}
	}

	return builder.Build()
}

func (*DefaultTranslator) Banned() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("You have been permanently banned")

	return builder.Build()
}

func (*DefaultTranslator) Terminated() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("The administrator has terminated the session")
	builder.Text("You can send a new message to start a new session")

	return builder.Build()
}

func (*DefaultTranslator) Sender(user *botapi.User) (text string, entities []botapi.MessageEntity) {
	name := user.FirstName
	if user.LastName != "" {
		name += " " + user.LastName
	}

	builder := formatter.NewBuilder()

	builder.Textln("Sender")
	builder.Text("User ID ").Codeln(user.ID)
	builder.Text("Name ").Codeln(name)
	if user.UserName != "" {
		builder.Textfln("Username @%s", user.UserName)
	}

	return builder.Build()
}

func (*DefaultTranslator) Success() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Operation successful")

	return builder.Build()
}

func (*DefaultTranslator) Blocked(user_id int64) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("User ").Code(user_id).Text(" has blocked this bot")

	return builder.Build()
}

func (*DefaultTranslator) BanUser(user_id int64) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("User ").Code(user_id).Text(" has been permanently banned")

	return builder.Build()
}

func (*DefaultTranslator) UnbanUser(user_id int64) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("User ").Code(user_id).Text(" has been unbanned")

	return builder.Build()
}

func (*DefaultTranslator) VerifyModeUpdated(arg string) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Verification mode updated to ").Code(arg)

	return builder.Build()
}

func (*DefaultTranslator) Captcha_FirstQuestion(a, b int) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Please complete the verification before you can send messages:\n")
	builder.Text("Calculate: ").Code(a).Text(" + ").Code(b).Text(" = ?")

	return builder.Build()
}

func (*DefaultTranslator) Captcha_RepeatQuestion(a, b int) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Please answer the calculation result:\n")
	builder.Code(a).Text(" + ").Code(b).Text(" = ?")

	return builder.Build()
}

func (*DefaultTranslator) Captcha_FormatError(a, b int) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Invalid answer. Please reply with a number and try again.\n")
	builder.Text("Calculate: ").Text(a).Text(" + ").Text(b).Text(" = ?")

	return builder.Build()
}

func (*DefaultTranslator) Captcha_WrongAnswer(a, b int) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Wrong answer. Please try again:\n")
	builder.Text("Calculate: ").Code(a).Text(" + ").Code(b).Text(" = ?")

	return builder.Build()
}

func (*DefaultTranslator) Captcha_Passed() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Verification passed! You can now send messages.")

	return builder.Build()
}
