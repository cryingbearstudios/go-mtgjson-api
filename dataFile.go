package mtgJson

import (
	"encoding/json"
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

type MetadataCallback func(meta model.Meta) error

var IgnoreMetadata MetadataCallback = func(meta model.Meta) error {
	return nil
}
var LogMetadata MetadataCallback = func(meta model.Meta) error {
	slog.Info("metadata", "date", meta.Date, "version", meta.Version)
	return nil
}

func PreprocessDataFile(dec *json.Decoder, cb MetadataCallback) error {
	// read start of root object and the "meta" key
	if err := jsonStreaming.AssertNextTokens(dec, json.Delim('{'), "meta"); err != nil {
		return err
	}
	// read file metadata
	var meta model.Meta
	if err := dec.Decode(&meta); err != nil {
		return err
	}
	// pass metadata to callback
	if err := cb(meta); err != nil {
		return err
	}
	// read 'data' key and start of 'data' object
	return jsonStreaming.AssertNextTokens(dec, "data", json.Delim('{'))
}
