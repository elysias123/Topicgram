package translator

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*Translator) CommandDescription_Ban() string {
	return "封禁 user"
}

func (*Translator) CommandDescription_Unban() string {
	return "解封 user"
}

func (*Translator) CommandDescription_Terminate() string {
	return "结束当前会话"
}

func (*Translator) CommandDescription_Verify() string {
	return "设置验证模式"
}

func (*Translator) CommandUsage_Ban() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用法: ")
	builder.Code("/ban <用户ID>")

	return builder.Build()
}

func (*Translator) CommandUsage_Unban() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用法: ")
	builder.Code("/unban <用户ID>")

	return builder.Build()
}

func (*Translator) CommandUsage_Terminate() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用法: ")
	builder.Code("/terminate <用户ID>")

	return builder.Build()
}

func (*Translator) CommandUsage_Verify() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用法: ")
	builder.Code("/verify <calculate|off>")

	return builder.Build()
}
