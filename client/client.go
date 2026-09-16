package client

import (
	"compress/gzip"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

	mtgJson "cryingbear.net/go-mtgjson-api"
	"cryingbear.net/go-mtgjson-api/model"
	"github.com/google/uuid"
)

const (
	baseUrl          = "https://mtgjson.com/api/v5"
	allPrintingsGz   = "AllPrintings.json.gz"
	allPrintingsSha  = allPrintingsGz + ".sha256"
	tcgPlayerSkusGz  = "TcgplayerSkus.json.gz"
	tcgPlayerSkusSha = tcgPlayerSkusGz + ".sha256"
)

func GetExpectedSha() ([]byte, error) {
	resp, err := http.Get(path.Join(baseUrl, tcgPlayerSkusSha))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	shaText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result, err := hex.DecodeString(string(shaText))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetTcgPlayerSkus(cb func(item *mtgJson.SkusById) error) error {
	expectedSha, err := GetExpectedSha()
	if err != nil {
		return err
	}
	resp, err := http.Get(path.Join(baseUrl, tcgPlayerSkusGz))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	hasher := sha256.New()

	dec := json.NewDecoder(io.TeeReader(gzipReader, hasher))
	err = mtgJson.PreprocessDataFile(dec, mtgJson.IgnoreMetadata)
	if err != nil {
		return err
	}

	for dec.More() {
		var key uuid.UUID
		if err := dec.Decode(&key); err != nil {
			return err
		}
		var val []model.TcgplayerSKUs
		if err := dec.Decode(&val); err != nil {
			return err
		}
		item := mtgJson.SkusById{
			Id:   key,
			SKUs: val,
		}
		if err := cb(&item); err != nil {
			return err
		}
	}

	actualHash := hasher.Sum(nil)
	if subtle.ConstantTimeCompare(expectedSha, actualHash) != 1 {
		return fmt.Errorf("sha256 sum mismatch, expected %x, got %x", expectedSha, actualHash)
	}
	return nil
}
