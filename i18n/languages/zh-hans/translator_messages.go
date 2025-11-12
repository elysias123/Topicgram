package translator

import (
	botapi "github.com/OvyFlash/telegram-bot-api"
	formatter "gitlab.com/CoiaPrant/telegram-bot-formatter"
)

func (*Translator) Welcome(isPro bool, ad string) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()
	builder.Textln("你好!")
	builder.EOL()
	builder.Textln("你可以使用此 Bot 与我们联系")

	if !isPro {
		builder.EOL()
		builder.Italicln("此 Bot 由 @TopicgramMessager 创建")

		if ad != "" {
			builder.Textln("────────────")
			builder.Italicln("广告:")
			builder.Italic(ad)
		}
	}

	return builder.Build()
}

func (*Translator) Banned() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("你已被永久封禁")

	return builder.Build()
}

func (*Translator) Terminated() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Textln("管理员已终止会话")
	builder.Text("你可以发送新消息来创建新会话")

	return builder.Build()
}

func (*Translator) Sender(user *botapi.User) (text string, entities []botapi.MessageEntity) {
	name := user.FirstName
	if user.LastName != "" {
		name += " " + user.LastName
	}

	builder := formatter.NewBuilder()

	builder.Textln("发送者")
	builder.Text("用户ID ").Codeln(user.ID)
	builder.Text("昵称 ").Codeln(name)
	if user.UserName != "" {
		builder.Textfln("用户名 @%s", user.UserName)
	}

	return builder.Build()
}

func (*Translator) Success() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("操作成功")

	return builder.Build()
}

func (*Translator) Blocked(user_id int64) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用户 ").Code(user_id).Text(" 已屏蔽此 Bot")

	return builder.Build()
}

func (*Translator) BanUser(user_id int64) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用户 ").Code(user_id).Text(" 已被永久封禁")

	return builder.Build()
}

func (*Translator) UnbanUser(user_id int64) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("用户 ").Code(user_id).Text(" 已被解封")

	return builder.Build()
}

func (*Translator) VerifyModeUpdated(arg string) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("验证模式已修改为 ").Code(arg)

	return builder.Build()
}

func (*Translator) Captcha_FirstQuestion(a, b int) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("请先完成验证才能发送消息：\n")
	builder.Text("计算: ").Code(a).Text(" + ").Code(b).Text(" = ?")

	return builder.Build()
}

func (*Translator) Captcha_RepeatQuestion(a, b int) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("请回答计算结果：\n")
	builder.Code(a).Text(" + ").Code(b).Text(" = ?")

	return builder.Build()
}

func (*Translator) Captcha_FormatError(a, b int) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("答案错误 只需要回答数字 请再试一次\n")
	builder.Text("计算: ").Text(a).Text(" + ").Text(b).Text(" = ?")

	return builder.Build()
}

func (*Translator) Captcha_WrongAnswer(a, b int) (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("答案错误 请再试一次：\n")
	builder.Text("计算: ").Code(a).Text(" + ").Code(b).Text(" = ?")

	return builder.Build()
}

func (*Translator) Captcha_Passed() (text string, entities []botapi.MessageEntity) {
	builder := formatter.NewBuilder()

	builder.Text("验证通过！您现在可以发送消息。")

	return builder.Build()
}
