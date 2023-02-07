package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/BRUHItsABunny/gOkHttp/requests"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"github.com/BRUHItsABunny/go-reverso-api/constants"
	"net/http"
	"net/url"
	"strconv"
)

func QueryRequest(ctx context.Context, device *andutils.Device, text string, srcLang, dstLang *Language, page int) (*http.Request, error) {
	parameters := DefaultQueryParameters()
	parameters["source_text"] = []string{text}
	parameters["source_lang"] = []string{srcLang.Code}
	parameters["target_lang"] = []string{dstLang.Code}
	parameters["npage"] = []string{strconv.Itoa(page)}

	req, err := requests.MakePOSTRequest(ctx, constants.EndpointContextQuery,
		requests.NewHeaderOption(DefaultQueryHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakePOSTRequest: %w", err)
	}
	return req, nil
}

func SuggestRequest(ctx context.Context, device *andutils.Device, text string, srcLang, dstLang *Language) (*http.Request, error) {
	parameters := url.Values{
		"search":      {text},
		"source_lang": {srcLang.Code},
		"target_lang": {dstLang.Code},
	}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointContextSuggest,
		requests.NewHeaderOption(DefaultQueryHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func TranslateRequest(ctx context.Context, device *andutils.Device, request *TranslateRequestBody) (*http.Request, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := requests.MakePOSTRequest(ctx, constants.EndpointTranslate,
		requests.NewHeaderOption(DefaultTranslateHeaders(device)),
		requests.NewPOSTJSONOption(data, true),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakePOSTRequest: %w", err)
	}
	return req, nil
}

func SynonymsRequest(ctx context.Context, device *andutils.Device, text string, lang *Language, rude, abc, merge bool, limit int) (*http.Request, error) {
	parameters := url.Values{
		"limit": {strconv.Itoa(limit)},
		"rude":  {strconv.FormatBool(rude)},
		"abc":   {strconv.FormatBool(abc)},
		"merge": {strconv.FormatBool(merge)},
	}

	headers := DefaultSynonymsHeaders(device, true)
	headers["x-reverso-ui-lang"] = []string{lang.Code}

	req, err := requests.MakeGETRequest(ctx, fmt.Sprintf("%s%s/%s", constants.EndpointSynonymsSearch, lang.Code, url.QueryEscape(text)),
		requests.NewHeaderOption(headers),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func AutoCompleteRequest(ctx context.Context, device *andutils.Device, text string, lang *Language, rude, colloquial bool) (*http.Request, error) {
	parameters := url.Values{
		"term":       {text},
		"rude":       {strconv.FormatBool(rude)},
		"colloquial": {strconv.FormatBool(colloquial)},
	}

	headers := DefaultSynonymsHeaders(device, false)
	headers["x-reverso-ui-lang"] = []string{lang.Code}

	req, err := requests.MakeGETRequest(ctx, fmt.Sprintf("%s%s/%s", constants.EndpointSynonymsAutoComplete, lang.Code, url.QueryEscape(text)),
		requests.NewHeaderOption(headers),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func RephraseRequest(ctx context.Context, device *andutils.Device, text string, lang *Language) (*http.Request, error) {
	parameters := url.Values{
		"sentence":      {text},
		"clientVersion": {device.Version.ToAndroidVersion()},
		"language":      {lang.Code},
	}

	headers := DefaultSynonymsHeaders(device, false)
	headers["x-reverso-ui-lang"] = []string{lang.Code}
	delete(headers, "content-type")

	req, err := requests.MakeGETRequest(ctx, constants.EndpointRephrase,
		requests.NewHeaderOption(DefaultQueryHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func VoiceRequest(ctx context.Context, device *andutils.Device, text, voice string, bitrate, voiceSpeed int) (*http.Request, error) {
	parameters := url.Values{
		"inputText":  {base64.URLEncoding.EncodeToString([]byte(text))},
		"mp3BitRate": {strconv.Itoa(bitrate)},
		"voiceSpeed": {strconv.Itoa(voiceSpeed)},
	}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointVoice+voice,
		requests.NewHeaderOption(DefaultVoiceHeaders(device)),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func ConjugateRequest(ctx context.Context, device *andutils.Device, text string, lang *Language) (*http.Request, error) {
	req, err := requests.MakeGETRequest(ctx, fmt.Sprintf("%s%s/%s", constants.EndpointConjugator, lang.Code, url.QueryEscape(text)),
		requests.NewHeaderOption(DefaultQueryHeaders(device)),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}
