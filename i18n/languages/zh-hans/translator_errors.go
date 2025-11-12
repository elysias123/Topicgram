package translator

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*Translator) Error() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("发生错误, 请重试")

	return builder.Build()
}

func (*Translator) Error_Database() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("连接数据库失败, 请重试")

	return builder.Build()
}

func (*Translator) Error_TopicRequired() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("请在此群组启用话题")

	return builder.Build()
}

func (*Translator) Error_UnknwonCommand() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("未知命令")

	return builder.Build()
}
func (*Translator) Error_UnsupportedMessage() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("不支持的消息")

	return builder.Build()
}

func (*Translator) Error_ForwardForbidden() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("此 Bot 当前无法转发消息, 因为转发受限")

	return builder.Build()
}

func (*Translator) Error_FailedToCreateTopic(forUser bool) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	if forUser {
		builder.Text("创建话题失败, 请重试")
	} else {
		builder.Text("创建话题失败")
	}

	return builder.Build()
}

func (*Translator) Error_FailedToEdit() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("此消息无法编辑")

	return builder.Build()
}