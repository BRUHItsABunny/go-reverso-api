package go_reverso_api

import (
	"context"
	"fmt"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"github.com/BRUHItsABunny/go-reverso-api/api"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

// Loads the .env file in our repo - if you have one
func loadTestEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

var languages = GetLanguages()

func TestQuery(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	text := os.Getenv("VAL_TEXT")
	srcLang := languages[os.Getenv("VAL_SRC_LANG")]
	dstLang := languages[os.Getenv("VAL_DST_LANG")]
	page, err := strconv.Atoi(os.Getenv("VAL_PAGE"))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewReversoClient(hClient, nil)
	res, err := c.Query(context.Background(), text, srcLang, dstLang, page)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestSuggest(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	text := os.Getenv("VAL_WORD")
	srcLang := languages[os.Getenv("VAL_SRC_LANG")]
	dstLang := languages[os.Getenv("VAL_DST_LANG")]
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewReversoClient(hClient, nil)
	res, err := c.Suggest(context.Background(), text, srcLang, dstLang)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestTranslate(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	text := os.Getenv("VAL_TEXT")
	srcLang := languages[os.Getenv("VAL_SRC_LANG")]
	dstLang := languages[os.Getenv("VAL_DST_LANG")]
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewReversoClient(hClient, nil)
	request := api.NewTranslateRequestBody(text, srcLang, dstLang)
	res, err := c.Translate(context.Background(), request)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestSynonyms(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	text := os.Getenv("VAL_WORD")
	lang := languages[os.Getenv("VAL_SRC_LANG")]
	rude := os.Getenv("VAL_RUDE") == "true"
	abc := os.Getenv("VAL_ABC") == "true"
	merge := os.Getenv("VAL_MERGE") == "true"
	limit, err := strconv.Atoi(os.Getenv("VAL_LIMIT"))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewReversoClient(hClient, nil)
	res, err := c.Synonyms(context.Background(), text, lang, rude, abc, merge, limit)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestAutoComplete(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	text := os.Getenv("VAL_WORD")
	lang := languages[os.Getenv("VAL_SRC_LANG")]
	rude := os.Getenv("VAL_RUDE") == "true"
	colloquial := os.Getenv("VAL_COLLOQUIAL") == "true"
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewReversoClient(hClient, nil)
	res, err := c.AutoComplete(context.Background(), text, lang, rude, colloquial)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestRephrase(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	text := os.Getenv("VAL_TEXT")
	lang := languages[os.Getenv("VAL_SRC_LANG")]
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewReversoClient(hClient, nil)
	res, err := c.Rephrase(context.Background(), text, lang)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestConjugate(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	text := os.Getenv("VAL_VERB")
	lang := languages[os.Getenv("VAL_SRC_LANG")]
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewReversoClient(hClient, nil)
	res, err := c.Conjugate(context.Background(), text, lang)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestVoice(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	text := os.Getenv("VAL_TEXT")
	voice := os.Getenv("VAL_VOICE")
	bitrate, err := strconv.Atoi(os.Getenv("VAL_BITRATE"))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	speed, err := strconv.Atoi(os.Getenv("VAL_VOICESPEED"))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewReversoClient(hClient, nil)
	res, err := c.Voice(context.Background(), text, voice, bitrate, speed)
	if err == nil {
		os.Remove("test.mp3")
		os.WriteFile("test.mp3", res, 0666)
	} else {
		fmt.Println("err: ", err)
	}
}
