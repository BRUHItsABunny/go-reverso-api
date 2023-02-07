package api

import (
	"fmt"
	"github.com/BRUHItsABunny/gOkHttp/responses"
	"net/http"
)

func ErrorParser(resp *http.Response, fallbackErr error) error {
	errRsp := new(ErrorResponse)
	err := responses.ResponseJSON(resp, errRsp)
	if err != nil {
		return fmt.Errorf("responses.CheckHTTPCode: %w", fallbackErr)
	}
	return errRsp
}

func InterfaceParser(resp *http.Response) (interface{}, error) {
	result := new(interface{})
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func BytesParser(resp *http.Response) ([]byte, error) {
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	result, err := responses.ResponseBytes(resp)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}

func ObjectParser[T any](obj T, resp *http.Response) (*T, error) {
	result := new(T)
	err := responses.CheckHTTPCode(resp, 200)
	if err != nil {
		return nil, ErrorParser(resp, err)
	}

	err = responses.ResponseJSON(resp, result)
	if err != nil {
		return nil, fmt.Errorf("responses.ResponseJSON: %w", err)
	}
	return result, nil
}
