package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/divvot/distil"
	"github.com/divvot/distil/model"
)

type distilHandler struct{}

func (h *distilHandler) Encrypt(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	var msg struct {
		AppManifest string `json:"appManifest"`
		ArcRandom   uint32 `json:"arcRandom"`
	}

	var manifestStruct struct {
		AppManifest string `json:"application_manifest"`
	}

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	manifestStruct.AppManifest = msg.AppManifest

	b, err := json.Marshal(manifestStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := distil.Encrypt(string(b), msg.ArcRandom)
	if err != nil {
		data = err.Error()
	}

	response := map[string]string{
		"p": data,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *distilHandler) Solve(w http.ResponseWriter, r *http.Request) {

	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	var challenge model.Challenge
	err := json.NewDecoder(r.Body).Decode(&challenge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var fingeprint model.Fingerprint

	if challenge.Fingerprint != nil {
		fingeprint = *challenge.Fingerprint
	} else if challenge.IsDebug {
		fingeprint = distil.DebugFingerprint()
	} else {
		fingeprint = distil.GenerateFingerprint(challenge.BundleId, challenge.BundleVersion, challenge.IOSVersion,
			challenge.LanguageCode, challenge.CountryCode)
	}

	solver := distil.NewDistilSolver(fingeprint)
	solver.SetArcRandom(challenge.ArcRandom)

	solution, answer, err := solver.Solve(challenge.Question, challenge.Session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	csolution := model.ChallengeSolution{
		Answer:      answer,
		Solution:    solution,
		Magic:       solver.Magic(),
		Fingerprint: fingeprint,
	}

	log.Printf("Solved challenge: %s\n", challenge.Question)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(csolution)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
