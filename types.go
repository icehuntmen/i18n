package i18n

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// Vars is the collection used to inject variables during translation.
// This type only exists to be less verbose.
type Vars map[string]interface{}

type translator interface {
	SetDefault(locale discordgo.Locale)
	LoadBundle(locale discordgo.Locale, file string) error
	Get(locale discordgo.Locale, key string, values Vars) string
	GetDefault(key string, values Vars) string
	GetLocalizations(key string, variables Vars) *map[discordgo.Locale]string
}

type translatorImpl struct {
	defaultLocale discordgo.Locale
	translations  map[discordgo.Locale]bundle
	loadedBundles map[string]bundle
	logger        *logrus.Logger
}

type translatorMock struct {
	SetDefaultFunc       func(locale discordgo.Locale)
	LoadBundleFunc       func(locale discordgo.Locale, file string) error
	GetFunc              func(locale discordgo.Locale, key string, values Vars) string
	GetDefaultFunc       func(key string, values Vars) string
	GetLocalizationsFunc func(key string, variables Vars) *map[discordgo.Locale]string
}

type bundle map[string][]string
