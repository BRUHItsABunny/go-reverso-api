package api

import (
	constants2 "github.com/BRUHItsABunny/gOkHttp/constants"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"github.com/BRUHItsABunny/go-reverso-api/constants"
	"net/http"
	"net/url"
)

func DefaultQueryParameters() url.Values {
	return url.Values{
		"nrows":       {"4"},
		"expr_sug":    {"1"},
		"json":        {"1"},
		"dym_apply":   {"true"},
		"pos_reorder": {"5"},
	}
}

func DefaultQueryHeaders(device *andutils.Device) http.Header {
	return http.Header{
		"user-agent": {"Dalvik/2.1.0 " + device.GetUserAgent() + " ReversoContext"},
	}
}

func DefaultTranslateHeaders(device *andutils.Device) http.Header {
	return http.Header{
		"user-agent": {"Dalvik/2.1.0 " + device.GetUserAgent() + " ReversoContext 10.7.0.10000081"},
	}
}

func DefaultSynonymsHeaders(device *andutils.Device, isSearch bool) http.Header {
	result := http.Header{
		"x-reverso-ui-lang": {"en"},
		"x-reverso-origin":  {"synonyms.app.android"},
		"user-agent":        {"Dalvik/2.1.0 " + device.GetUserAgent() + " ReversoSynonyms"},
		"authorization":     {"Basic " + constants.AuthSynonyms},
	}
	if isSearch {
		result["content-type"] = []string{constants2.MIMEApplicationJSON}
		result["os"] = []string{"Android " + device.Version.ToAndroidVersion()}
		result["origin"] = []string{"synonyms.app.android"}
	}
	return result
}

func DefaultVoiceHeaders(device *andutils.Device) http.Header {
	return http.Header{
		"user-agent": {"stagefright/1.2 (Linux;Android " + device.Version.ToAndroidVersion() + ")"},
	}
}
