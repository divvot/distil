package distil

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"math/rand/v2"
	"strings"

	"github.com/Hyper-Solutions/orderedobject"
)

func Encrypt(data string, magic uint32) (string, error) {

	// This code sucks

	var temporal map[string]any

	err := json.Unmarshal([]byte(data), &temporal)
	if err != nil {
		return "", errors.New("invalid payload")
	}

	// Order matters)?
	orden := []string{"application_manifest", "old_token", "session", "answer"}
	ordered := orderedobject.NewObject[any](len(temporal))
	for _, k := range orden {
		if v, ok := temporal[k]; ok {
			ordered.Set(k, v)
		}
	}

	// JSON format matters)?
	payloadBytes, _ := json.MarshalIndent(ordered, "", "  ")
	payload := strings.ReplaceAll(string(payloadBytes), "\": ", "\" : ")

	originalLen := len(payload)
	if (len(payload) % 2) == 1 {
		// Payload length needs to be odd
		payload += "="
	}

	if (len(payload) % 4) != 0 {
		return "", errors.New("insufficient length")
	}

	paddedLength := len(payload)

	if magic == 0 {
		magic = rand.Uint32()
	}

	bytes := []byte(payload)

	buffer := make([]byte, 0)

	t := magicEncryptNumber * magic
	buffer = binary.BigEndian.AppendUint32(buffer, t)
	buffer = binary.BigEndian.AppendUint32(buffer, uint32(originalLen)*uint32(magicEncryptNumber)^magic)

	for i := 0; paddedLength > i; i += 4 {
		chunk := bytes[i : i+4]

		j := binary.LittleEndian.Uint32(chunk)
		obscured := (uint32(magicEncryptNumber) * j) ^ magic
		magic = j

		buffer = binary.BigEndian.AppendUint32(buffer, obscured)
	}

	return "2." + base64.URLEncoding.EncodeToString(buffer), nil
}

// Slow function, It takes about 10 minutes to decrypt the second payload
// and 3 - 7 minutes for first payload
func Decrypt(data string) (string, uint32) {
	decoded, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}

	chunk := decoded[0:4]
	actual := binary.BigEndian.Uint32(chunk)

	magic := crack(actual)

	chunk = decoded[4:8]
	actual = binary.BigEndian.Uint32(chunk)

	length := crack(actual ^ magic)
	decoded = decoded[8:]
	decrypted := make([]byte, 0)

	for i := 0; len(decoded) > i; i += 4 {
		chunk = decoded[i : i+4]

		actual = binary.BigEndian.Uint32(chunk)
		cracked := crack(actual ^ magic)

		decrypted = binary.LittleEndian.AppendUint32(decrypted, cracked)
		magic = cracked
	}

	return string(decrypted[:length]), magic
}
