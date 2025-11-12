package translator

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	"gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*Translator) BotFather_CommandDescription_AddBot() string {
	return "添加 Bot"
}

func (*Translator) BotFather_CommandDescription_DelBot() string {
	return "删除 Bot"
}

func (*Translator) BotFather_CommandDescription_MyBots() string {
	return "列出我的 Bots"
}

func (*Translator) CommandDescription_SetGroup() string {
	return "将此群组设置为转发群组"
}

func (*Translator) BotFather_CommandUsage_AddBot() (string, []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用法: ")
	builder.Code("/addbot <bot_token>")

	return builder.Build()
}

func (*Translator) BotFather_CommandUsage_DelBot() (string, []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用法: ")
	builder.Code("/delbot <bot_id>")

	return builder.Build()
}

func (*Translator) BotFather_CommandUsage_Notice() (string, []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用法: ")
	builder.Code("/notice <用户ID> <内容>")

	return builder.Build()
}

func (*Translator) BotFather_CommandUsage_NoticeAll() (string, []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用法: ")
	builder.Code("/notice_all <内容>")

	return builder.Build()
}
