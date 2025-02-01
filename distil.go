package distil

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/divvot/distil/model"
)

type DistilSolver struct {
	fingeprint model.Fingerprint
	magic      uint32
	arcRandom  uint32
}

func (ds *DistilSolver) Solve(question string, session string) (string, string, error) {
	if ds.magic == 0 {
		ds.magic = calculateMagic(ds.fingeprint)
	}

	answerString := solve(question, ds.magic)
	answer := model.Answer{
		Answer:      answerString,
		Fingerprint: ds.fingeprint,
	}

	bytes, err := json.Marshal(answer)
	if err != nil {
		return "", "", err
	}

	str := strings.ReplaceAll(string(bytes), "/", "\\/")
	answerString = base64.StdEncoding.EncodeToString([]byte(str))

	clue := model.Clue{
		Session: session,
		Answer:  answerString,
	}

	bytes, err = json.Marshal(clue)
	if err != nil {
		return "", "", err
	}

	solution, _ := Encrypt(string(bytes), ds.arcRandom)

	return solution, answerString, nil
}

func (ds *DistilSolver) Magic() uint32 {
	return ds.magic
}

func (ds *DistilSolver) SetArcRandom(arcRandom uint32) {
	ds.arcRandom = arcRandom
}

func NewDistilSolver(fingeprint model.Fingerprint) *DistilSolver {
	return &DistilSolver{
		fingeprint: fingeprint,
		magic:      0,
	}
}
