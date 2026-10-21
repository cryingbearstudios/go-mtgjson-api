package jsonStreaming

import (
	"encoding/json"
	"fmt"
)

func AssertNextTokens(dec *json.Decoder, expectedTokens ...json.Token) error {
	for _, expectedTok := range expectedTokens {
		if tok, err := dec.Token(); err != nil { // read first key
			return err
		} else if tok != expectedTok {
			return fmt.Errorf("expected '%v' key, got %v", expectedTok, tok)
		}
	}
	return nil
}

func ProcessArray(dec *json.Decoder, cb func(int) error) error {
	if err := AssertNextTokens(dec, json.Delim('[')); err != nil {
		return err
	}
	index := 0
	for dec.More() {
		if err := cb(index); err != nil {
			return err
		}
		index++
	}
	return AssertNextTokens(dec, json.Delim(']'))
}

func ProcessObjectByProperties(dec *json.Decoder, cb func(key json.Token) error) error {
	if err := AssertNextTokens(dec, json.Delim('{')); err != nil {
		return err
	}
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return err
		}
		if err := cb(key); err != nil {
			return err
		}
	}
	return AssertNextTokens(dec, json.Delim('}'))
}
