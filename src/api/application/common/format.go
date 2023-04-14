package common

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	ErrorLog = "ERROR:: [DOMAIN]%s [LAYER]%s [ERROR]%s"
)

// NormalizeString remove white spaces and set case to name
func NormalizeString(name *string) {
	*name = cases.Title(language.English, cases.Compact).String(strings.TrimSpace(strings.ToLower(*name)))
}
