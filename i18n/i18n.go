package i18n

import (
	"strings"
	"sync"

	botapi "github.com/OvyFlash/telegram-bot-api"
)

type (
	LanguageCode = string

	Translator interface {
		embeddedCheck()

		/* For cloud-hosted Topicgram */
		BotFather_CommandDescription_AddBot() string
		BotFather_CommandDescription_DelBot() string
		BotFather_CommandDescription_MyBots() string
		CommandDescription_SetGroup() string

		BotFather_CommandUsage_AddBot() (text string, entities []botapi.MessageEntity)
		BotFather_CommandUsage_DelBot() (text string, entities []botapi.MessageEntity)
		BotFather_CommandUsage_Notice() (text string, entities []botapi.MessageEntity)
		BotFather_CommandUsage_NoticeAll() (text string, entities []botapi.MessageEntity)

		BotFather_Welcome() (text string, entities []botapi.MessageEntity)
		BotFather_BotAdded() (text string, entities []botapi.MessageEntity)
		BotFather_BotRemoved() (text string, entities []botapi.MessageEntity)
		BotFather_OwnBots(usernames []string, ids []int64) (text string, entities []botapi.MessageEntity)

		BotFather_Error_InvalidToken() (text string, entities []botapi.MessageEntity)
		BotFather_Error_BotExists() (text string, entities []botapi.MessageEntity)
		BotFather_Error_BotNotExists() (text string, entities []botapi.MessageEntity)
		Error_PermissionDenied() (text string, entities []botapi.MessageEntity)
		Error_AdminRequired() (text string, entities []botapi.MessageEntity)
		Error_NoForwardGroup() (text string, entities []botapi.MessageEntity)

		/* For all Topicgram */
		CommandDescription_Ban() string
		CommandDescription_Unban() string
		CommandDescription_Terminate() string
		CommandDescription_Verify() string

		CommandUsage_Ban() (text string, entities []botapi.MessageEntity)
		CommandUsage_Unban() (text string, entities []botapi.MessageEntity)
		CommandUsage_Terminate() (text string, entities []botapi.MessageEntity)
		CommandUsage_Verify() (text string, entities []botapi.MessageEntity)

		Error() (text string, entities []botapi.MessageEntity)
		Error_Database() (text string, entities []botapi.MessageEntity)
		Error_TopicRequired() (text string, entities []botapi.MessageEntity)
		Error_UnknwonCommand() (text string, entities []botapi.MessageEntity)
		Error_UnsupportedMessage() (text string, entities []botapi.MessageEntity)
		Error_ForwardForbidden() (text string, entities []botapi.MessageEntity)
		Error_FailedToCreateTopic(forUser bool) (text string, entities []botapi.MessageEntity)
		Error_FailedToEdit() (text string, entities []botapi.MessageEntity)

		Welcome(isPro bool, ad string) (text string, entities []botapi.MessageEntity)
		Banned() (text string, entities []botapi.MessageEntity)
		Terminated() (text string, entities []botapi.MessageEntity)

		Sender(user *botapi.User) (text string, entities []botapi.MessageEntity)
		Success() (text string, entities []botapi.MessageEntity)
		Blocked(user_id int64) (text string, entities []botapi.MessageEntity)
		BanUser(user_id int64) (text string, entities []botapi.MessageEntity)
		UnbanUser(user_id int64) (text string, entities []botapi.MessageEntity)
		VerifyModeUpdated(arg string) (text string, entities []botapi.MessageEntity)

		/* Captcha / Verification */
		Captcha_FirstQuestion(a, b int) (text string, entities []botapi.MessageEntity)
		Captcha_RepeatQuestion(a, b int) (text string, entities []botapi.MessageEntity)
		Captcha_FormatError(a, b int) (text string, entities []botapi.MessageEntity)
		Captcha_WrongAnswer(a, b int) (text string, entities []botapi.MessageEntity)
		Captcha_Passed() (text string, entities []botapi.MessageEntity)
	}
)

var (
	mutex       sync.RWMutex
	translators = make(map[LanguageCode]Translator)
)

func RegisterTranslator(code LanguageCode, translator Translator) {
	mutex.Lock()
	defer mutex.Unlock()

	translators[code] = translator
}

func GetOrDefault(code LanguageCode) Translator {
	mutex.RLock()
	defer mutex.RUnlock()

	if translator, ok := translators[code]; ok {
		return translator
	}

	if strings.Contains(code, "-") {
		code, _, _ = strings.Cut(code, "-")

		if translator, ok := translators[code]; ok {
			return translator
		}
	}

	return defaultTranslator
}

func Range(f func(code LanguageCode, translator Translator)) {
	f("", defaultTranslator)

	mutex.RLock()
	defer mutex.RUnlock()

	for code, translator := range translators {
		f(code, translator)
	}
}
