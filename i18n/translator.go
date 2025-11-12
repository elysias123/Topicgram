package i18n

var defaultTranslator Translator = &DefaultTranslator{}

type DefaultTranslator struct{}

func (*DefaultTranslator) embeddedCheck() {}
