package distil

import (
	"encoding/base64"
	"encoding/json"

	"github.com/divvot/distil/model"
)

func solve(question_string string, magic uint32) string {
	question := []byte(question_string)

	lenP := len(question)
	lenO := len(question)

	for ok := true; ok; {
		xor := magic ^ uint32(question[lenO-1])
		lenT := lenO
		if xor != 0 {
			for stop := true; stop; {
				question[lenT%lenP] += byte(xor)
				stop = xor >= 4
				lenT += 1
				xor = xor >> 2
			}

			ok = 1 < lenO
			lenO -= 1
		}
	}

	return base64.URLEncoding.EncodeToString(question)
}

func calculateMagic(fingerprint model.Fingerprint) uint32 {
	var secret uint32 = 1571 * 5

	secret += sumChars(fingerprint.BundleIdentifier)
	secret += sumChars(fingerprint.BundleVersion)
	secret += sumChars(fingerprint.Model)
	secret += sumChars(fingerprint.Name)
	secret += sumChars(fingerprint.VenderId)
	secret += sumChars(fingerprint.Version)
	secret += sumChars(fingerprint.CountryCode)
	secret += sumChars(fingerprint.LanguageCode)
	secret += uint32(fingerprint.Device.Width) + uint32(fingerprint.Device.Height)

	var ipAddresses map[string]string

	b, _ := json.Marshal(fingerprint.IpAddresses)
	json.Unmarshal(b, &ipAddresses)

	// I don't gonna add next params, not revelant I guess
	// ident6cfdc2e2399f0cf722b3f6de3075c54d --> is jailbreak
	// ident0073c744e76e32a46d36953b1d849e5c --> is simulator
	// ident61fa153f2827c887a48a351ae3c6cfd3 --> is debugger
	// (for tbh I dont remember each but that's the idea lol)

	for k, v := range ipAddresses {
		secret += sumChars(k) + sumChars(v)
	}

	return secret
}
