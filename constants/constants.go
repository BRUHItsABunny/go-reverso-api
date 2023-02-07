package constants

const (
	urlSynonyms                  = "https://synonyms.reverso.net/api/v2/"
	EndpointSynonymsSearch       = urlSynonyms + "search/"
	EndpointSynonymsAutoComplete = urlSynonyms + "autocomplete/"

	urlAccount              = "https://account.reverso.net/api/v1/account/"
	EndpointAccountLogin    = urlAccount + "Login"
	EndpointAccountRegister = urlAccount + "Register"

	urlContext             = "https://context.reverso.net/"
	EndpointContextSuggest = urlContext + "bst-suggest-service"
	EndpointContextQuery   = urlContext + "bst-query-service"

	EndpointTranslate  = "https://api.reverso.net/translate/v1/translation"
	EndpointConjugator = "https://api-conjugator.reverso.net/api/conjugator/"
	EndpointRephrase   = "https://rephraser-api.reverso.net/v1/rephrase"
	EndpointVoice      = "https://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetVoiceStream/voiceName="

	AuthSynonyms = "c3lub255bXM6REtiZTUyRjNZRExZdVFFOHlk" // synonyms:DKbe52F3YDLYuQE8yd
)
