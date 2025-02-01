package distil

import (
	"encoding/json"
	"strings"

	"github.com/divvot/distil/model"

	"github.com/google/uuid"
)

// fingerprint provider

const (
	debugFingerprint = `{ "uuid": "69EB7A23-F97D-4674-B847-CFDDD5DA6F94", "vendorId": "6B7F7A6B-9065-46EC-B967-92E42526CD88", "ipAddresses": { "utun3/ipv6": "fe80::ce81:b1c:bd2c:69e", "anpi0/ipv6": "fe80::9009:43ff:fe06:3343", "en0/ipv4": "192.168.0.10", "utun2/ipv6": "fe80::be20:f7cd:5bba:d9ab", "lo0/ipv6": "fe80::1", "awdl0/ipv6": "fe80::8c75:ff:fe9f:f817", "llw0/ipv6": "fe80::8c75:ff:fe9f:f817", "utun1/ipv6": "fe80::81c4:e42b:12de:67d4", "en0/ipv6": "fe80::1cf1:7cb1:16c1:1665", "utun0/ipv6": "fe80::eddb:ea6f:6b28:6518", "utun4/ipv6": "fe80::1c36:4767:b2ec:b690", "lo0/ipv4": "127.0.0.1" }, "version": "16.7.8", "bundleIdentifier": "uk.co.coop.app", "languageCode": "en", "countryCode": "SV", "bundleVersion": "1.53.0", "isSimulator": false, "isJailBroken": false, "debugger": false, "device": { "width": 375, "height": 812 }, "model": "iPhone", "name": "" }`
)

func DebugFingerprint() model.Fingerprint {
	var fp model.Fingerprint
	json.Unmarshal([]byte(debugFingerprint), &fp)
	return fp
}

func GenerateFingerprint(bundleId, bundleVersion, iosVersion, languageCode, countryCode string) model.Fingerprint {

	// I'll just modify the fingeprint debug
	sample := DebugFingerprint()

	if iosVersion == "" {
		iosVersion = sample.Version
	}

	if languageCode == "" {
		iosVersion = sample.LanguageCode
	}
	if countryCode == "" {
		iosVersion = sample.CountryCode
	}

	sample.IpAddresses.Anpi0Ipv6 = generateIPv6()
	sample.IpAddresses.Awdl0Ipv6 = generateIPv6()
	sample.IpAddresses.En0Ipv6 = generateIPv6()
	sample.IpAddresses.Llw0Ipv6 = generateIPv6()
	sample.IpAddresses.Lo0Ipv6 = generateIPv6()
	sample.IpAddresses.Utun0Ipv6 = generateIPv6()
	sample.IpAddresses.Utun1Ipv6 = generateIPv6()
	sample.IpAddresses.Utun2Ipv6 = generateIPv6()
	sample.IpAddresses.Utun3Ipv6 = generateIPv6()
	sample.IpAddresses.Utun4Ipv6 = generateIPv6()
	sample.IpAddresses.En0Ipv4 = generateIPv4()

	sample.BundleIdentifier = bundleId
	sample.BundleVersion = bundleVersion
	sample.CountryCode = countryCode
	sample.LanguageCode = languageCode
	sample.Version = iosVersion

	sample.Uuid = strings.ToUpper(uuid.New().String())
	sample.VenderId = strings.ToUpper(uuid.New().String())
	return sample
}

func ToFingeprint(fps string) model.Fingerprint {
	var fp model.Fingerprint
	json.Unmarshal([]byte(fps), &fp)
	return fp
}
