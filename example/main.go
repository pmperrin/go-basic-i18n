package main

import (
	"fmt"

	basicI18n "github.com/pmperrin/go-basic-i18n/i18n"
)

func main() {
	i18nInstance, err := basicI18n.InitI18n("./example/local", "messages", "fr")
	if err != nil {
		fmt.Println(err.Error())
	}
	lang, err := i18nInstance.GetLang("en")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(lang.GetText("signup.title"))
	fmt.Println(lang.GetText("signup.button"))
}
