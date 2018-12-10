// Package analyze implements the Flake8 runner and parser.
package analyze

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/analyze/cache"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/satori/go.uuid"
)

const (
	engineName   = "flake8"
	languageName = "python3"
)

var domainUUID = uuid.Must(
	uuid.FromString("badefdd6-8997-425d-9d9e-ae31a01daf0c"))
var flake8Cache = cache.New(30*time.Minute, 5*time.Minute)

type flake8Message struct {
	Code         string `json:"code"`
	Filename     string `json:"filename"`
	LineNumber   uint   `json:"line_number"`
	ColumnNumber uint   `json:"column_number"`
	Text         string `json:"text"`
	PhysicalLine string `json:"physical_line"`
}

type flake8Result map[string][]flake8Message

func (f flake8Result) toRoast(username string, code string) (roast *model.RoastResult) {
	result := f["stdin"]

	roast = model.NewRoastResult(username, languageName, code)

	for _, message := range result {
		hash := uuid.NewV5(domainUUID, message.Code)

		switch []rune(message.Code)[0] {
		case 'W', 'C':
			roast.AddWarning(hash,
				message.LineNumber,
				message.ColumnNumber,
				engineName,
				message.Code,
				message.Text)
		case 'F', 'I', 'R', 'E':
			roast.AddError(hash,
				message.LineNumber,
				message.ColumnNumber,
				engineName,
				message.Code,
				message.Text)
		default:
			log.Printf("unhandled message code: %s", message.Code)
		}
	}

	roast.CalculateScore()

	return
}

// WithFlake8 statically analyzes the code with Flake8 and parses the result.
func WithFlake8(username, code string) (result *model.RoastResult, err error) {
	if r, ok := flake8Cache.Get(code); ok {
		if result, ok := r.(model.RoastResult); ok {
			return &result, err
		}
	}

	var r flake8Result

	cmd := exec.Command("python3", "-m",
		"flake8",
		"--format=json",
		"--max-complexity=10",
		"--isolated",
		"--exit-zero",
		"-")

	cmd.Stdin = strings.NewReader(code)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	if err = json.NewDecoder(stdout).Decode(&r); err != nil {
		return
	}

	if err = cmd.Wait(); err != nil {
		return
	}

	result = r.toRoast(username, code)
	flake8Cache.Set(code, *result)

	return
}