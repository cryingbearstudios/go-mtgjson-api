package mtgJson

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"

	"github.com/cryingbearstudios/go-mtgjson-api/jsonStreaming"
	"github.com/cryingbearstudios/go-mtgjson-api/model"
	"github.com/cryingbearstudios/go-mtgjson-api/util"
)

// These data types are meant to represent the semantic model of the downloaded files, but should likely not actually
// be used directly for ser/de purposes, due to the size of the data files they describe. For efficiency reasons,
// we want to use a streaming decoding strategy of the Data payload instead.
type dataFile[keyType comparable, ValueType any] struct {
	Meta model.Meta            `json:"meta"`
	Data map[keyType]ValueType `json:"data"`
}

type _ dataFile[util.SetCode, model.Set] // all printings

func IgnoreMetadata(meta *model.Meta) error {
	return nil
}

func LogMetadata(meta *model.Meta) error {
	slog.Info("metadata", "date", meta.Date, "version", meta.Version)
	return nil
}

type metaCb func(*model.Meta) error
type dataCb func(context.Context, json.Token, *json.Decoder) error

func ProcessDataFile(ctx context.Context, input io.Reader, metadataCallback metaCb, dataCallback dataCb) error {
	dec := json.NewDecoder(input)
	if err := jsonStreaming.AssertNextTokens(dec, json.Delim('{'), "meta"); err != nil {
		return err
	}
	var meta model.Meta
	if err := dec.Decode(&meta); err != nil {
		return err
	}
	if err := metadataCallback(&meta); err != nil {
		return err
	}
	if err := jsonStreaming.AssertNextTokens(dec, "data"); err != nil {
		return err
	}
	err := jsonStreaming.ProcessObjectByProperties(dec, func(key json.Token) error {
		return dataCallback(ctx, key, dec)
	})
	if err != nil {
		return err
	}
	return jsonStreaming.AssertNextTokens(dec, json.Delim('}'))
}
