package main

import (
	"fmt"
	go_reverso_api "github.com/BRUHItsABunny/go-reverso-api"
)

func main() {
	// init client
	client := go_reverso_api.GetReversoClient()
	// init languages
	langs := go_reverso_api.GetLanguages()
	// Translate
	fmt.Println(client.Translate("bunny", langs[go_reverso_api.LanguageEnglish], langs[go_reverso_api.LanguageFrench]))
	// Speak
	fmt.Println(client.Speak("son petit lapin", go_reverso_api.VoiceFrenchMale, 128, 85))
	// Synonyms
	fmt.Println(client.Synonyms("bunny", langs[go_reverso_api.LanguageEnglish]))
	// AutoComplete
	fmt.Println(client.AutoComplete("bunny", langs[go_reverso_api.LanguageEnglish]))
	// Query
	fmt.Println(client.Query("sky", langs[go_reverso_api.LanguageEnglish], langs[go_reverso_api.LanguageFrench], 1))
	// Suggest
	fmt.Println(client.Suggest("sky", langs[go_reverso_api.LanguageEnglish], langs[go_reverso_api.LanguageFrench]))
}
