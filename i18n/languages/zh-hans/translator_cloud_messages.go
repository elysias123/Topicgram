package translator

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*Translator) BotFather_Welcome() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("欢迎使用 Topicgram!")
	builder.Textln("我们的通知频道 @TopicgramMessager")
	builder.EOL()
	builder.Text("现在，您可以使用 /addbot 添加新的机器人")

	return builder.Build()
}

func (*Translator) BotFather_BotAdded() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("成功添加 Bot")
	builder.EOL()
	builder.Textln("将此 Bot 以管理员身份添加到您的群组, 然后使用 /setgroup 设置转发群组")
	builder.Italic("注意! 新群组的创建者默认匿名. 您必须使用当前账户, 否则提示权限不足")

	return builder.Build()
}

func (*Translator) BotFather_BotRemoved() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("成功删除 Bot")

	return builder.Build()
}

func (*Translator) BotFather_OwnBots(usernames []string, ids []int64) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("您拥有以下机器人:")
	for i, username := range usernames {
		builder.Textf("- @%s (", username).Code(ids[i]).Textln(")")
	}

	return builder.Build()
}
