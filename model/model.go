package model

type Fingerprint struct {
	Uuid             string      `json:"uuid"`
	VenderId         string      `json:"vendorId"`
	IpAddresses      IpAddresses `json:"ipAddresses"`
	Version          string      `json:"version"`
	BundleIdentifier string      `json:"bundleIdentifier"`
	LanguageCode     string      `json:"languageCode"`
	CountryCode      string      `json:"countryCode"`
	BundleVersion    string      `json:"bundleVersion"`
	IsSimulator      bool        `json:"isSimulator"`
	IsJailBroken     bool        `json:"isJailBroken"`
	Debugger         bool        `json:"debugger"`
	Device           Device      `json:"device"`
	Model            string      `json:"model"`
	Name             string      `json:"name"`
}

type Device struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type IpAddresses struct {
	Utun3Ipv6 string `json:"utun3/ipv6"`
	Anpi0Ipv6 string `json:"anpi0/ipv6"`
	En0Ipv4   string `json:"en0/ipv4"`
	Utun2Ipv6 string `json:"utun2/ipv6"`
	Lo0Ipv6   string `json:"lo0/ipv6"`
	Awdl0Ipv6 string `json:"awdl0/ipv6"`
	Llw0Ipv6  string `json:"llw0/ipv6"`
	Utun1Ipv6 string `json:"utun1/ipv6"`
	En0Ipv6   string `json:"en0/ipv6"`
	Utun0Ipv6 string `json:"utun0/ipv6"`
	Utun4Ipv6 string `json:"utun4/ipv6"`
	Lo0Ipv4   string `json:"lo0/ipv4"`
}

type Clue struct {
	// (I ran out of names lol)
	Session string `json:"session"`
	Answer  string `json:"answer"`
}

type Answer struct {
	Answer      string      `json:"answer"`
	Fingerprint Fingerprint `json:"fingerprint"`
}

type Challenge struct {
	Question      string       `json:"question"`
	Session       string       `json:"session"`
	BundleId      string       `json:"bundleId,omitempty"`
	BundleVersion string       `json:"bundleVersion,omitempty"`
	CountryCode   string       `json:"countryCode,omitempty"`
	LanguageCode  string       `json:"languageCode,omitempty"`
	IOSVersion    string       `json:"iosVersion,omitempty"`
	ArcRandom     uint32       `json:"arcRandom,omitempty"`
	IsDebug       bool         `json:"isDebug,omitempty"`
	Fingerprint   *Fingerprint `json:"fingerprint,omitempty"`
}

type ChallengeSolution struct {
	Solution    string      `json:"solution"`
	Answer      string      `json:"answer"`
	Magic       uint32      `json:"magic"`
	Fingerprint Fingerprint `json:"fingeprint"`
}
