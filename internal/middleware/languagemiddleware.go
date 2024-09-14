package middleware

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
	"strings"
)

type LanguageMiddleware struct {
	bundle    *i18n.Bundle
	Localizer *i18n.Localizer
}

func NewLanguageMiddleware(bundle *i18n.Bundle) *LanguageMiddleware {
	return &LanguageMiddleware{
		bundle:    bundle,
		Localizer: i18n.NewLocalizer(bundle, "zh-Hans"),
	}
}
func (m *LanguageMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acceptLang := r.Header.Get("Accept-Language")
		lang := parseAcceptLanguage(acceptLang)
		m.Localizer = i18n.NewLocalizer(m.bundle, lang)
		next(w, r)
	}
}
func parseAcceptLanguage(acceptLang string) string {
	languages := strings.Split(acceptLang, ",")
	if len(languages) > 0 {
		lang := strings.Split(languages[0], ";")[0]
		if lang == "zh" {
			return "zh"
		}
	}
	return "en"
}
