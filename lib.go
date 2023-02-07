package go_reverso_api

import (
	"encoding/json"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"github.com/BRUHItsABunny/go-reverso-api/api"
	"github.com/BRUHItsABunny/go-reverso-api/client"
	"net/http"
)

func GetLanguages() map[string]*api.Language {
	result := make(map[string]*api.Language, 0)
	languageStr := "{\"arabic\":{\"code\":\"ar\",\"alpha3\":\"ara\"},\"german\":{\"code\":\"de\",\"alpha3\":\"ger\"},\"english\":{\"code\":\"en\",\"alpha3\":\"eng\"},\"spanish\":{\"code\":\"es\",\"alpha3\":\"spa\"},\"french\":{\"code\":\"fr\",\"alpha3\":\"fra\"},\"portuguese\":{\"code\":\"pt\",\"alpha3\":\"por\"},\"italian\":{\"code\":\"it\",\"alpha3\":\"ita\"},\"dutch\":{\"code\":\"nl\",\"alpha3\":\"dut\"},\"russian\":{\"code\":\"ru\",\"alpha3\":\"rus\"},\"hebrew\":{\"code\":\"he\",\"alpha3\":\"heb\"},\"polish\":{\"code\":\"pl\",\"alpha3\":\"pol\"},\"romanian\":{\"code\":\"ro\",\"alpha3\":\"rum\"},\"japanese\":{\"code\":\"ja\",\"alpha3\":\"jap\"},\"turkish\":{\"code\":\"tr\",\"alpha3\":\"tur\"},\"chinese\":{\"code\":\"zh\",\"alpha3\":\"chi\"}}"
	_ = json.Unmarshal([]byte(languageStr), &result)
	return result
}

func NewReversoClient(hClient *http.Client, device *andutils.Device) *client.ReversoClient {
	return client.NewReversoClient(hClient, device)
}
