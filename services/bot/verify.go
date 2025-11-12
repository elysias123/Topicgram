package bot

import (
	. "Topicgram/database"
	"Topicgram/i18n"
	"Topicgram/model"
	"Topicgram/utils"
	"strconv"
	"strings"

	botapi "github.com/OvyFlash/telegram-bot-api"
)

type CaptchaInfo struct {
	a        int
	b        int
	answer   int
	attempts int
}

func newCalculateCaptcha() *CaptchaInfo {
	for {
		a := utils.GetRand(1, 60)
		b := utils.GetRand(1, 60)
		if a+b <= 100 {
			return &CaptchaInfo{a: a, b: b, answer: a + b}
		}
	}
}

func (bot *Bot) Userverify(msg *botapi.Message, translator i18n.Translator, currentChat botapi.BaseChat) bool {
	return bot.calculateCaptcha(msg, translator, currentChat)
}

func (bot *Bot) calculateCaptcha(msg *botapi.Message, translator i18n.Translator, currentChat botapi.BaseChat) bool {
	userID := msg.From.ID

	bot.captchaMu.Lock()
	if bot.captcha == nil {
		bot.captcha = make(map[int64]*CaptchaInfo)
	}
	ci := bot.captcha[userID]
	bot.captchaMu.Unlock()

	if ci == nil {
		ci = newCalculateCaptcha()
		bot.captchaMu.Lock()
		bot.captcha[userID] = ci
		bot.captchaMu.Unlock()
		text, entities := translator.Captcha_FirstQuestion(ci.a, ci.b)
		bot.Send(botapi.MessageConfig{BaseChat: currentChat, Text: text, Entities: entities})
		return false
	}

	if msg.Text == "" || strings.HasPrefix(msg.Text, "/") {
		text, entities := translator.Captcha_RepeatQuestion(ci.a, ci.b)
		bot.Send(botapi.MessageConfig{BaseChat: currentChat, Text: text, Entities: entities})
		return false
	}

	answer, err := strconv.ParseInt(strings.TrimSpace(msg.Text), 10, 64)
	if err != nil {
		ci.attempts++
		text, entities := translator.Captcha_FormatError(ci.a, ci.b)
		bot.Send(botapi.MessageConfig{BaseChat: currentChat, Text: text, Entities: entities})
		return false
	}

	if int(answer) != ci.answer {
		ci.attempts++
		if ci.attempts >= 3 {

			bot.captchaMu.Lock()
			delete(bot.captcha, userID)
			bot.captchaMu.Unlock()

			var t model.Topic
			if err := DB().Where("user_id", userID).Find(&t).Error; err == nil {
				if t.Id == 0 {
					t = model.Topic{UserId: userID, TopicId: 0, IsBan: true, LanguageCode: msg.From.LanguageCode}
					DB().Create(&t)
				} else if !t.IsBan {
					t.IsBan = true
					DB().Save(&t)
				}
			}

			text, entities := translator.Banned()
			bot.Send(botapi.MessageConfig{BaseChat: currentChat, Text: text, Entities: entities})
			return false
		}

		newCi := newCalculateCaptcha()
		bot.captchaMu.Lock()
		bot.captcha[userID] = newCi
		bot.captchaMu.Unlock()
		text, entities := translator.Captcha_WrongAnswer(newCi.a, newCi.b)
		bot.Send(botapi.MessageConfig{BaseChat: currentChat, Text: text, Entities: entities})
		return false
	}

	bot.captchaMu.Lock()
	delete(bot.captcha, userID)
	bot.captchaMu.Unlock()

	bot.skipMu.Lock()
	if bot.skipVerifiedForward == nil {
		bot.skipVerifiedForward = make(map[int64]bool)
	}

	bot.skipVerifiedForward[userID] = true
	bot.skipMu.Unlock()
	text, entities := translator.Captcha_Passed()
	bot.Send(botapi.MessageConfig{BaseChat: currentChat, Text: text, Entities: entities})
	return true
}
