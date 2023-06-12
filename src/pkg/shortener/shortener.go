package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"github.com/manlikeNacho/Sissors/src/models"
	"math/big"
)

// sha256Of to hash the initial input of the link, to maintain consistency and uniqueness.
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// base58Encoded the encoding algorithm.
func base58Encoded(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

// GenerateShortLink to generate short Link.
func GenerateShortLink(m models.Url) (string, error) {
	if m.ShortUrl == "" {
		urlHashBytes := sha256Of(m.Url)
		generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
		finalString, err := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
		if err != nil {
			return "", err
		}
		return finalString[:8], nil
	}
	return m.ShortUrl, nil
}
