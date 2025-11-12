package translator

import "Topicgram/i18n"

var _ i18n.Translator = &Translator{}

type Translator struct {
	i18n.DefaultTranslator
}

func init() {
	i18n.RegisterTranslator("zh-hans", &Translator{})
}
