package client

import (
	"context"
	"fmt"
	andutils "github.com/BRUHItsABunny/go-android-utils"
	"github.com/BRUHItsABunny/go-reverso-api/api"
	"net/http"
)

type ReversoClient struct {
	Client *http.Client
	Device *andutils.Device
}

func NewReversoClient(hClient *http.Client, device *andutils.Device) *ReversoClient {
	if hClient == nil {
		hClient = http.DefaultClient
	}

	if device == nil {
		device = andutils.GetRandomDevice()
	}

	return &ReversoClient{
		Client: hClient,
		Device: device,
	}
}

func (c *ReversoClient) Query(ctx context.Context, text string, srcLang, dstLang *api.Language, page int) (*api.QueryResponse, error) {
	req, err := api.QueryRequest(ctx, c.Device, text, srcLang, dstLang, page)
	if err != nil {
		return nil, fmt.Errorf("api.QueryRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.QueryResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *ReversoClient) Suggest(ctx context.Context, text string, srcLang, dstLang *api.Language) (*api.SuggestResponse, error) {
	req, err := api.SuggestRequest(ctx, c.Device, text, srcLang, dstLang)
	if err != nil {
		return nil, fmt.Errorf("api.SuggestRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.SuggestResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *ReversoClient) Translate(ctx context.Context, request *api.TranslateRequestBody) (*api.TranslateResponse, error) {
	req, err := api.TranslateRequest(ctx, c.Device, request)
	if err != nil {
		return nil, fmt.Errorf("api.TranslateRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.TranslateResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *ReversoClient) Synonyms(ctx context.Context, text string, lang *api.Language, rude, abc, merge bool, limit int) (*api.SynonymsResponse, error) {
	req, err := api.SynonymsRequest(ctx, c.Device, text, lang, rude, abc, merge, limit)
	if err != nil {
		return nil, fmt.Errorf("api.SynonymsRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.SynonymsResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *ReversoClient) AutoComplete(ctx context.Context, text string, lang *api.Language, rude, colloquial bool) (*api.AutoCompleteResponse, error) {
	req, err := api.AutoCompleteRequest(ctx, c.Device, text, lang, rude, colloquial)
	if err != nil {
		return nil, fmt.Errorf("api.AutoCompleteRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.AutoCompleteResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *ReversoClient) Rephrase(ctx context.Context, text string, lang *api.Language) (*api.RephraseResponse, error) {
	req, err := api.RephraseRequest(ctx, c.Device, text, lang)
	if err != nil {
		return nil, fmt.Errorf("api.RephraseRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.RephraseResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *ReversoClient) Voice(ctx context.Context, text, voice string, bitrate, voiceSpeed int) ([]byte, error) {
	req, err := api.VoiceRequest(ctx, c.Device, text, voice, bitrate, voiceSpeed)
	if err != nil {
		return nil, fmt.Errorf("api.VoiceRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.BytesParser(resp)
	if err != nil {
		return nil, fmt.Errorf("api.BytesParser: %w", err)
	}
	return result, nil
}

func (c *ReversoClient) Conjugate(ctx context.Context, text string, lang *api.Language) (*api.ConjugatorResponse, error) {
	req, err := api.ConjugateRequest(ctx, c.Device, text, lang)
	if err != nil {
		return nil, fmt.Errorf("api.ConjugateRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.ConjugatorResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}
