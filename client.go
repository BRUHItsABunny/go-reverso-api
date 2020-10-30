package go_reverso_api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetReversoClient() *ReversoClient {
	httpClient := gokhttp.GetHTTPClient(nil)
	return &ReversoClient{Client: &httpClient}
}

func GetLanguages() map[string]*Language {
	result := make(map[string]*Language, 0)
	languageStr := "{\"arabic\":{\"code\":\"ar\",\"alpha3\":\"ara\"},\"german\":{\"code\":\"de\",\"alpha3\":\"ger\"},\"english\":{\"code\":\"en\",\"alpha3\":\"eng\"},\"spanish\":{\"code\":\"es\",\"alpha3\":\"spa\"},\"french\":{\"code\":\"fr\",\"alpha3\":\"fra\"},\"portuguese\":{\"code\":\"pt\",\"alpha3\":\"por\"},\"italian\":{\"code\":\"it\",\"alpha3\":\"ita\"},\"dutch\":{\"code\":\"nl\",\"alpha3\":\"dut\"},\"russian\":{\"code\":\"ru\",\"alpha3\":\"rus\"},\"hebrew\":{\"code\":\"he\",\"alpha3\":\"heb\"},\"polish\":{\"code\":\"pl\",\"alpha3\":\"pol\"},\"romanian\":{\"code\":\"ro\",\"alpha3\":\"rum\"},\"japanese\":{\"code\":\"ja\",\"alpha3\":\"jap\"},\"turkish\":{\"code\":\"tr\",\"alpha3\":\"tur\"},\"chinese\":{\"code\":\"zh\",\"alpha3\":\"chi\"}}"
	_ = json.Unmarshal([]byte(languageStr), &result)
	return result
}

func (rc *ReversoClient) Query(text string, srcLang, dstLang *Language, page int) (*SearchResponse, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = new(SearchResponse)
	)

	params := url.Values{
		"source_text": []string{text},
		"source_lang": []string{srcLang.Code},
		"target_lang": []string{dstLang.Code},
		"npage":       []string{strconv.Itoa(page)},
		"nrows":       []string{"4"},
		"expr_sug":    []string{"1"},
		"json":        []string{"1"},
		"dym_apply":   []string{"true"},
		"pos_reorder": []string{"5"},
	}

	headers := map[string]string{
		"User-Agent": userAgentContextApp,
	}

	req, err = rc.Client.MakePOSTRequestWithParameters(urlContextQuery, params, map[string]string{}, headers)
	if err == nil {
		resp, err = rc.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (rc *ReversoClient) Suggest(text string, srcLang, dstLang *Language) (*SuggestResponse, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = new(SuggestResponse)
	)

	params := url.Values{
		"search":      []string{text},
		"source_lang": []string{srcLang.Code},
		"target_lang": []string{dstLang.Code},
	}

	headers := map[string]string{
		"User-Agent": userAgentContextApp,
	}

	req, err = rc.Client.MakeGETRequest(urlContextSuggest, params, headers)
	if err == nil {
		resp, err = rc.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (rc *ReversoClient) Translate(text string, srcLang, dstLang *Language) (*TranslateResponse, error) {
	var (
		err           error
		req           *http.Request
		resp          *gokhttp.HttpResponse
		result        = new(TranslateResponse)
		postBodyBytes []byte
	)

	postBody := TranslateRequest{
		Input:  text,
		From:   srcLang.Alpha3,
		To:     dstLang.Alpha3,
		Format: "text",
		Options: TranslateOptions{
			Origin:           "contextappandroid", // contextappandroid vs reversomobile
			SentenceSplitter: false,
		},
	}
	if len(strings.Split(text, " ")) == 1 {
		postBody.Options.ContextResults = true
		postBody.Options.LanguageDetection = true
	}

	postBodyBytes, err = json.Marshal(&postBody)

	headers := map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
		"User-Agent":   "",
	}
	if err == nil {
		req, err = rc.Client.MakeRawPOSTRequest(urlTranslate, url.Values{}, bytes.NewReader(postBodyBytes), headers)
		if err == nil {
			resp, err = rc.Client.Do(req)
			if err == nil {
				err = resp.Object(&result)
				return result, err
			}
		}
	}
	return nil, err
}

func (rc *ReversoClient) Synonyms(text string, language *Language) (*SynonymsResponse, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = new(SynonymsResponse)
	)

	params := url.Values{
		"limit": []string{"50"},
		"rude":  []string{"true"},
		"abc":   []string{"false"},
		"merge": []string{"true"},
	}

	headers := map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
		"User-Agent":   "",
	}

	req, err = rc.Client.MakeGETRequest(urlSynonyms+endpointSynonymsSearch+language.Code+"/"+url.PathEscape(text), params, headers)
	if err == nil {
		resp, err = rc.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (rc *ReversoClient) AutoComplete(text string, language *Language) ([]string, error) {
	var (
		err    error
		req    *http.Request
		resp   *gokhttp.HttpResponse
		result = make([]string, 0)
	)

	params := url.Values{
		"term":       []string{text},
		"rude":       []string{"true"},
		"colloquial": []string{"true"},
	}

	headers := map[string]string{
		"Content-Type":      "application/json",
		"User-Agent":        userAgentSynonymsApp,
		"x-reverso-origin":  "synonymapp",
		"x-reverso-ui-lang": "en",
		"authorization":     "Basic " + bearerSynonyms,
	}

	req, err = rc.Client.MakeGETRequest(urlSynonyms+endpointSynonymsAutoComplete+language.Code+"/", params, headers)
	if err == nil {
		resp, err = rc.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			return result, err
		}
	}
	return nil, err
}

func (rc *ReversoClient) Speak(text, voice string, bitRate, voiceSpeed int) (string, error) {
	var (
		err     error
		req     *http.Request
		resp    *gokhttp.HttpResponse
		result  = make([]byte, 0)
		fileOut *os.File
	)

	fileName := "pronunciation_" + strconv.FormatInt(time.Now().Unix(), 10) + ".mp3"
	fileOut, err = os.Create(fileName)

	if err == nil {
		params := url.Values{
			"inputText":  []string{base64.URLEncoding.EncodeToString([]byte(text))},
			"mp3BitRate": []string{strconv.Itoa(bitRate)},
			"voiceSpeed": []string{strconv.Itoa(voiceSpeed)},
		}

		headers := map[string]string{
			"User-Agent": userAgentContextApp,
		}

		req, err = rc.Client.MakeGETRequest(urlVoice+voice, params, headers)
		if err == nil {
			resp, err = rc.Client.Do(req)
			if err == nil {
				result, err = resp.Bytes()
				if err == nil {
					_, err = fileOut.Write(result)
					return fileName, err
				}
			}
		}
	}
	return "", err
}
