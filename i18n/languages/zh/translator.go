package translator

import (
	"Topicgram/i18n"
	translator "Topicgram/i18n/languages/zh-hans"
)

func init() {
	i18n.RegisterTranslator("zh", &translator.Translator{})
}
