package translator

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*Translator) BotFather_Error_InvalidToken() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("无效的 Bot Token")

	return builder.Build()
}

func (*Translator) BotFather_Error_BotExists() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Bot 已存在")

	return builder.Build()
}

func (*Translator) BotFather_Error_BotNotExists() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("Bot 不存在")

	return builder.Build()
}

func (*Translator) Error_PermissionDenied() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("权限不足")

	return builder.Build()
}

func (*Translator) Error_AdminRequired() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("请将此机器人提升为管理员并授予以下权限")
	builder.EOL()
	builder.Textln("所需权限:")
	builder.Textln("- 删除消息")
	builder.Textln("- 置顶消息")
	builder.Textln("- 管理话题")

	return builder.Build()
}

func (*Translator) Error_NoForwardGroup() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("此机器人当前无法转发消息, 因为尚未设置转发群组")

	return builder.Build()
}

