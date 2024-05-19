package i18n_test

import (
	"testing"

	"github.com/pmperrin/go-basic-i18n/i18n"
	"github.com/stretchr/testify/assert"
)

func TestLoadPropertiesContante(t *testing.T) {

	instance, e1 := i18n.InitI18n("../example/local", "messages", "en")
	assert.NoError(t, e1)

	// Test loading English properties
	_, err := instance.GetLang("en")
	assert.NoError(t, err)

	// Test loading French properties
	_, err = instance.GetLang("fr")
	assert.NoError(t, err)
}

func TestGetText(t *testing.T) {

	// Initialize the i18n system
	i18nInstance, e1 := i18n.InitI18n("../example/local", "messages", "en")
	assert.NoError(t, e1)

	// Retrieve the English language file
	englishLang, err := i18nInstance.GetLang("en")
	assert.NoError(t, err)

	// Test getting a text in English
	text := englishLang.GetText("welcome.message")
	assert.Equal(t, "Welcome to our application", text)

	// Test getting a text in French
	frenchLang, err := i18nInstance.GetLang("fr")
	assert.NoError(t, err)

	text = frenchLang.GetText("welcome.message")
	assert.Equal(t, "Bienvenue dans notre application", text)
}

func TestGetTextWithParam(t *testing.T) {

	// Initialize the i18n system
	i18nInstance, err := i18n.InitI18n("../example/local", "messages", "en")
	assert.NoError(t, err)

	// Retrieve the English language file
	englishLang, err := i18nInstance.GetLang("en")
	assert.NoError(t, err)

	// Define parameters for the text
	params := map[string]string{"name": "John"}

	// Test getting a text with parameters in English
	text := englishLang.GetTextWithParam("greeting", params)
	assert.Equal(t, "Hello John!", text)

	// Test getting a text with parameters in French
	frenchLang, err := i18nInstance.GetLang("fr")
	assert.NoError(t, err)

	text = frenchLang.GetTextWithParam("greeting", params)
	assert.Equal(t, "Salut John!", text)
}
