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

func ReadStringToken(dec *json.Decoder) (string, error) {
	setCodeToken, err := dec.Token() // read key (set code)
	if err != nil {
		return "", err
	}
	setCode, ok := setCodeToken.(string)
	if !ok {
		return "", fmt.Errorf("expected set code as string, got %v", setCodeToken)
	}
	return setCode, nil
}

func ReadStringAtPath(dec *json.Decoder, path ...json.Token) (string, error) {
	if err := AssertNextTokens(dec, path); err != nil {
		return "", err
	}
	// Return the value of the next token as a string
	return ReadStringToken(dec)
}

func ProcessArray(dec *json.Decoder, cb func() error) error {
	if err := AssertNextTokens(dec, json.Delim('[')); err != nil {
		return err
	}
	for dec.More() {
		if err := cb(); err != nil {
			return err
		}
	}
	return AssertNextTokens(dec, json.Delim(']'))
}

func ProcessOnlyOnePropertyInObject(dec *json.Decoder, property string, cb func() error) error {
	return ProcessObjectByProperties(dec, func(key json.Token) error {
		if key != property {
			// read and discard the value for this key
			if err := dec.Decode(new(any)); err != nil {
				return err
			}
			return nil
		}
		return cb()
	})
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

