package i18n

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestMock(t *testing.T) {
	// Must not panic
	mock := newMock()
	mock.SetDefault(discordgo.ChineseCN)
	assert.NoError(t, mock.LoadBundle(discordgo.SpanishES, ""))
	assert.Empty(t, mock.Get(discordgo.Croatian, "", nil))
	assert.Empty(t, mock.GetDefault("", nil))
	assert.Nil(t, mock.GetLocalizations("", nil))

	var called bool
	mock.SetDefaultFunc = func(_ discordgo.Locale) {
		called = true
	}
	mock.LoadBundleFunc = func(_ discordgo.Locale, _ string) error {
		called = true
		return nil
	}
	mock.GetFunc = func(_ discordgo.Locale, _ string, _ Vars) string {
		called = true
		return ""
	}
	mock.GetDefaultFunc = func(_ string, _ Vars) string {
		called = true
		return ""
	}
	mock.GetLocalizationsFunc = func(_ string, _ Vars) *map[discordgo.Locale]string {
		called = true
		return nil
	}

	called = false
	mock.SetDefault(discordgo.ChineseCN)
	assert.True(t, called)

	called = false
	assert.NoError(t, mock.LoadBundle(discordgo.SpanishES, ""))
	assert.True(t, called)

	called = false
	assert.Empty(t, mock.Get(discordgo.Croatian, "", nil))
	assert.True(t, called)

	called = false
	assert.Empty(t, mock.GetDefault("", nil))
	assert.True(t, called)

	called = false
	assert.Empty(t, mock.GetLocalizations("", nil))
	assert.True(t, called)
}
